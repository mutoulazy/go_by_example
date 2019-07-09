package tailf

import (
	"errors"
	"time"
	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
)

type CollectConf struct {
	LogPath string
	Topic   string
}

type TailfObj struct {
	tail        *tail.Tail
	collectConf CollectConf
}

type TextMsg struct {
	Msg string
	Topic string
}

type TailfObjMgr struct {
	tailfObjs []*TailfObj
	msgChan chan *TextMsg
}

var (
	tailfObjMgr *TailfObjMgr
)

func InitTailf(collectConfs []CollectConf, chanSize int) (err error) {
	if len(collectConfs) == 0 {
		err = errors.New("collectConfs len err")
		return
	}

	tailfObjMgr = &TailfObjMgr{
		msgChan: make(chan *TextMsg, chanSize),
	}

	for _, v := range collectConfs {
		tails, loadErr := tail.TailFile(v.LogPath, tail.Config{
			ReOpen:    true,
			Follow:    true,
			MustExist: false,
			Poll:      true,
		})
		if loadErr != nil {
			logs.Warn("tail file err:", loadErr)
			err = loadErr
			return
		}

		tailfObj := &TailfObj{
			collectConf: v,
			tail:        tails,
		}
		// tailfObj.tail = tails

		tailfObjMgr.tailfObjs = append(tailfObjMgr.tailfObjs, tailfObj)

		go readFileTail(tailfObj)
	}

	return
}

// 读取日志文件 把对应的数据输入到管道中
func readFileTail(tailObj *TailfObj) {
	for true {
		line, ok := <-tailObj.tail.Lines
		if !ok {
			logs.Warn("tail file close reopen, filename:%s\n", tailObj.tail.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		textMsg := &TextMsg{
			Msg: line.Text,
			Topic: tailObj.collectConf.Topic,
		}

		tailfObjMgr.msgChan <- textMsg
	}
}

// 从管道中获取一条数据
func GetOneLine() (msg *TextMsg) {
	msg = <- tailfObjMgr.msgChan
	return
}
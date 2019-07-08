package tailf

import (
	"errors"
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

type TiilfObjMgr struct {
	tailfs []*TailfObj
}

var (
	tiilfObjMgr *TiilfObjMgr
)

func InitTailf(collectConfs []CollectConf) (err error) {
	if len(collectConfs) == 0 {
		err = errors.New("collectConfs len err")
		return
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

		tiilfObjMgr.tailfs = append(tiilfObjMgr.tailfs, tailfObj)
	}

	return
}

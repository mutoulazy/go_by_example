package tailf

import (
	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
)

type CollectConf struct {
	LogPath string
	Topic string
}

func InitTailf(collectConfs []CollectConf) (err error) {
	filename := collectConfs[0].LogPath
	tails, err := tail.TailFile(filename, tail.Config{
		ReOpen:    true,
		Follow:    true,
		// Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		logs.Warn("tail file err:", err)
		return
	}
	logs.Debug(tails)
	return
}
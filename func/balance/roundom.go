package balance

import (
	"errors"
)

type RoundomBalance struct {
	curIndex int
}

func (p *RoundomBalance) DoBalance(insts []*Instance) (inst *Instance,err error) {
	length := len(insts)
	if length == 0 {
		err = errors.New("No instance")
		return
	}

	if (p.curIndex >= length) {
		p.curIndex = 0
	}
	inst = insts[p.curIndex]
	p.curIndex ++ 
	return
}
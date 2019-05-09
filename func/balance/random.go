package balance

import (
	"errors"
	"math/rand"
)

func init() {
	RegisterBalancer("random", &RandonBalance{})
}

type RandonBalance struct {

}

func (p *RandonBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance,err error) {
	if len(insts) == 0 {
		err = errors.New("No instance")
		return
	}

	length := len(insts)
	index := rand.Intn(length)
	inst = insts[index]
	return
}
package main

import (
	"../balance"
	"errors"
	"hash/crc32"
	"fmt"
	"math/rand"
)

// 利用接口注册新的负载算法
func init() {
	balance.RegisterBalancer("hash", &HashBalance{})
}

type HashBalance struct {

}

func (p *HashBalance) DoBalance(insts []*balance.Instance, key ...string) (inst *balance.Instance,err error) {
	length := len(insts)
	var defkey string = fmt.Sprintf("%d", rand.Int())
	if length == 0 {
		err = errors.New("No instance")
		return
	}

	if len(key) > 0 {
		defkey = key[0]
	}
	crcTable := crc32.MakeTable(crc32.IEEE)
	hashval := crc32.Checksum([]byte(defkey), crcTable)
	index := int(hashval) % length
	inst = insts[index]
	return
}
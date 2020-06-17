package main

import (
	"context"
	"fmt"
	etcdClient "github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	cli, err := etcdClient.New(etcdClient.Config{
		Endpoints:   []string{"localhost:2379", "localhost:2379", "localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Errorf("connect etcd failed, err:%s", err)
		return
	}
	fmt.Println("connect etcd success")
	defer cli.Close()

	rch := cli.Watch(context.Background(), "/logAgeny/conf/")
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}

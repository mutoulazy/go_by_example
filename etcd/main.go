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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "/logAgeny/conf/", "sample_value")
	cancel()
	if err != nil {
		fmt.Errorf("put etcd failed, err:%s", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "/logAgeny/conf/")
	cancel()
	if err != nil {
		fmt.Errorf("get etcd failed, err:%s", err)
		return
	}

	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
}

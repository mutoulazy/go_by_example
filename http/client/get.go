package main

import (
	"io/ioutil"
	"fmt"
	"net/http"
)

var url = []string{
	"http://www.baidu.com",
	"http://www.google.com",
	"http://www.taobao.com",
}

func head() {
	for _, v := range url {
		resp, err := http.Head(v)
		if err != nil {
			fmt.Printf("head %s failed, err:%v\n", v, err)
			continue
		}

		fmt.Printf("%s head succ, status:%v\n",v , resp.Status)
	}
}

func main() {
	res, err := http.Get("https://www.baidu.com/")
	if err != nil {
		fmt.Println("get err: ", err)
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("get data err: ", err)
		return
	}

	fmt.Println(string(data))
	head()
}
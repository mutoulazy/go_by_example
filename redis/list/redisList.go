package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	defer conn.Close()
	
	// 向redis里面添加值
	_, err = conn.Do("lpush", "book_list", "abc", "ceg", 400)
	if err != nil {
		fmt.Println("Set value err: ", err)
		return
	}
	// 从redis中获取值
	r, err := redis.String(conn.Do("lpop", "book_list"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}
	fmt.Println(r)
}
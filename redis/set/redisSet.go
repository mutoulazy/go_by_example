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
	_, err = conn.Do("MSet", "abc", 100, "efg", 300)
	if err != nil {
		fmt.Println("Set value err: ", err)
		return
	}
	// 从redis中获取值
	r, err := redis.Ints(conn.Do("MGet", "abc", "efg"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}

	// 设置过期时间
	_, err = conn.Do("expire", "abc", 10)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range r {
		fmt.Println(v)
	}
}
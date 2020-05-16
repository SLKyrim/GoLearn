package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	// 通过go，向redis写入和读取数据
	// 1. 连接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	fmt.Println("conn succ...", conn)
}

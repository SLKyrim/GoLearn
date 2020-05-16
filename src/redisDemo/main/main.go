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
	defer conn.Close() // 重要：延时关闭

	// 2. 通过go向redis写入数据 string [key-val]
	_, err = conn.Do("Set", "name", "SingleLong")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}

	// 3. 通过go向redis读取数据 string [key-val]
	// r 是 interface{}，需要类型转换
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get err=", err)
		return
	}

	fmt.Println("操作成功", r)
}

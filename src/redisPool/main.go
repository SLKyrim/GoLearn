package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// 定义一个全局的pool
var pool *redis.Pool

// 当启动程序时，就初始化连接池
func init() {

	pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲连接数
		MaxActive:   0,   // 和数据库的最大连接数，0表示无限制
		IdleTimeout: 100, // 最大空闲时间，一个连接在100秒内没被用过则放回到空闲连接中
		Dial: func() (redis.Conn, error) { // 初始化连接的代码，声明连接哪个ip的redis
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	// 先从连接池中取出一个连接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "name", "RedisPoll")
	if err != nil {
		fmt.Println("conn.Do err =", err)
		return
	}

	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("conn.Do err =", err)
		return
	}

	fmt.Println("r =", r)
}

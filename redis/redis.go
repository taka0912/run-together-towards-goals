package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"os"
)

func Connection() redis.Conn {
	const IpPort = "127.0.0.1:6379"

	//redisに接続
	c, err := redis.Dial("tcp", IpPort)
	if err != nil {
		panic(err)
	}
	return c
}

func Set(key string, value string, c redis.Conn){
	c.Do("SET", key, value)
}

func redisGet(key string, c redis.Conn) string {
	s, err := redis.String(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return s
}

func SetList(key uint, value []string, c redis.Conn){
	for _ , v := range value {
		c.Do("RPUSH", key, v)
	}
}
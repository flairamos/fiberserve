package redisutil

import (
	"errors"
	"fmt"
	"shoollove/config"

	"github.com/gomodule/redigo/redis"
)

var Client *redis.Conn

func Init() {
	conf := config.Config.Redis
	conn, err := redis.Dial("tcp", conf.Host)
	if err != nil {
		panic(errors.New("conn redis failed"))
	}
	_, err = conn.Do("auth", conf.Password)
	if err != nil {
		panic(errors.New("redis auth failed"))
	}

	// 选择第一个数据库
	_, err = conn.Do("SELECT", conf.Database)
	if err != nil {
		panic(errors.New("redis select database failed"))
	}
	Client = &conn
	fmt.Println("conn redis success")
}

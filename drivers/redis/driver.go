package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

type ConfigRedis struct {
	DB_Host string
	DB_Port string
}

func (c ConfigRedis) InitRedis() redis.Conn {
	dsn := fmt.Sprintf("%v:%v", c.DB_Host, c.DB_Port)
	conn, err := redis.Dial("tcp", dsn)
	if err != nil {
		log.Panic(err)
	}
	return conn
}

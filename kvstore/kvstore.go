package kvstore

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

var Redis redis.Conn

func Dial(connUrl string) {
	conn, err := redis.Dial("tcp", connUrl)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer conn.Close()

	Redis = conn
}

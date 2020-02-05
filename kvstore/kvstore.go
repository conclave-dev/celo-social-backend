package kvstore

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

func Dial() redis.Conn {
	client, err := redis.Dial(connProto, connURL)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func IsUpdateExists(hash string) bool {
	client := Dial()
	defer client.Close()

	exists, err := redis.Bool(client.Do("HEXISTS", UPDATES, hash))
	if err != nil {
		log.Fatal(err)
	}

	return exists
}

func SetUpdate(hash string, update string) {
	client := Dial()
	defer client.Close()

	_, err := client.Do("HSET", UPDATES, hash, update)
	if err != nil {
		log.Fatal(err)
	}
}

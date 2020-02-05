package kvstore

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

// Dial connects to redis and returns the connected client
func Dial() redis.Conn {
	client, err := redis.Dial(connProto, connURL)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

// IsUpdateExists checks whether an update already exists
func IsUpdateExists(hash string) bool {
	client := Dial()
	defer client.Close()

	exists, err := redis.Bool(client.Do("HEXISTS", UPDATES, hash))
	if err != nil {
		log.Fatal(err)
	}

	return exists
}

// SetUpdate stores the update hash along with the data for reference
func SetUpdate(hash string, update string) {
	client := Dial()
	defer client.Close()

	_, err := client.Do("HSET", UPDATES, hash, update)
	if err != nil {
		log.Fatal(err)
	}
}

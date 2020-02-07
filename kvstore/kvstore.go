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

func Exists(hashKey string, field string) (bool, error) {
	client := Dial()
	defer client.Close()
	return redis.Bool(client.Do("HEXISTS", hashKey, field))
}

func Set(hashKey string, field string, value string) (interface{}, error) {
	client := Dial()
	defer client.Close()
	return client.Do("HSET", hashKey, field, value)
}

func Get(hashKey string, field string) (string, error) {
	client := Dial()
	defer client.Close()
	return redis.String(client.Do("HGET", hashKey, field))
}

func Delete(hashKey string, field string) (interface{}, error) {
	client := Dial()
	defer client.Close()
	return client.Do("HDEL", hashKey, field)
}

// IsUpdateExists checks whether an update already exists
func DoesUpdateExist(updateHash string) bool {
	b, err := Exists(UPDATES, updateHash)
	if err != nil {
		log.Fatal(err)
	}

	return b
}

// SetUpdate stores the update hash along with the data for reference
func SetUpdate(updateHash string, update string) (interface{}, error) {
	return Set(UPDATES, updateHash, update)
}

// GetUpdate retries the update at the updateHash field
func GetUpdate(updateHash string) (string, error) {
	return Get(UPDATES, updateHash)
}

// DeleteUpdate deletes the hash from the store
func DeleteUpdate(updateHash string) (interface{}, error) {
	return Delete(UPDATES, updateHash)
}

func DoesUserExist(username string) bool {
	b, err := Exists(USERS, username)
	if err != nil {
		log.Fatal(err)
	}

	return b
}

func SetUser(username string, user string) (interface{}, error) {
	return Set(USERS, username, user)
}

func GetUser(username string) (string, error) {
	return Get(USERS, username)
}

func DeleteUser(username string) (interface{}, error) {
	return Delete(USERS, username)
}

func DoesAddressExist(address string) bool {
	b, err := Exists(ADDRESSES, address)
	if err != nil {
		log.Fatal(err)
	}

	return b
}

func SetAddress(address string, username string) (interface{}, error) {
	return Set(ADDRESSES, address, username)
}

func GetAddress(address string) (string, error) {
	return Get(ADDRESSES, address)
}

func DeleteAddress(address string) (interface{}, error) {
	return Delete(ADDRESSES, address)
}

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

// DoesProfileExist checks whether an update already exists
func DoesProfileExist(updateProfile string) bool {
	b, err := Exists(PROFILES, updateProfile)
	if err != nil {
		log.Fatal(err)
	}

	return b
}

// SetProfile stores the update hash along with the data for reference
func SetProfile(updateProfile string, update string) (interface{}, error) {
	return Set(PROFILES, updateProfile, update)
}

// GetProfile retries the update at the updateProfile field
func GetProfile(updateProfile string) (string, error) {
	return Get(PROFILES, updateProfile)
}

// DeleteProfile deletes the hash from the store
func DeleteProfile(updateProfile string) (interface{}, error) {
	return Delete(PROFILES, updateProfile)
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

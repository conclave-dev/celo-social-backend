package kvstore

import (
	"log"
	"strings"

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
	u := strings.ToLower(username)
	b, err := Exists(USERS, u)
	if err != nil {
		log.Fatal(err)
	}

	return b
}

func SetUser(username string, hash string) (interface{}, error) {
	u := strings.ToLower(username)
	return Set(USERS, u, hash)
}

func GetUser(username string) (string, error) {
	u := strings.ToLower(username)
	return Get(USERS, u)
}

func DeleteUser(username string) (interface{}, error) {
	u := strings.ToLower(username)
	return Delete(USERS, u)
}

func DoesAddressExist(address string) bool {
	a := strings.ToLower(address)
	b, err := Exists(ADDRESSES, a)
	if err != nil {
		log.Fatal(err)
	}

	return b
}

func SetAddress(address string, hash string) (interface{}, error) {
	a := strings.ToLower(address)
	return Set(ADDRESSES, a, hash)
}

func GetAddress(address string) (string, error) {
	a := strings.ToLower(address)
	return Get(ADDRESSES, a)
}

func DeleteAddress(address string) (interface{}, error) {
	a := strings.ToLower(address)
	return Delete(ADDRESSES, a)
}

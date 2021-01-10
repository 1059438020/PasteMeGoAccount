package handler

import (
	"github.com/go-redis/redis/v7"
	"github.com/wonderivan/logger"
	"time"
)

var redisClient *redis.Client

func LoadRedisClient(client *redis.Client) {
	redisClient = client
	Exist("hello")
	Exist("abc")
}

func Set(key string, value string, expiration int64) {
	err := redisClient.Set(key, value, time.Duration(expiration)).Err()
	if err != nil {
		logger.Painc(err)
		panic(err)
	}
}

func SetPerm(key string, value string) {
	Set(key, value, -1)
}

func Exist(key string) bool {
	exists := redisClient.Exists(key)
	result, err := exists.Result()
	if err != nil {
		logger.Painc(err)
		panic(err)
	}
	return result == 1
}

func Get(key string) string {
	get := redisClient.Get(key)
	result, err := get.Result()
	if err != nil {
		logger.Painc(err)
		panic(err)
	}
	return result
}

func HGet(key string, filed string) string {
	if !Exist(key) {
		return ""
	}
	get := redisClient.HGet(key, filed)
	result, err := get.Result()
	if err != nil {
		logger.Painc(err)
		panic(err)
	}
	return result
}
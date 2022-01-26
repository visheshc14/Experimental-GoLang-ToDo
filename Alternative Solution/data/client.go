package data

import (
	"github.com/go-redis/redis/v8"
)

func Client() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No Password Set
		DB:       0,  // Use Default DB
	})

	return rdb
}

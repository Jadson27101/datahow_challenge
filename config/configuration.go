package config

import "github.com/go-redis/redis/v8"

func NewDefaultRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}
func NewDefaultServerPorts() (string, string) {
	return "9102", "5000"
}

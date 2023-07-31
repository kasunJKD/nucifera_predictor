package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
)

type redisConfig struct {
	RedisClient *redis.Client
}

func Connect(host, port, password string) *redis.Client {

	log.Printf("Connecting to Redis\n")

	redisConn := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       0,
	})

	_, err := redisConn.Ping(context.Background()).Result()

	if err != nil {
		log.Fatal("Error connecting to redis: ", err)
	}

	return redisConn

}

func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST") + ":" +os.Getenv("REDIS_PORT"),
		Password: os.Getenv("PASSWORD"),
		DB: dbNo,
	})

	return rdb
}
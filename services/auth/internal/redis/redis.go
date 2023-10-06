package redis

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

type Handler struct {
	*redis.Client
}

func Init() Handler {
	rc := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})

	if _, err := rc.Ping(context.Background()).Result(); err != nil {
		log.Fatalf("%v", err.Error())
	}

	log.Println("✅ Redis connected!")

	return Handler{rc}
}

package redis

import (
	"context"
	"fmt"
	"pwsd_keeper/pkg/utility"

	"github.com/redis/go-redis/v9"
)

const (
	REDIS_ADDRESS  = "REDIS_ADDRESS"
	REDIS_PORT     = "REDIS_PORT"
	REDIS_PASSWORD = "REDIS_PASSWORD"
)

func handleErr(err error) {
	if err != nil {
		panic("can not read from env")
	}
}

func New() (context.Context, *redis.Client) {
	var ctx = context.Background()
	redisAddress, err := utility.LoadFromEnv(REDIS_ADDRESS)
	handleErr(err)
	redisPort, err := utility.LoadFromEnv(REDIS_PORT)
	handleErr(err)
	redisPassword, err := utility.LoadFromEnv(REDIS_PASSWORD)
	handleErr(err)

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisAddress, redisPort),
		Password: redisPassword,
		DB:       0,
	})

	return ctx, rdb
}

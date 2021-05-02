package cache

import (
	"os"
	"backend/util"
	"strconv"

	"github.com/go-redis/redis"
)

// RedisClient Redis Client
var RedisClient *redis.Client

// Redis init
func Redis() {
	db, _ := strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDIS_ADDR"),
		Password:   os.Getenv("REDIS_PW"),
		DB:         int(db),
		MaxRetries: 1,
	})

	_, err := client.Ping().Result()

	if err != nil {
		util.Log().Panic("连接Redis不成功", err)
	}

	RedisClient = client
}

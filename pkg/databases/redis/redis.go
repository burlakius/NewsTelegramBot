package redisdb

import (
	"context"
	"strconv"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

var redisCache cache.Cache

func RedisConnect(address, port string) {
	redisRing := *redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			address: port,
		},
	})

	redisCache = *cache.New(&cache.Options{
		Redis: redisRing,
	})
}

func SetLanguage(chatID int64, language string) error {
	err := redisCache.Set(&cache.Item{
		Key:   strconv.FormatInt(chatID, 10),
		Value: language,
	})

	return err
}

func GetLanguage(chatID int64) (string, error) {
	var wanted string
	err := redisCache.Get(context.TODO(), strconv.FormatInt(chatID, 10), &wanted)

	if err != nil {
		return "", err
	}

	return wanted, nil
}

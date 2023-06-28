package redisdb

import (
	"context"
	"strconv"
	"sync"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

var redisCache botCache

type botCache struct {
	rc cache.Cache
	mu sync.Mutex
}

func RedisConnect(address, port string) {
	redisRing := *redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			address: port,
		},
	})

	redisCache.rc = *cache.New(&cache.Options{
		Redis: redisRing,
	})
}

func SetLanguage(chatID int64, language string) error {
	redisCache.mu.Lock()
	defer redisCache.mu.Unlock()

	err := redisCache.rc.Set(&cache.Item{
		Key:   strconv.FormatInt(chatID, 10),
		Value: language,
	})

	return err
}

func GetLanguage(chatID int64) (string, error) {
	redisCache.mu.Lock()
	defer redisCache.mu.Unlock()

	var wanted string
	err := redisCache.rc.Get(context.TODO(), strconv.FormatInt(chatID, 10), &wanted)

	if err != nil {
		return "", err
	}

	return wanted, nil
}

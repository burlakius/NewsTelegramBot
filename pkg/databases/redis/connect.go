package redisdb

import (
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

func RedisConnect(addrToLangSess, addrToChatSt string) {
	chatStates.rc = *cache.New(&cache.Options{
		Redis: *redis.NewClient(&redis.Options{
			Addr: addrToChatSt,
		}),
	})
	languageSessions.rc = *cache.New(&cache.Options{
		Redis: *redis.NewClient(&redis.Options{
			Addr: addrToLangSess,
		}),
	})
}

package redisdb

import (
	"context"
	"strconv"
	"sync"

	"github.com/go-redis/cache/v8"
)

var languageSessions botCache

type botCache struct {
	rc cache.Cache
	mu sync.Mutex
}

func GetLanguage(chatID int64) (string, error) {
	languageSessions.mu.Lock()
	defer languageSessions.mu.Unlock()

	var wanted string
	err := languageSessions.rc.Get(context.TODO(), strconv.FormatInt(chatID, 10), &wanted)

	if err != nil {
		return "", err
	}

	return wanted, nil
}

func SetLanguage(chatID int64, language string) error {
	languageSessions.mu.Lock()
	defer languageSessions.mu.Unlock()

	err := languageSessions.rc.Set(&cache.Item{
		Key:   strconv.FormatInt(chatID, 10),
		Value: language,
	})

	return err
}

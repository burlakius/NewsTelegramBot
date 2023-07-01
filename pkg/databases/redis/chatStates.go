package redisdb

import (
	"context"
	"strconv"

	"github.com/go-redis/cache/v8"
)

var chatStates botCache

func GetChatState(chatID int64) (string, error) {
	chatStates.mu.Lock()
	defer chatStates.mu.Unlock()

	var wanted string
	err := chatStates.rc.Get(context.TODO(), strconv.FormatInt(chatID, 10), &wanted)

	if err != nil {
		return "", err
	}
	return wanted, nil
}

func SetChatState(chatID int64, state string) error {
	chatStates.mu.Lock()
	defer chatStates.mu.Unlock()

	err := chatStates.rc.Set(&cache.Item{
		Key:   strconv.FormatInt(chatID, 10),
		Value: state,
	})

	return err
}

func DoneChatState(chatID int64) error {
	chatStates.mu.Lock()
	defer chatStates.mu.Unlock()

	err := chatStates.rc.Delete(context.TODO(), strconv.FormatInt(chatID, 10))

	return err
}

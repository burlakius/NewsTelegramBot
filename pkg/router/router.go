package router

import (
	"context"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type router struct {
	bot                *tgbotapi.BotAPI
	updateConfig       tgbotapi.UpdateConfig
	dispatcherChannels channelsCycle
}

func (r *router) StartPolling(closeContext context.Context) {
	for {
		select {
		case <-closeContext.Done():
			logrus.Infoln("Polling shutdown...")
			r.dispatcherChannels.close()
			logrus.Infoln("Dispatcher channels closed...")

			return
		default:
		}

		updates, err := r.bot.GetUpdates(r.updateConfig)
		if err != nil {
			logrus.Warnln(err)
			logrus.Warnln("Failed to get updates, retrying in 3 seconds...")
			time.Sleep(time.Second * 3)

			continue
		}

		for _, update := range updates {
			logrus.Infof("Update[%v] received\n", update.UpdateID)
			logrus.Debugln(update)
			if update.UpdateID >= r.updateConfig.Offset {
				r.updateConfig.Offset = update.UpdateID + 1

				r.dispatcherChannels.next() <- update
				logrus.Infof("Update[%v] sended to dispatcher\n", update.UpdateID)
			}
		}
	}
}

type channelsCycle struct {
	channels []chan tgbotapi.Update
	length   uint
	index    uint
}

func (cq *channelsCycle) next() chan tgbotapi.Update {
	cq.index += 1
	if cq.index > cq.length {
		cq.index = 0
	}

	return cq.channels[cq.index]
}

func (cq *channelsCycle) close() {
	for _, channel := range cq.channels {
		close(channel)
	}
}

func newChannelsCycle(channels []chan tgbotapi.Update) channelsCycle {
	return channelsCycle{
		channels: channels,
		length:   uint(len(channels)),
		index:    0,
	}
}

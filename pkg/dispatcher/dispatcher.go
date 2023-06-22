package dispatcher

import (
	handlers "news_telegram_bot/pkg/handlers"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type Dispatcher struct {
	MessageHandlersList       []handlers.MessageHandler
	EditedMessageHandlersList []handlers.MessageHandler
	CallbackQueryHandlersList []handlers.CallbackQueryHandler

	// NOT USED HANDLERS
	// ChannelPostHandlersList       []handler.MessageHandler
	// EditedChannelPostHandlersList []handler.MessageHandler
	// InlineQueryHandlersList        []handler.Handler
	// ChosenInlineResultHandlersList []handler.Handler
	// ShippingQueryHandlersList    []handler.Handler
	// PreCheckoutQueryHandlersList []handler.Handler
	// PollHandlersList             []handler.Handler
	// PollAnswerHandlersList       []handler.Handler
	// MyChatMemberHandlersList    []handler.Handler
	// ChatMemberHandlersList      []handler.Handler
	// ChatJoinRequestHandlersList []handler.Handler
}

func (d *Dispatcher) WaitUpdates(bot *tgbotapi.BotAPI, dispatcherChannel chan tgbotapi.Update, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	for update := range dispatcherChannel {
		logrus.Debugf("Update[%v] is received by dispatcher\n", update.UpdateID)

		if update.Message != nil {
			logrus.Debugf("Update[%v] is a message", update.UpdateID)
			for _, handler := range d.MessageHandlersList {
				if handler.Check(update.Message) {
					handler.CallbackFunc(update.Message, bot)
					break
				}
			}
		} else if update.EditedMessage != nil {
			logrus.Debugf("Update[%v] is a edited message", update.UpdateID)
			for _, handler := range d.EditedMessageHandlersList {
				if handler.Check(update.Message) {
					handler.CallbackFunc(update.Message, bot)
					break
				}
			}
		} else if update.CallbackQuery != nil {
			logrus.Debugf("Update[%v] is a callback query", update.UpdateID)
			for _, handler := range d.CallbackQueryHandlersList {
				if handler.Check(update.CallbackQuery) {
					handler.CallbackFunc(update.CallbackQuery, bot)
					break
				}
			}
		}
	}

	logrus.Infoln("Dispatcher closed...")
}

func (d *Dispatcher) RegisterMessageHandler(
	callbackFunc func(*tgbotapi.Message, *tgbotapi.BotAPI),
	filters ...func(*tgbotapi.Message) bool,
) {

	d.MessageHandlersList = append(
		d.MessageHandlersList,
		handlers.MessageHandler{
			CallbackFunc: callbackFunc,
			Filters:      filters,
		},
	)
	logrus.Debugln("Registrated new message handler...")
}

func (d *Dispatcher) RegisterEditedMessageHandler(
	callbackFunc func(*tgbotapi.Message, *tgbotapi.BotAPI),
	filters ...func(*tgbotapi.Message) bool,
) {

	d.EditedMessageHandlersList = append(
		d.EditedMessageHandlersList,
		handlers.MessageHandler{
			CallbackFunc: callbackFunc,
			Filters:      filters,
		},
	)
	logrus.Debugln("Registrated new edited message handler...")
}

func (d *Dispatcher) RegisterCallbackQueryHandler(
	callbackFunc func(*tgbotapi.CallbackQuery, *tgbotapi.BotAPI),
	Filters ...func(*tgbotapi.CallbackQuery) bool,
) {

	d.CallbackQueryHandlersList = append(
		d.CallbackQueryHandlersList,
		handlers.CallbackQueryHandler{
			CallbackFunc: callbackFunc,
			Filters:      Filters,
		},
	)
	logrus.Debugln("Registrated new callback query handler...")
}

func NewDispatcher() Dispatcher {
	return Dispatcher{
		MessageHandlersList:       make([]handlers.MessageHandler, 0, 1),
		EditedMessageHandlersList: make([]handlers.MessageHandler, 0, 1),
		CallbackQueryHandlersList: make([]handlers.CallbackQueryHandler, 0, 1),
	}
}

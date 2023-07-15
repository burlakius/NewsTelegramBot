package handlers

import (
	"news_telegram_bot/internal/config"
	"news_telegram_bot/pkg/databases/mariadb"
	redisdb "news_telegram_bot/pkg/databases/redis"
	"news_telegram_bot/pkg/translator"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang.org/x/text/message"
)

func adminAuthentication(chatID int64, printer *message.Printer, bot *tgbotapi.BotAPI) {
	responceMessage := tgbotapi.NewMessage(chatID, printer.Sprintf(
		"Привіт усім!\nЯ - @%v, створений, щоб допомагати адмінам у публікації нових новин та відповідати на питання студентів.\n\nЯ співпрацюю тільки зі своїми адміністраторами, тому пойдіть, будь ласка, аутентифікацію.\n\nВведіть пароль:",
		bot.Self.UserName,
	))

	if !mariadb.IsAdminChat(chatID) {
		err := redisdb.SetChatState(chatID, "WaitPassword")
		if err != nil {
			sendBotStorageError(chatID, bot)
			return
		}
		bot.Send(responceMessage)
	}
}

func setAdminsCommands(chatID int64, printer *message.Printer, bot *tgbotapi.BotAPI) {
	adminCommands := tgbotapi.NewSetMyCommandsWithScope(
		tgbotapi.NewBotCommandScopeChat(chatID),
		tgbotapi.BotCommand{
			Command:     "get_question",
			Description: printer.Sprintf("Вивести питання користувачів з можливістю відповіді"),
		},
		tgbotapi.BotCommand{
			Command:     "set_news",
			Description: printer.Sprintf("Додати нову звичайну новину"),
		},
		tgbotapi.BotCommand{
			Command:     "set_important_news",
			Description: printer.Sprintf("Додати нову важливу новину"),
		},
		tgbotapi.BotCommand{
			Command:     "help",
			Description: printer.Sprintf("Детальніше про команди"),
		},
	)

	bot.Send(adminCommands)
}

func receiveAdminPassword(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	defer redisdb.DoneChatState(message.Chat.ID)

	printer, err := translator.GetPrinterByChatID(message.Chat.ID)
	if err != nil {
		sendLanguageError(message.Chat.ID, bot)
		return
	}

	responceMessage := tgbotapi.NewMessage(message.Chat.ID, "")
	if message.Text == config.AdminPassword {
		responceMessage.Text = printer.Sprintf("Пароль вірний ✅\nНатисніть /help якщо вам потрібно дізнатись, як зі мною співпрацювати, або у вас не відображаються меню команд")
		mariadb.AddNewAdminChat(message.Chat.ID)
		defer setAdminsCommands(message.Chat.ID, printer, bot)
	} else {
		responceMessage.Text = printer.Sprintf("Пароль невірний ❌")
		defer bot.Send(
			tgbotapi.LeaveChatConfig{ChatID: message.Chat.ID},
		)
	}

	bot.Send(responceMessage)
}

func helpForAdmins(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	printer, err := translator.GetPrinterByChatID(message.Chat.ID)
	if err != nil {
		sendLanguageError(message.Chat.ID, bot)
		return
	}

	responceMessage := tgbotapi.NewMessage(
		message.Chat.ID,
		printer.Sprintf("Я ваш особистий телеграм-бот, готовий допомогти вам з отриманням та управлінням новинами.\n\nОсь команди, які ви можете використовувати:\n\n/get_question - отримати одне питання з можливістю відповіді, або видалення(питання не видалиться до тих пір, поки ви самі цього не захочете)\n/add_news - Додати нову звичайну новину.\n/add_important_news - Додати нову важливу новину.\n\n!Важливі новини відрізняються від звичайних тим, що користувач не зможе відключити собі надходження важливих\n\n/help - Детальніше про команди та їх використання.\n\nЯ готовий надати вам всю необхідну інформацію та виконувати ваші запити. Просто надішліть мені одну з цих команд, і я буду радий допомогти!"),
	)

	setAdminsCommands(message.Chat.ID, printer, bot)
	bot.Send(responceMessage)
}

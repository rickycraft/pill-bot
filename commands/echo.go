package commands

import (
	"time"

	bot "pill-bot/bot"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Echo(tgbot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	text := time.Now().Format("2006-01-02 15:04:05")
	bot.SendMessage(tgbot, update.Message.Chat.ID, text)
}

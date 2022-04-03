package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

func SendMessage(bot *tgbotapi.BotAPI, chatId int64, message string) {
	msg := tgbotapi.NewMessage(chatId, message)
	_, err := bot.Send(msg)
	if err != nil {
		log.Error("Error sending message <%s> to %d", message, chatId)
	}
}

func SendMessageAdmin(bot *tgbotapi.BotAPI, message string) {
	admin := GetAdmin()
	SendMessage(bot, admin, message)
}

func SendMessagePatient(bot *tgbotapi.BotAPI, message string) {
	patient := GetPatient()
	SendMessage(bot, patient, message)
}

func SendMessageAll(bot *tgbotapi.BotAPI, message string) {
	admin := GetAdmin()
	patient := GetPatient()
	SendMessage(bot, admin, message)
	SendMessage(bot, patient, message)
}

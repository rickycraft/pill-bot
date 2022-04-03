package main

import (
	"pill-bot/server"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

func StartBot() {
	server, err := server.StartServer()
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
		return
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := server.Bot.GetUpdatesChan(u)

	for update := range updates {
		// handle commands
		if update.Message == nil {
			continue
		}
		var message string
		if update.Message.IsCommand() {
			message = handle_commands(update)
		} else {
			message = handle_messages(update)
		}
		log.Infof("msg response %s", message)
	}
}

func handle_commands(update tgbotapi.Update) string {
	switch update.Message.Command() {
	case "start":
		log.Info("recognized start command")
		return "starting bot"
	case "echo":
		log.Info(time.Now().Format("2006-01-02 15:04:05"))
		return "echo"
	case "stats":
		log.Info("stats not yet implemented")
		return "stats not yet implemented"
	default:
		log.Info("unknown command")
		return "unknown command"
	}
}

func handle_messages(update tgbotapi.Update) string {
	switch update.Message.Text {
	case "pill":
		log.Infof("recognized pill message from %d", update.Message.Chat.ID)
		return "pill message"
	case "box":
		log.Infof("recognized box message from %d", update.Message.Chat.ID)
		return "box not yet implemented"
	default:
		log.Info("unknown message")
		return "unknown message"
	}
}

func main() {
	StartBot()
}

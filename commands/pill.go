package commands

import (
	"fmt"
	"pill-bot/bot"
	"pill-bot/database"
	"pill-bot/server"

	log "github.com/sirupsen/logrus"
)

const (
	takeMsg = "Presa"
)

func Pill(server *server.Server) {
	if !database.CanTake(server.DB) {
		updateFailed(server, fmt.Errorf("cannot take pill"))
		return
	}
	if err := database.Take(server.DB); err != nil {
		updateFailed(server, err)
		return
	}
	updateSuccess(server)
}

func updateFailed(server *server.Server, err error) {
	log.Errorf("Error updating pill: %v", err)
	bot.SendMessageAdmin(server.Bot, fmt.Sprintf("Error updating pill: %v", err))
}

func updateSuccess(server *server.Server) {
	bot.SendMessageAll(server.Bot, takeMsg)
}

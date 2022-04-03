package server

import (
	"database/sql"
	"os"
	"pill-bot/database"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	DB  *sql.DB
	Bot *tgbotapi.BotAPI
}

const (
	token_fname = "bot_token"
)

func StartServer() (*Server, error) {
	db := database.Init()

	bot, err := tgbotapi.NewBotAPI(read_token(token_fname))
	if err != nil {
		return nil, err
	}

	return &Server{
		DB:  db,
		Bot: bot,
	}, nil
}

func read_token(fileName string) string {
	dat, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Cannot read token from %s", fileName)
	}
	return string(dat)
}

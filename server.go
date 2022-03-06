package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	token_fname = "bot_token"
)

func main() {
	log.Info(read_token(token_fname))
}

func read_token(fileName string) string {
	dat, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Cannot read token from %s", fileName)
	}
	return string(dat)
}

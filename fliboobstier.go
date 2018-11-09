package main

import (
	"log"
	"os"

	"github.com/chesnovsky/fliboobstier/boobs"

	"gopkg.in/telegram-bot-api.v4"
)

func main() {

	tgToken := os.Getenv("TG_TOKEN")

	bot, err := tgbotapi.NewBotAPI(tgToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		boobs.Showmeyourboobs(bot, update.Message)
	}

}

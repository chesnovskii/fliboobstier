package main

import (
	"os"

	"github.com/chesnovsky/fliboobstier/boobs"
	"github.com/chesnovsky/fliboobstier/logger"
	"gopkg.in/telegram-bot-api.v4"
)

func main() {

	tgToken := os.Getenv("TG_TOKEN")

	bot, err := tgbotapi.NewBotAPI(tgToken)
	if err != nil {
		logger.Logger.Fatal(err)
	}

	logger.Logger.Infof("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		boobs.Showmeyourboobs(bot, update.Message)
	}

}

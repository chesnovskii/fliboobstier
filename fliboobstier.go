package main

import (
	"github.com/chesnovsky/fliboobstier/bot"
	"github.com/chesnovsky/fliboobstier/config"
	"github.com/chesnovsky/fliboobstier/logger"
)

func main() {

	logger.Init("")
	err, mainConfig := config.GetConfig()
	if err != nil {
		logger.Logger.Fatalf("Config load failed:\n%v\n", err)
	}
	bot.RunBot(mainConfig)

	// tgToken := os.Getenv("TG_TOKEN")

	// bot, err := tgbotapi.NewBotAPI(tgToken)
	// if err != nil {
	// 	logger.Logger.Fatalf("Cannot init Telegram BotApi:\n%v", err)
	// }

	// logger.Logger.Infof("Authorized on account %s", bot.Self.UserName)

	// u := tgbotapi.NewUpdate(0)
	// u.Timeout = 60

	// updates, err := bot.GetUpdatesChan(u)

	// for update := range updates {
	// 	boobs.Showmeyourboobs(bot, update.Message)
	// }

}

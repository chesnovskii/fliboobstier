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
}

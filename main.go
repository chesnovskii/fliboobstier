// Main package of Fliboobstier TG bot
package main

import (
	"github.com/chesnovsky/fliboobstier/bot"
	"github.com/chesnovsky/fliboobstier/config"
	"github.com/chesnovsky/fliboobstier/logger"
	"github.com/chesnovsky/fliboobstier/storage"
)

func main() {
	logger.Init("")
	configPath := "config.yml"
	mainConfig, err := config.GetConfig(configPath)
	if err != nil {
		logger.Logger.Fatalf("Config load failed: %+v\n", err)
	}
	Storage, err := storage.NewSqLite("fliboobstier.db")
	if err != nil {
		logger.Logger.Fatalf("Database init failed: %+v\n", err)
	}

	Bot, _ := bot.InitBot(&mainConfig, &Storage)
	Bot.RunBot()
}

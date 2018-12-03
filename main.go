// Main package of Fliboobstier TG bot
package main

import (
	"github.com/chesnovsky/fliboobstier/bot"
	"github.com/chesnovsky/fliboobstier/config"
	"github.com/chesnovsky/fliboobstier/logger"
)

func main() {
	logger.Init("")
	configPath := "config.yml"
	mainConfig, err := config.GetConfig(configPath)
	if err != nil {
		logger.Logger.Fatalf("Config load failed:\n%v\n", err)
	}
	bot.RunBot(mainConfig)
}

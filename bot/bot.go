package bot

import (
	"github.com/chesnovsky/fliboobstier/config"
	"github.com/chesnovsky/fliboobstier/logger"
	"github.com/chesnovsky/fliboobstier/storage"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

type Bot struct {
	TgBot       *tgbotapi.BotAPI
	Storage     *storage.StorageInstance
	Config      *config.MainConfig
	Transitions []Transition
}

type Transition struct {
	Name   string
	ChatID int64
	UserID int
	Data   map[string]string
}

func InitBot(mainConfig *config.MainConfig, storage *storage.StorageInstance) (Bot, error) {
	var BotInstance Bot
	logger.Logger.Info("Starting a bot...")
	logger.Logger.Debugf("Got config:\n%v", mainConfig)

	tgBot, err := tgbotapi.NewBotAPI(mainConfig.TgToken)
	if err != nil {
		logger.Logger.Fatalf("Cannot start a TG API:\n%v", err)
	}
	logger.Logger.Infof("Started on account http://t.me/%s", tgBot.Self.UserName)

	BotInstance = Bot{
		TgBot:   tgBot,
		Storage: storage,
		Config:  mainConfig,
	}
	return BotInstance, nil
}

func (BotInstance *Bot) RunBot() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := BotInstance.TgBot.GetUpdatesChan(u)
	for update := range updates {
		logger.Logger.Debugf("Got raw message: %+v\n", update.Message)
		if update.Message != nil {
			if !BotInstance.processTransitions(update.Message) {
				if update.Message.IsCommand() {
					BotInstance.processAdminCommands(update.Message)
				} else {
					BotInstance.processRegexActions(update.Message)
				}
			}
		}
	}
}

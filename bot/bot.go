package bot

import (
    "github.com/chesnovsky/fliboobstier/config"
    "github.com/chesnovsky/fliboobstier/logger"
    tgBotApi "gopkg.in/telegram-bot-api.v4"
)

// var (
//     TgBotApi = tgBotApi.BotAPI{}
// )

// func initBotApi(token string) error {
//     _, TgBotApi := tgBotApi.NewBotAPI(token)
//     return nil
// }

func RunBot(config config.MainConfig) {
    logger.Logger.Debug("Starting a bot")
    logger.Logger.Debugf("Got config:\n%v", config)
    // err := initBotApi(config.TgToken)
    // var TgBotApi tgBotApi.BotAPI
    err, tgBot := tgBotApi.NewBotAPI(config.TgToken)
    if err != nil {
        logger.Logger.Fatalf("Cannot start a TG API:\n%v", err)
    }
    logger.Logger.Infof("Authorized on account %v", tgBot)
}

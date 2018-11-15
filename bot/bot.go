package bot

import (
    "regexp"

    "github.com/chesnovsky/fliboobstier/config"
    "github.com/chesnovsky/fliboobstier/logger"
    "gopkg.in/telegram-bot-api.v4"
)

var (
    TgBot = tgbotapi.BotAPI{}
)

// func initBotApi(token string) error {
//     var err error
//     TgBot, err = tgbotapi.NewBotAPI(token)
//     return err
// }

func RunBot(mainConfig config.MainConfig) {
    logger.Logger.Debug("Starting a bot")
    logger.Logger.Debugf("Got config:\n%v", mainConfig)

    TgBot, err := tgbotapi.NewBotAPI(mainConfig.TgToken)
    if err != nil {
        logger.Logger.Fatalf("Cannot start a TG API:\n%v", err)
    }
    logger.Logger.Infof("Started on account @%s", TgBot.Self.UserName)

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60
    updates, err := TgBot.GetUpdatesChan(u)
    for update := range updates {
        for catchType, catchConfig := range mainConfig.Catches {
            logger.Logger.Debugf("Starting <%s> sender", catchType)
            sender(TgBot, update.Message, catchConfig)
        }
    }
}

func sender(bot *tgbotapi.BotAPI, message *tgbotapi.Message, senderConfig config.WordCatch) {
    logger.Logger.Debugf("Got message:\n%v\n", message.Sticker)
    logger.Logger.Debugf("Got local config:\n%v\n", senderConfig)
    localRegex := regexp.MustCompile(senderConfig.Regex)
    matched := localRegex.MatchString(message.Text)
    logger.Logger.Debugf("Regex is mached:\t%v\t", matched)
    if matched {
        msg := tgbotapi.NewStickerShare(message.Chat.ID, senderConfig.Stickers[0])
        logger.Logger.Debugf("Sending NORMA message:\n%v\n", msg)
        bot.Send(msg)
    }
}

package bot

import (
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

// RunBot is an init func for TG Bot
func RunBot(mainConfig config.MainConfig) {
    logger.Logger.Debug("Starting a bot")
    logger.Logger.Debugf("Got config:\n%v", mainConfig)

    TgBot, err := tgbotapi.NewBotAPI(mainConfig.TgToken)
    if err != nil {
        logger.Logger.Fatalf("Cannot start a TG API:\n%v", err)
    }
    logger.Logger.Infof("Started on account http://t.me/%s", TgBot.Self.UserName)

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60
    updates, err := TgBot.GetUpdatesChan(u)
    for update := range updates {
        logger.Logger.Debugf("Got raw update:\n%v\n", update.Message)
        if update.Message != nil {
            logger.Logger.Debugf("Got message images:\n%v\n", update.Message.Photo)
            logger.Logger.Debugf("Got message stickers:\n%v\n", update.Message.Sticker)
            logger.Logger.Debugf("Got message gif:\n%v\n", update.Message.Animation)
            logger.Logger.Debugf("Got message documents:\n%v\n", update.Message.Document)
            for catchType, catchConfig := range mainConfig.Catches {
                logger.Logger.Debugf("Starting <%s> sender", catchType)
                sender(TgBot, update.Message, catchConfig)
            }
        }
    }
}

func sender(bot *tgbotapi.BotAPI, message *tgbotapi.Message, senderConfig config.WordCatch) {
    logger.Logger.Debugf("Got local config:\n%v\n", senderConfig)
    logger.Logger.Debugf("Got message text:\n%v\n", message.Text)
    matched := senderConfig.Regex.MatchString(message.Text)
    logger.Logger.Debugf("Regex is mached:\t%v\t", matched)
    if matched {
        msg, err := createSenderMsg(senderConfig, message.Chat.ID)
        if err == nil {
            logger.Logger.Debugf("Sending message:\n%v\n", msg)
            bot.Send(msg)
        } else {
            logger.Logger.Errorf("Cannot send an reply to message:\n%v\n%v", message, err)
        }
    }
}

func createSenderMsg(senderConfig config.WordCatch, chatID int64) (tgbotapi.Chattable, error) {
    // select media type
    var err error
    msg := tgbotapi.NewStickerShare(chatID, senderConfig.Stickers[0])
    return msg, err
}

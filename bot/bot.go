package bot

import (
    "github.com/chesnovsky/fliboobstier/config"
    "github.com/chesnovsky/fliboobstier/logger"
    "gopkg.in/telegram-bot-api.v4"
    "math/rand"
)

type MediaTypesType struct{
    Wtf int
    Image int
    Sticker int
    Gif int
}

var (
    TgBot = tgbotapi.BotAPI{}
    MediaTypes = MediaTypesType{-1, 0, 1, 2}
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

func isNumInRange(rMin int, rMax int, num int) bool{
    if (rMin < num && num < rMax) {
        return true
    } else {
        return false
    }
}

func selectMediaType(config config.WordCatch) int {
    imagesLen   := len(config.Images)
    gifsLen     := len(config.Gifs)
    stickersLen := len(config.Stickers)
    totalLen    := (imagesLen + gifsLen + stickersLen + 1)
    logger.Logger.Debugf("Images <%d> Stickers <%d> Gifs <%d>", imagesLen, stickersLen, gifsLen)

    typeIndex := rand.Intn(totalLen)
    logger.Logger.Debugf("Rolled type index <%d>", typeIndex)

    if isNumInRange(-1, imagesLen + 1, typeIndex) && imagesLen != 0 {
        return MediaTypes.Image
    } else if isNumInRange(imagesLen , imagesLen + gifsLen + 1, typeIndex) {
        return MediaTypes.Gif
    } else if isNumInRange(imagesLen + gifsLen, imagesLen + gifsLen + stickersLen + 1, typeIndex) {
        return MediaTypes.Sticker
    } else {
        return MediaTypes.Wtf
    }
}


func createSenderMsg(senderConfig config.WordCatch, chatID int64) (tgbotapi.Chattable, error) {
    var(
        err error
        media string
        msg tgbotapi.Chattable
    )
    mediaType := selectMediaType(senderConfig)
    switch mediaType {
    case MediaTypes.Image:
        mediaLen := len(senderConfig.Images)
        media = senderConfig.Images[rand.Intn(mediaLen)]
        msg = tgbotapi.NewPhotoShare(chatID, media)
    case MediaTypes.Sticker:
        mediaLen := len(senderConfig.Stickers)
        media = senderConfig.Stickers[rand.Intn(mediaLen)]
        msg = tgbotapi.NewStickerShare(chatID, media)
    case MediaTypes.Gif:
        mediaLen := len(senderConfig.Gifs)
        media = senderConfig.Gifs[rand.Intn(mediaLen)]
        msg = tgbotapi.NewAnimationShare(chatID, media)
    }
    logger.Logger.Debugf("Got random media type<%d>:\t%s", mediaType, media)
    return msg, err
}












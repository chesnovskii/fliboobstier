package bot

import (
	"fmt"
	"math/rand"

	"github.com/chesnovsky/fliboobstier/config"
	"github.com/chesnovsky/fliboobstier/logger"
	"github.com/chesnovsky/fliboobstier/storage"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

type MediaTypesType struct {
	None     int
	Image    int
	Sticker  int
	Gif      int
	Document int
}

var (
	MediaTypes = MediaTypesType{-1, 0, 1, 2, 3}
)

func numInRange(rMin int, rMax int, num int) bool {
	if rMin <= num && num <= rMax {
		return true
	} else {
		return false
	}
}

func selectMediaType(action_elements storage.RegexActionElements) int {
	imagesLen := len(action_elements.Images)
	gifsLen := len(action_elements.Gifs)
	stickersLen := len(action_elements.Stickers)
	totalLen := (imagesLen + gifsLen + stickersLen)
	logger.Logger.Debugf("Got media for to send:  <%d> images, <%d> stickers, <%d> gifs", imagesLen, stickersLen, gifsLen)
	if totalLen == 0 {
		// No media found, bye-bye
		return MediaTypes.None
	}

	typeIndex := (rand.Intn(totalLen) + 1)
	logger.Logger.Debugf("Randomed media index <%d>", typeIndex)

	// Selecting type of media based on count of each media type
	if numInRange(0, imagesLen, typeIndex) && imagesLen != 0 {
		return MediaTypes.Image
	} else if numInRange(imagesLen, imagesLen+gifsLen, typeIndex) && gifsLen != 0 {
		return MediaTypes.Gif
	} else if numInRange(imagesLen+gifsLen, imagesLen+gifsLen+stickersLen, typeIndex) && stickersLen != 0 {
		return MediaTypes.Sticker
	} else {
		return MediaTypes.None
	}
}

func (BotInstance *Bot) processRegexActions(message *tgbotapi.Message) {
	for catchType, catchConfig := range BotInstance.Config.RegexActions {
		logger.Logger.Debugf("Trying to catch <%s> regex: %s\n", catchType, catchConfig.RawRegex)
		matched := catchConfig.Regex.MatchString(message.Text)
		logger.Logger.Debugf("Regex is mached: %v\n", matched)
		if matched {
			msg, err := BotInstance.createSenderMsg(catchConfig, message.Chat.ID)
			if err == nil {
				logger.Logger.Debugf("Sending message: %v\n", msg)
				BotInstance.TgBot.Send(msg)
			} else {
				logger.Logger.Errorf("Cannot send an reply to message: %v\n%v", message, err)
			}
		}
	}
}

func (BotInstance *Bot) createSenderMsg(actionObj config.WordCatch, chatID int64) (tgbotapi.Chattable, error) {
	var (
		err   error
		media string
		msg   tgbotapi.Chattable
	)

	action_elements, err := BotInstance.Storage.GetRegexActionElements(actionObj.ID)
	mediaType := selectMediaType(action_elements)
	switch mediaType {
	case MediaTypes.Image:
		mediaLen := len(action_elements.Images)
		media = action_elements.Images[rand.Intn(mediaLen)]
		msg = tgbotapi.NewPhotoShare(chatID, media)
	case MediaTypes.Sticker:
		mediaLen := len(action_elements.Stickers)
		media = action_elements.Stickers[rand.Intn(mediaLen)]
		msg = tgbotapi.NewStickerShare(chatID, media)
	case MediaTypes.Gif:
		mediaLen := len(action_elements.Gifs)
		media = action_elements.Gifs[rand.Intn(mediaLen)]
		msg = tgbotapi.NewAnimationShare(chatID, media)
	case MediaTypes.None:
		err = fmt.Errorf("No media found for action <%s>", actionObj.ID)
	}
	logger.Logger.Debugf("Got random media type<%d>:\t%s", mediaType, media)
	return msg, err
}

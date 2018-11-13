package boobs

import (
	"github.com/chesnovsky/fliboobstier/logger"
	"github.com/chesnovsky/fliboobstier/writer"
	"gopkg.in/telegram-bot-api.v4"
)

// Showmeyourboobs show boobs
func Showmeyourboobs(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {

	// if message == nil {
	// 	return
	// }
	logger.Logger.Debugf("Recieved message from [%s] %s [ChatId: %s] %s", message.From.UserName, message.Text, message.Chat.ID, message.Photo)

	// writer.WritePhoto(message.Photo)
	writer.LogPhotoId(message.Photo)

	switch message.Text {
	case "сиськи":
		pic := "AgADAgAD0KkxG2BUMEtjAgH_n4NbopXZtw4ABHuQO-dPEy-VtVEGAAEC"
		// msg := tgbotapi.NewMessage(message.Chat.ID, "Привет, я покажу тебе сиськи, обязательно покажу, но потом, не сейчас, а сейчас хочу в пляс!!!")
		msg := tgbotapi.NewPhotoShare(message.Chat.ID, pic)
		logger.Logger.Debug(msg)
		bot.Send(msg)
	default:
		msg := tgbotapi.NewDocumentShare(message.Chat.ID, "CAADAgADnQUAAlOx9wMjvcls38LyPwI")
		logger.Logger.Debug(msg)
		bot.Send(msg)
	}

}

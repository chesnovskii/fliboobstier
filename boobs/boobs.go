package boobs

import (
	"gopkg.in/telegram-bot-api.v4"
)

func Showmeyourboobs(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {

	if message == nil {
		return
	}
	switch message.Text {
	case "сиськи":
		msg := tgbotapi.NewMessage(message.Chat.ID, "Привет, я покажу тебе сиськи, обязательно покажу, но потом, не сейчас, а сейчас хочу в пляс!!!")
		bot.Send(msg)
	default:
		// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		msg := tgbotapi.NewDocumentShare(message.Chat.ID, "CAADAgADnQUAAlOx9wMjvcls38LyPwI")
		bot.Send(msg)
	}

}

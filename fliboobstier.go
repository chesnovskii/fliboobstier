package main

import (
	"log"
	"os"

	// "fliboobstier/tg_bot"

	"gopkg.in/telegram-bot-api.v4"
)

func main() {

	tgToken := os.Getenv("TG_TOKEN")

	// tg_bot.Init()
	bot, err := tgbotapi.NewBotAPI(tgToken)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		log.Panic(err)
	}
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.Text == "" {
			log.Print("It's a penis!")
		}
		switch update.Message.Text {
		case "сиськи":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, я покажу тебе сиськи, обязательно покажу, но потом, не сейчас, а сейчас хочу в пляс!!!")
			bot.Send(msg)
		default:
			// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			msg := tgbotapi.NewDocumentShare(update.Message.Chat.ID, "CAADAgADnQUAAlOx9wMjvcls38LyPwI")
			bot.Send(msg)

		}
	}
}

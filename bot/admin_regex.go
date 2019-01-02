package bot

import (
	"fmt"
	"strings"

	"github.com/chesnovsky/fliboobstier/logger"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

func getElementTypeAndId(message *tgbotapi.Message) (string, string) {
	if message.Animation != nil {
		return "gif", message.Animation.FileID
	}
	if message.Photo != nil && len(*message.Photo) != 0 {
		return "image", (*message.Photo)[len(*message.Photo)-1].FileID
	}
	if message.Sticker != nil {
		return "sticker", message.Sticker.FileID
	}
	return "", ""
}

func (BotInstance *Bot) listRegexActions(message *tgbotapi.Message) {
	action_strings := []string{}
	for action_name, action_data := range BotInstance.Config.RegexActions {
		action_string := fmt.Sprintf("*%s*: `%s`", action_name, action_data.RawRegex)
		action_strings = append(action_strings, action_string)
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, strings.Join(action_strings, "\n"))
	msg.ParseMode = "Markdown"
	logger.Logger.Debugf("Sending message: %+v\n", msg)
	BotInstance.TgBot.Send(msg)
}

func (BotInstance *Bot) addRegexActionElement(message *tgbotapi.Message) {
	if !BotInstance.checkCommandArgs(message, 1) {
		return
	}
	regex_action_name := strings.Split(message.CommandArguments(), " ")[0]
	if !BotInstance.checkRegexActionExists(regex_action_name, message.Chat.ID) {
		return
	}
	msg_text := "Ok, send me an Image/Sticker/Gif"
	msg := tgbotapi.NewMessage(message.Chat.ID, msg_text)
	BotInstance.TgBot.Send(msg)
	transition_data := map[string]string{
		"regex_action_name": regex_action_name,
	}
	transition := Transition{
		Name:   "AddRegexActionElement",
		UserID: message.From.ID,
		ChatID: message.Chat.ID,
		Data:   transition_data,
	}
	BotInstance.addTransition(transition)
}

func (BotInstance *Bot) continueAddRegexActionElement(message *tgbotapi.Message, transition Transition) {
	element_type, element_id := getElementTypeAndId(message)
	if element_type == "" {
		msg_text := "You must send Image/Gif/Sticker!"
		msg := tgbotapi.NewMessage(message.Chat.ID, msg_text)
		BotInstance.TgBot.Send(msg)
		BotInstance.removeTransition(transition)
	}
	action_id := transition.Data["regex_action_name"]
	BotInstance.Storage.AddRegexActionElement(action_id, element_type, element_id)
	BotInstance.removeTransition(transition)
}

func (BotInstance *Bot) removeRegexActionElement(message *tgbotapi.Message) {
	if !BotInstance.checkCommandArgs(message, 1) {
		return
	}
	regex_action_name := strings.Split(message.CommandArguments(), " ")[0]
	if !BotInstance.checkRegexActionExists(regex_action_name, message.Chat.ID) {
		return
	}
	msg_text := "Ok, send me an Image/Sticker/Gif"
	msg := tgbotapi.NewMessage(message.Chat.ID, msg_text)
	BotInstance.TgBot.Send(msg)
	transition_data := map[string]string{
		"regex_action_name": regex_action_name,
	}
	transition := Transition{
		Name:   "RemoveRegexActionElement",
		UserID: message.From.ID,
		ChatID: message.Chat.ID,
		Data:   transition_data,
	}
	BotInstance.addTransition(transition)
}

func (BotInstance *Bot) continueRemoveRegexActionElement(message *tgbotapi.Message, transition Transition) {
	element_type, element_id := getElementTypeAndId(message)
	if element_type == "" {
		msg_text := "You must send Image/Gif/Sticker!"
		msg := tgbotapi.NewMessage(message.Chat.ID, msg_text)
		BotInstance.TgBot.Send(msg)
		BotInstance.removeTransition(transition)
	}
	action_id := transition.Data["regex_action_name"]
	BotInstance.Storage.RemoveRegexActionElement(action_id, element_type, element_id)
	BotInstance.removeTransition(transition)
}

func (BotInstance *Bot) checkRegexActionExists(action string, chatID int64) bool {
	for action_name := range BotInstance.Config.RegexActions {
		if action_name == action {
			return true
		}
	}
	msg_text := fmt.Sprintf("Cannot find regex action <%s>", action)
	msg := tgbotapi.NewMessage(chatID, msg_text)
	BotInstance.TgBot.Send(msg)
	return false
}

func (BotInstance *Bot) showRegexAction(message *tgbotapi.Message) {
	if !BotInstance.checkCommandArgs(message, 1) {
		return
	}
	regex_action_name := strings.Split(message.CommandArguments(), " ")[0]
	if !BotInstance.checkRegexActionExists(regex_action_name, message.Chat.ID) {
		return
	}
	// Sending all the medias for action
	elements, _ := BotInstance.Storage.GetRegexActionElements(regex_action_name)
	for _, media := range elements.Images {
		msg := tgbotapi.NewPhotoShare(message.Chat.ID, media)
		BotInstance.TgBot.Send(msg)
	}
	for _, media := range elements.Gifs {
		msg := tgbotapi.NewAnimationShare(message.Chat.ID, media)
		BotInstance.TgBot.Send(msg)
	}
	for _, media := range elements.Stickers {
		msg := tgbotapi.NewStickerShare(message.Chat.ID, media)
		BotInstance.TgBot.Send(msg)
	}
}

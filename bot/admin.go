package bot

import (
	"fmt"
	"strings"

	"github.com/chesnovsky/fliboobstier/logger"
	tgbotapi "gopkg.in/telegram-bot-api.v5"
)

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

var AdminCommands = []string{"list_admins", "list_regex_actions", "show_regex_action", "add_regex_action_element", "remove_regex_action_element"}

func (BotInstance *Bot) processAdminCommands(message *tgbotapi.Message) {
	if !(contains(AdminCommands, message.Command())) {
		return
	}
	if !BotInstance.isUserAdmin(message.From.UserName) {
		logger.Logger.Infof("User @<%s> doesn't have admin privileges for command <%s>", message.From.UserName, message.Command())
		msg_text := fmt.Sprintf("Звуки высокомерного игнорирования @%s", message.From.UserName)
		msg := tgbotapi.NewMessage(message.Chat.ID, msg_text)
		logger.Logger.Debugf("Sending message: %+v\n", msg)
		BotInstance.TgBot.Send(msg)
		return
	}

	// Inline description for bot, set it via @BotFather:
	// list_admins					- List admin users
	// list_regex_actions			- List actions based on message regex
	// show_regex_action			- Show all media for regex action
	// add_regex_action_element 	- Add new media to regex action
	// remove_regex_action_element	- Remove media from regex action
	logger.Logger.Debugf("Processing admin command <%s>", message.Text)
	var err error
	switch message.Command() {
	case "list_admins":
		err = BotInstance.listAdmins(message)
	case "list_regex_actions":
		err = BotInstance.listRegexActions(message)
	case "show_regex_action":
		err = BotInstance.showRegexAction(message)
	case "add_regex_action_element":
		err = BotInstance.addRegexActionElement(message)
	case "remove_regex_action_element":
		err = BotInstance.removeRegexActionElement(message)
	}
	if err != nil {
		BotInstance.CritError(message.Chat.ID, err)
	}
}

func (BotInstance *Bot) isUserAdmin(user string) bool {
	is_admin := false
	admin_list, _ := BotInstance.Storage.GetAdminListSqlite()
	if contains(admin_list, user) {
		is_admin = true
	}
	return is_admin
}

func (BotInstance *Bot) listAdmins(message *tgbotapi.Message) error {
	admin_list := []string{}
	admin_list, err := BotInstance.Storage.GetAdminListSqlite()
	if err != nil {
		return err
	}
	logger.Logger.Debugf("Got admins list: %s", admin_list)
	msg := tgbotapi.NewMessage(message.Chat.ID, strings.Join(admin_list, "\n"))
	logger.Logger.Debugf("Sending message: %+v\n", msg)
	BotInstance.TgBot.Send(msg)
	return nil
}

func (BotInstance *Bot) checkCommandArgs(message *tgbotapi.Message, required_args_num int) bool {
	command_args := strings.Split(message.CommandArguments(), " ")
	if len(command_args) == required_args_num {
		return true
	}
	msg_text := fmt.Sprintf("Need <%d> arguments to command, got <%d>", required_args_num, len(command_args))
	msg := tgbotapi.NewMessage(message.Chat.ID, msg_text)
	BotInstance.TgBot.Send(msg)
	return false
}

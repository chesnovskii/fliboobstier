package bot

import (
	"reflect"

	"github.com/chesnovsky/fliboobstier/logger"
	tgbotapi "gopkg.in/telegram-bot-api.v5"
)

func (BotInstance *Bot) addTransition(src_transition Transition) {
	logger.Logger.Debugf("Adding transition: %+v", src_transition)
	BotInstance.Transitions = append(BotInstance.Transitions, src_transition)
}

func (BotInstance *Bot) removeTransition(src_transition Transition) {
	for index, transition := range BotInstance.Transitions {
		if reflect.DeepEqual(src_transition, transition) {
			logger.Logger.Debugf("Removing transition: %+v", src_transition)
			BotInstance.Transitions = append(BotInstance.Transitions[:index], BotInstance.Transitions[index+1:]...)
		}
	}
}

func (BotInstance *Bot) getTransitions() []Transition {
	return BotInstance.Transitions
}

func (BotInstance *Bot) processTransitions(message *tgbotapi.Message) bool {
	for _, transition := range BotInstance.getTransitions() {
		if transition.ChatID == message.Chat.ID && transition.UserID == message.From.ID {
			switch transition.Name {
			case "AddRegexActionElement":
				logger.Logger.Debugf("Continuing transition <%s>", transition.Name)
				BotInstance.continueAddRegexActionElement(message, transition)
			case "RemoveRegexActionElement":
				logger.Logger.Debugf("Continuing transition <%s>", transition.Name)
				BotInstance.continueRemoveRegexActionElement(message, transition)
			}
			return true
		}
	}
	return false
}

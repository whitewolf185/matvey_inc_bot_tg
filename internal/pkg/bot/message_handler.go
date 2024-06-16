package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type messageHandler struct {
	messagesChan tgbotapi.UpdatesChannel
}

func NewMessageHandler(tgApi *tgbotapi.BotAPI) messageHandler {
	messagesChan := tgApi.GetUpdatesChan(tgbotapi.NewUpdate(0))
	return messageHandler{
		messagesChan: messagesChan,
	}
}

func (m *messageHandler) HandleMessages() {
	logrus.Infoln("starting to handling messages...")
	for update := range m.messagesChan {
		if update.Message == nil { // ignore messages
			continue
		}
		logrus.Debugln(update.Message.Chat.ID)

	}
}

// func (m *messageHandler) getAllUsers(update tgbotapi.Update) ([]string, error) {
// 	update.Message.Chat.
// }

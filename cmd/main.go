package main

import (
	"github.com/sirupsen/logrus"
	"github.com/whitewolf185/matvey_inc_bot_tg/internal/config"
	"github.com/whitewolf185/matvey_inc_bot_tg/internal/config/flags"
	"github.com/whitewolf185/matvey_inc_bot_tg/internal/pkg/bot"
)

func main() {
	flags.InitServiceFlags()
	logrus.SetLevel(logrus.DebugLevel)
	tgApi, err := config.NewTgBot()
	if err != nil {
		logrus.Fatalln(err.Error())
	}

	messageHandler := bot.NewMessageHandler(tgApi)
	messageHandler.HandleMessages()
}

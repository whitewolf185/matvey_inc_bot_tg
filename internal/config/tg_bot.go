package config

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewTgBot() (*tgbotapi.BotAPI, error) {
	currApi, err := tgbotapi.NewBotAPI(GetValue(BotApiToken))
	if err != nil {
		return nil, fmt.Errorf("could not create bot: %w", err)
	}
	return currApi, nil
}

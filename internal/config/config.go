package config

import (
	"os"
)

type configName string

const (
	BotApiToken = configName("api_token")
)

func GetValue(cnfg configName) string {
	return os.Getenv(string(cnfg))
}

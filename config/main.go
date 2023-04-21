package config

import (
	"github.com/mattermost/mattermost-server/v5/plugin"
	"go.uber.org/atomic"
)

var (
	config     atomic.Value
	Mattermost plugin.API
	BotUserID  string
)

type Configuration struct {
	Secret string `json:"secret"`
}

func GetConfig() *Configuration {
	return config.Load().(*Configuration)
}

func SetConfig(c *Configuration) {
	config.Store(c)
}

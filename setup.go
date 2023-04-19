package main

import (
	"io/ioutil"
	"path/filepath"

	"github.com/lgu-it-sre/mm-github-pr/config"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/pkg/errors"
)

const (
	botUserName    = "github.bot"
	botDisplayName = "GitHub Bot"
	botDescription = "Created by @juny"
)

func (p *Plugin) setUpBotUser() error {
	botUserID, err := p.Helpers.EnsureBot(&model.Bot{
		Username:    botUserName,
		DisplayName: botDisplayName,
		Description: botDescription,
	})
	if err != nil {
		return err
	}

	bundlePath, err := p.API.GetBundlePath()
	if err != nil {
		return err
	}

	profileImage, err := ioutil.ReadFile(filepath.Join(bundlePath, "assets", "icon.png"))
	if err != nil {
		return err
	}

	if appErr := p.API.SetProfileImage(botUserID, profileImage); appErr != nil {
		return errors.Wrap(appErr, "couldn't set profile image")
	}

	config.BotUserID = botUserID
	return nil
}

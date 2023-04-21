package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/lgu-it-sre/mm-github-pr/config"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

type Plugin struct {
	plugin.MattermostPlugin
}

func (p *Plugin) OnActivate() error {
	config.Mattermost = p.API

	if err := p.setUpBotUser(); err != nil {
		config.Mattermost.LogError("Failed to create a bot user", "Error", err.Error())
		return err
	}

	if err := p.OnConfigurationChange(); err != nil {
		return err
	}

	return nil
}

func (p *Plugin) OnConfigurationChange() error {
	if config.Mattermost == nil {
		return nil
	}
	var configuration config.Configuration

	if err := config.Mattermost.LoadPluginConfiguration(&configuration); err != nil {
		config.Mattermost.LogError("Failed to load plugin configuration", "Error", err.Error())
		return err
	}

	config.SetConfig(&configuration)
	return nil
}

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	switch event := r.Header.Get("X-GitHub-Event"); event {
	case "pull_request":
		p.handlePullRequestEvent(w, r)
	}
}

func (p *Plugin) createPost(client *model.Client4, username, message, title, title_link, text string) {
	user, res := client.GetUserByUsername(username, "")
	if res.StatusCode >= 400 {
		fmt.Println(res.Error.Message)
		return
	}

	channel, res := client.CreateDirectChannel(MMBOTID, user.Id)
	if res.StatusCode >= 400 {
		fmt.Println(res.Error.Message)
		return
	}

	attachment := &model.SlackAttachment{
		Fallback:  "",
		Color:     "#db3b21",
		Title:     title,
		TitleLink: title_link,
		Text:      text,
	}
	post := &model.Post{
		UserId:    MMBOTID,
		ChannelId: channel.Id,
		Message:   message,
	}

	model.ParseSlackAttachment(post, []*model.SlackAttachment{attachment})

	client.CreatePost(post)
}

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	plugin.ClientMain(&Plugin{})
}

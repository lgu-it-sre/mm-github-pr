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
	case "pull_request_review":
		p.handlePullRequestReviewEvent(w, r)
	case "issue_comment":
		p.handleIssueCommentEvent(w, r)
	}
}

func (p *Plugin) createPost(username, message, title, title_link, text string) {
	user, err := config.Mattermost.GetUserByUsername(username)
	if err != nil {
		fmt.Printf("[HTTP %d] Failed to get user %s: %s", err.StatusCode, username, err.Message)
		return
	}

	channel, err := config.Mattermost.GetDirectChannel(config.BotUserID, user.Id)
	if err != nil {
		fmt.Printf("[HTTP %d] Failed to get user %s: %s", err.StatusCode, username, err.Message)
		return
	}

	attachment := &model.SlackAttachment{
		Fallback:  "",
		Color:     "#1a7f37",
		Title:     title,
		TitleLink: title_link,
		Text:      text,
	}
	post := &model.Post{
		UserId:    config.BotUserID,
		ChannelId: channel.Id,
		Message:   message,
	}

	model.ParseSlackAttachment(post, []*model.SlackAttachment{attachment})

	config.Mattermost.CreatePost(post)
}

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	plugin.ClientMain(&Plugin{})
}

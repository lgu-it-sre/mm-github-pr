package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

type Plugin struct {
	plugin.MattermostPlugin
}

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	return
	// }

	client := model.NewAPIv4Client(MMDOMAIN)
	client.SetToken(MMTOKEN)
}

func createPost(client *model.Client4, username, message, title, title_link, text string) {
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

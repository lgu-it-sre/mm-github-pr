package main

import (
	"errors"
	"net/http"

	"github.com/lgu-it-sre/mm-github-pr/github"
	"github.com/lgu-it-sre/mm-github-pr/util"
	"github.com/mattermost/mattermost-server/v5/model"
)

func (p *Plugin) handlePullRequestEvent(w http.ResponseWriter, r *http.Request) {
	var pr github.PullRequestEvent
	err := util.DecodeJSONBody(w, r, &pr)
	if err != nil {
		var mr *util.MalformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.Msg, mr.Status)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	client := model.NewAPIv4Client(MMDOMAIN)
	client.SetToken(MMTOKEN)

	p.createPost(client, "admin", "message", "title", pr.PullRequest.HtmlURL, "description")
}

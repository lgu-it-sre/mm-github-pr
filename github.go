package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/lgu-it-sre/mm-github-pr/config"
	"github.com/lgu-it-sre/mm-github-pr/github"
	"github.com/lgu-it-sre/mm-github-pr/util"
)

func (p *Plugin) handlePullRequestEvent(w http.ResponseWriter, r *http.Request) {
	var pr github.PullRequestEvent
	if err := util.DecodeJSONBody(w, r, &pr); err != nil {
		var mr *util.MalformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.Msg, mr.Status)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	var sender_nickname string
	sender_username := pr.Sender.Login
	pr_action := pr.Action
	pr_title := pr.PullRequest.Title
	pr_url := pr.PullRequest.HtmlURL
	pr_repo_fullname := pr.PullRequest.Head.Repo.FullName
	pr_repo_html := pr.PullRequest.Head.Repo.HTML
	pr_body := pr.PullRequest.Body
	user, err := config.Mattermost.GetUserByUsername(sender_username)
	if err != nil || user.Nickname == "" {
		sender_nickname = sender_username
	} else {
		sender_nickname = user.Nickname
	}

	payload := fmt.Sprintf(
		"%s (%s) %s merge request [%s](%s) in [%s](%s)",
		sender_nickname,
		sender_username,
		pr_action,
		pr_title,
		pr_url,
		pr_repo_fullname,
		pr_repo_html)
	for _, username := range pr.GetReviewers() {
		p.createPost(username, payload, pr_title, pr_url, pr_body)
	}
}

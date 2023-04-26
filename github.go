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

func (p *Plugin) handleIssueCommentEvent(w http.ResponseWriter, r *http.Request) {
	var ic github.IssueCommentEvent
	if err := util.DecodeJSONBody(w, r, &ic); err != nil {
		var mr *util.MalformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.Msg, mr.Status)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	var sender_nickname string
	sender_username := ic.Sender.Login
	ic_username := ic.Issue.User.Login
	if sender_username == ic_username {
		http.Error(w, "OK", http.StatusOK)
		return
	}

	ic_action := ic.Action
	ic_title := ic.Issue.Title
	ic_url := ic.Issue.HtmlURL
	ic_repo_fullname := ic.Repo.FullName
	ic_repo_html := ic.Repo.HTML
	ic_body := ic.Comment.Body
	user, err := config.Mattermost.GetUserByUsername(sender_username)
	if err != nil || user.Nickname == "" {
		sender_nickname = sender_username
	} else {
		sender_nickname = user.Nickname
	}

	payload := fmt.Sprintf(
		"%s (%s) %s comment to [%s](%s) in [%s](%s)",
		sender_nickname,
		sender_username,
		ic_action,
		ic_title,
		ic_url,
		ic_repo_fullname,
		ic_repo_html)
	p.createPost(ic.Issue.User.Login, payload, ic_title, ic_url, ic_body)
}

package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/lgu-it-sre/mm-github-pr/github"
	"github.com/lgu-it-sre/mm-github-pr/util"
)

var (
	admin = ""
)

func (p *Plugin) handlePullRequestEvent(w http.ResponseWriter, r *http.Request) {
	var pr github.PullRequestEvent
	if err := util.DecodeJSONBody(w, r, &pr); err != nil {
		var mr *util.MalformedRequest
		if errors.As(err, &mr) {
			p.createPost(admin, "Pull Request Error", fmt.Sprint(mr.Status), "", mr.Msg)
			http.Error(w, mr.Msg, mr.Status)
		} else {
			p.createPost(admin, "Pull Request Error (Internal Server Error)", fmt.Sprint(http.StatusInternalServerError), "", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	sender_username := pr.Sender.Login
	pr_action := pr.Action
	pr_title := pr.PullRequest.Title
	pr_url := pr.PullRequest.HtmlURL
	pr_repo_fullname := pr.PullRequest.Head.Repo.FullName
	pr_repo_html := pr.PullRequest.Head.Repo.HtmlURL
	pr_body := pr.PullRequest.Body

	payload := fmt.Sprintf(
		"@%s (%s) %s pull request [%s](%s) in [%s](%s)",
		sender_username,
		sender_username,
		pr_action,
		pr_title,
		pr_url,
		pr_repo_fullname,
		pr_repo_html)
	for _, username := range pr.GetReviewers() {
		p.createPost(username, payload, pr_title, pr_url, pr_body)
	}
	p.createPost(admin, payload, pr_title, pr_url, pr_body)
}

func (p *Plugin) handlePullRequestReviewEvent(w http.ResponseWriter, r *http.Request) {
	var pr github.PullRequestReviewEvent
	if err := util.DecodeJSONBody(w, r, &pr); err != nil {
		var mr *util.MalformedRequest
		if errors.As(err, &mr) {
			p.createPost(admin, "Pull Request Review Error", fmt.Sprint(mr.Status), "", mr.Msg)
			http.Error(w, mr.Msg, mr.Status)
		} else {
			p.createPost(admin, "Pull Request Review Error (Internal Server Error)", fmt.Sprint(http.StatusInternalServerError), "", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	sender_username := pr.Sender.Login
	pr_action := pr.Review.State
	pr_title := pr.PullRequest.Title
	pr_url := pr.Review.HtmlURL
	pr_repo_fullname := pr.PullRequest.Head.Repo.FullName
	pr_repo_html := pr.PullRequest.Head.Repo.HtmlURL
	pr_body := pr.Review.Body

	payload := fmt.Sprintf(
		"@%s (%s) %s pull request [%s](%s) in [%s](%s)",
		sender_username,
		sender_username,
		pr_action,
		pr_title,
		pr_url,
		pr_repo_fullname,
		pr_repo_html)
	p.createPost(pr.PullRequest.User.Login, payload, pr_title, pr_url, pr_body)
	for _, username := range pr.GetReviewers() {
		p.createPost(username, payload, pr_title, pr_url, pr_body)
	}
	p.createPost(admin, payload, pr_title, pr_url, pr_body)
}

func (p *Plugin) handleIssueCommentEvent(w http.ResponseWriter, r *http.Request) {
	var ic github.IssueCommentEvent
	if err := util.DecodeJSONBody(w, r, &ic); err != nil {
		var mr *util.MalformedRequest
		if errors.As(err, &mr) {
			p.createPost(admin, "Issue Comment Error", fmt.Sprint(mr.Status), "", mr.Msg)
			http.Error(w, mr.Msg, mr.Status)
		} else {
			p.createPost(admin, "Issue Comment Error (Internal Server Error)", fmt.Sprint(http.StatusInternalServerError), "", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

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
	ic_repo_html := ic.Repo.HtmlURL
	ic_body := ic.Comment.Body

	payload := fmt.Sprintf(
		"@%s (%s) %s comment to [%s](%s) in [%s](%s)",
		sender_username,
		sender_username,
		ic_action,
		ic_title,
		ic_url,
		ic_repo_fullname,
		ic_repo_html)
	p.createPost(ic_username, payload, ic_title, ic_url, ic_body)
	p.createPost(admin, payload, ic_title, ic_url, ic_body)
}

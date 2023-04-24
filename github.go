package main

import (
	"errors"
	"net/http"

	"github.com/lgu-it-sre/mm-github-pr/github"
	"github.com/lgu-it-sre/mm-github-pr/util"
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

	payload := "이름" + ` (` + pr.Sender.Login + `) ` + pr.Action + ` merge request ` + `[` + pr.PullRequest.Title + `](` + pr.PullRequest.HtmlURL + `) in [` + pr.PullRequest.Head.Repo.FullName + `](` + pr.PullRequest.Head.Repo.HTML + `)`
	// for _, username := range pr.GetReviewers() {
	// 	p.createPost(username, "message", "title", pr.PullRequest.HtmlURL, "description")
	// }
	p.createPost("admin", payload, pr.PullRequest.Title, pr.PullRequest.HtmlURL, pr.PullRequest.Body)
}

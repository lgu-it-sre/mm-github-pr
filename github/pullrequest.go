package github

type PullRequestEvent struct {
	Action      string      `json:"action"`
	Number      int         `json:"number"`
	PullRequest pullRequest `json:"pull_request"`
	Reviewer    user        `json:"requested_reviewer"`
	Assignee    user        `json:"assignee"`
	Label       label       `json:"label"`
	Changes     struct {
		Title struct {
			From string `json:"from"`
		} `json:"title"`
		Body struct {
			From string `json:"from"`
		} `json:"body"`
	} `json:"changes"`
	Before     string       `json:"before"`
	After      string       `json:"after"`
	Repository repo         `json:"repository"`
	Org        organization `json:"organization"`
	Ent        enterprise   `json:"enterprise"`
	Sender     user         `json:"sender"`
}

func (p *PullRequestEvent) GetReviewers() []string {
	var reviewers []string
	for _, reviewer := range p.PullRequest.Reviewers {
		reviewers = append(reviewers, reviewer.Login)
	}
	return reviewers
}

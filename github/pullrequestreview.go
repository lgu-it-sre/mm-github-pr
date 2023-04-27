package github

type PullRequestReviewEvent struct {
	Action      string       `json:"action"`
	Review      review       `json:"review"`
	PullRequest pullRequest  `json:"pull_request"`
	Repository  repo         `json:"repository"`
	Org         organization `json:"organization"`
	Ent         enterprise   `json:"enterprise"`
	Sender      user         `json:"sender"`
}

func (p *PullRequestReviewEvent) GetReviewers() []string {
	var reviewers []string
	for _, reviewer := range p.PullRequest.Reviewers {
		reviewers = append(reviewers, reviewer.Login)
	}
	return reviewers
}

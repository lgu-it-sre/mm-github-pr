package github

type IssueCommentEvent struct {
	Action  string `json:"action"`
	Changes struct {
		Body struct {
			From string `json:"from"`
		} `json:"body"`
	} `json:"changes"`
	Issue   issue        `json:"issue"`
	Comment comment      `json:"comment"`
	Repo    repo         `json:"repository"`
	Org     organization `json:"organization"`
	Ent     enterprise   `json:"enterprise"`
	Sender  user         `json:"sender"`
}

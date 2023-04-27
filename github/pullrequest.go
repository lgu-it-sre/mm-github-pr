package github

type PullRequestEvent struct {
	Action      string `json:"action"`
	Number      int    `json:"number"`
	PullRequest struct {
		URL            string `json:"url"`
		ID             int    `json:"id"`
		NodeID         string `json:"node_id"`
		HtmlURL        string `json:"html_url"`
		DiffURL        string `json:"diff_url"`
		PatchURL       string `json:"patch_url"`
		IssueURL       string `json:"issue_url"`
		Number         int    `json:"number"`
		State          string `json:"state"`
		Locked         bool   `json:"locked"`
		Title          string `json:"title"`
		User           user   `json:"user"`
		Body           string `json:"body"`
		CreatedAt      string `json:"created_at"`
		UpdatedAt      string `json:"updated_at"`
		ClosedAt       string `json:"closed_at"`
		MergedAt       string `json:"merged_at"`
		MergeCommitSha string `json:"merge_commit_sha"`
		Assignee       user   `json:"assignee"`
		Assignees      []user `json:"assignees"`
		Reviewers      []user `json:"requested_reviewers"`
		Team           []team `json:"requested_teams"`
		Label          []struct {
			ID    int    `json:"id"`
			Node  string `json:"node_id"`
			URL   string `json:"url"`
			Name  string `json:"name"`
			Desc  string `json:"description"`
			Color string `json:"color"`
			Def   bool   `json:"default"`
		} `json:"labels"`
		Milestone   milestone `json:"milestone"`
		Draft       bool      `json:"draft"`
		CommitsURL  string    `json:"commits_url"`
		ReviewsURL  string    `json:"review_comments_url"`
		ReviewURL   string    `json:"review_comment_url"`
		CommentsURL string    `json:"comments_url"`
		Statuses    string    `json:"statuses_url"`
		Head        struct {
			Label string `json:"label"`
			Ref   string `json:"ref"`
			SHA   string `json:"sha"`
			User  user   `json:"user"`
			Repo  repo   `json:"repo"`
		} `json:"head"`
		Base struct {
			Label string `json:"label"`
			Ref   string `json:"ref"`
			SHA   string `json:"sha"`
			User  user   `json:"user"`
			Repo  repo   `json:"repo"`
		} `json:"base"`
		Links struct {
			Self     href `json:"self"`
			Html     href `json:"html"`
			Issue    href `json:"issue"`
			Comments href `json:"comments"`
			Review   href `json:"review_comments"`
			Review2  href `json:"review_comment"`
			Commits  href `json:"commits"`
			Status   href `json:"statuses"`
		} `json:"_links"`
		AuthorAssociation   string `json:"author_association"`
		AutoMerge           bool   `json:"auto_merge"`
		ActiveLockReason    string `json:"active_lock_reason"`
		Merged              bool   `json:"merged"`
		Mergable            bool   `json:"mergeable"`
		Rebaseable          bool   `json:"rebaseable"`
		MergeableState      string `json:"mergeable_state"`
		MergedBy            user   `json:"merged_by"`
		Comments            int    `json:"comments"`
		ReviewComments      int    `json:"review_comments"`
		MaintainerCanModify bool   `json:"maintainer_can_modify"`
		Commits             int    `json:"commits"`
		Additions           int    `json:"additions"`
		Deletions           int    `json:"deletions"`
		ChangedFiles        int    `json:"changed_files"`
	} `json:"pull_request"`
	Reviewer user `json:"requested_reviewer"`
	Changes  struct {
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

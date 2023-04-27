package github

type user struct {
	Login   string `json:"login"`
	Id      int    `json:"id"`
	NodeID  string `json:"node_id"`
	Avatar  string `json:"avatar_url"`
	Grav    string `json:"gravatar_id"`
	URL     string `json:"url"`
	HtmlURL string `json:"html_url"`
	Follow  string `json:"followers_url"`
	Foll    string `json:"following_url"`
	Gists   string `json:"gists_url"`
	Star    string `json:"starred_url"`
	Subs    string `json:"subscriptions_url"`
	Orgs    string `json:"organizations_url"`
	Repos   string `json:"repos_url"`
	Events  string `json:"events_url"`
	Rcvd    string `json:"received_events_url"`
	Type    string `json:"type"`
	Site    bool   `json:"site_admin"`
}

type team struct {
	Id           int    `json:"id"`
	NodeID       string `json:"node_id"`
	URL          string `json:"url"`
	HtmlURL      string `json:"html_url"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	Desc         string `json:"description"`
	Privacy      string `json:"privacy"`
	Notification string `json:"notification_setting"`
	Perm         string `json:"permission"`
	Mem          string `json:"members_url"`
	Repo         string `json:"repositories_url"`
	Parent       string `json:"parent"`
}

type repo struct {
	Id              int    `json:"id"`
	NodeID          string `json:"node_id"`
	Name            string `json:"name"`
	FullName        string `json:"full_name"`
	Private         bool   `json:"private"`
	Owner           user   `json:"owner"`
	HtmlURL         string `json:"html_url"`
	Desc            string `json:"description"`
	Fork            bool   `json:"fork"`
	URL             string `json:"url"`
	ForksURL        string `json:"forks_url"`
	Keys            string `json:"keys_url"`
	Collab          string `json:"collaborators_url"`
	Teams           string `json:"teams_url"`
	Hooks           string `json:"hooks_url"`
	Issue           string `json:"issue_events_url"`
	Events          string `json:"events_url"`
	Assign          string `json:"assignees_url"`
	Branch          string `json:"branches_url"`
	Tags            string `json:"tags_url"`
	Blobs           string `json:"blobs_url"`
	GitTags         string `json:"git_tags_url"`
	GitRefs         string `json:"git_refs_url"`
	Trees           string `json:"trees_url"`
	Statuses        string `json:"statuses_url"`
	Langs           string `json:"languages_url"`
	Starg           string `json:"stargazers_url"`
	Contrib         string `json:"contributors_url"`
	Subs            string `json:"subscribers_url"`
	SubsC           string `json:"subscription_url"`
	Commits         string `json:"commits_url"`
	GitComm         string `json:"git_commits_url"`
	Comments        string `json:"comments_url"`
	IssueC          string `json:"issue_comment_url"`
	Contents        string `json:"contents_url"`
	Compare         string `json:"compare_url"`
	Merges          string `json:"merges_url"`
	Archive         string `json:"archive_url"`
	Downloads       string `json:"downloads_url"`
	Issues          string `json:"issues_url"`
	Pulls           string `json:"pulls_url"`
	Milestones      string `json:"milestones_url"`
	Notifications   string `json:"notifications_url"`
	Labels          string `json:"labels_url"`
	Releases        string `json:"releases_url"`
	Deploy          string `json:"deployments_url"`
	Created         string `json:"created_at"`
	Updated         string `json:"updated_at"`
	Pushed          string `json:"pushed_at"`
	GitURL          string `json:"git_url"`
	SshURL          string `json:"ssh_url"`
	CloneURL        string `json:"clone_url"`
	SvnURL          string `json:"svn_url"`
	Homepage        string `json:"homepage"`
	Size            int    `json:"size"`
	StargazersCount int    `json:"stargazers_count"`
	WatchersCount   int    `json:"watchers_count"`
	Language        string `json:"language"`
	HasIssues       bool   `json:"has_issues"`
	HasProjects     bool   `json:"has_projects"`
	HasDownloads    bool   `json:"has_downloads"`
	HasWiki         bool   `json:"has_wiki"`
	HasPages        bool   `json:"has_pages"`
	HasDiscussions  bool   `json:"has_discussions"`
	ForksCount      int    `json:"forks_count"`
	MirrorURL       string `json:"mirror_url"`
	Archived        bool   `json:"archived"`
	Disabled        bool   `json:"disabled"`
	OpenIssuesCount int    `json:"open_issues_count"`
	License         struct {
		Key    string `json:"key"`
		Name   string `json:"name"`
		SpdxID string `json:"spdx_id"`
		URL    string `json:"url"`
		NodeID string `json:"node_id"`
	} `json:"license"`
	AllowForking              bool     `json:"allow_forking"`
	IsTemplate                bool     `json:"is_template"`
	WebCommitSignoffRequired  bool     `json:"web_commit_signoff_required"`
	Topics                    []string `json:"topics"`
	Visibility                string   `json:"visibility"`
	Forks                     int      `json:"forks"`
	OpenIssues                int      `json:"open_issues"`
	Watchers                  int      `json:"watchers"`
	DefaultBranch             string   `json:"default_branch"`
	AllowSquashMerge          bool     `json:"allow_squash_merge"`
	AllowMergeCommit          bool     `json:"allow_merge_commit"`
	AllowRebaseMerge          bool     `json:"allow_rebase_merge"`
	AllowAutoMerge            bool     `json:"allow_auto_merge"`
	DeleteBranchOnMerge       bool     `json:"delete_branch_on_merge"`
	AllowUpdateBranch         bool     `json:"allow_update_branch"`
	UseSquashPRTitleAsDefault bool     `json:"use_squash_pr_title_as_default"`
	SquashMergeCommitMessage  string   `json:"squash_merge_commit_message"`
	SquashMergeCommitTitle    string   `json:"squash_merge_commit_title"`
	MergeCommitMessage        string   `json:"merge_commit_message"`
	MergeCommitTitle          string   `json:"merge_commit_title"`
}

type organization struct {
	Login         string `json:"login"`
	ID            int    `json:"id"`
	Node          string `json:"node_id"`
	URL           string `json:"url"`
	Repos         string `json:"repos_url"`
	Evnts         string `json:"events_url"`
	Hooks         string `json:"hooks_url"`
	Issues        string `json:"issues_url"`
	Members       string `json:"members_url"`
	PublicMembers string `json:"public_members_url"`
	Avatar        string `json:"avatar_url"`
	Desc          string `json:"description"`
}

type enterprise struct {
	ID      int    `json:"id"`
	Slug    string `json:"slug"`
	Name    string `json:"name"`
	Node    string `json:"node_id"`
	Avatar  string `json:"avatar_url"`
	Desc    string `json:"description"`
	Website string `json:"website_url"`
	HtmlURL string `json:"html_url"`
	Created string `json:"created_at"`
	Updated string `json:"updated_at"`
}

type label struct {
	ID    int    `json:"id"`
	Node  string `json:"node_id"`
	URL   string `json:"url"`
	Name  string `json:"name"`
	Desc  string `json:"description"`
	Color string `json:"color"`
	Def   bool   `json:"default"`
}

type href struct {
	Href string `json:"href"`
}

type milestone struct {
	URL          string `json:"url"`
	HtmlURL      string `json:"html_url"`
	Labels       string `json:"labels_url"`
	ID           int    `json:"id"`
	NodeID       string `json:"node_id"`
	Number       int    `json:"number"`
	State        string `json:"state"`
	Title        string `json:"title"`
	Desc         string `json:"description"`
	Creator      user   `json:"creator"`
	OpenIssues   int    `json:"open_issues"`
	ClosedIssues int    `json:"closed_issues"`
	Created      string `json:"created_at"`
	Updated      string `json:"updated_at"`
	Closed       string `json:"closed_at"`
	Due          string `json:"due_on"`
}
type reactions struct {
	URL      string `json:"url"`
	Total    int    `json:"total_count"`
	PlusOne  int    `json:"+1"`
	MinusOne int    `json:"-1"`
	Laugh    int    `json:"laugh"`
	Hooray   int    `json:"hooray"`
	Confused int    `json:"confused"`
	Heart    int    `json:"heart"`
	Rocket   int    `json:"rocket"`
	Eyes     int    `json:"eyes"`
}

type issue struct {
	URL         string    `json:"url"`
	RepoURL     string    `json:"repository_url"`
	LabelsURL   string    `json:"labels_url"`
	CommentsURL string    `json:"comments_url"`
	EventsURL   string    `json:"events_url"`
	HtmlURL     string    `json:"html_url"`
	ID          int       `json:"id"`
	NodeID      string    `json:"node_id"`
	Number      int       `json:"number"`
	Title       string    `json:"title"`
	User        user      `json:"user"`
	Labels      []label   `json:"labels"`
	State       string    `json:"state"`
	Locked      bool      `json:"locked"`
	Assignee    user      `json:"assignee"`
	Assignees   []user    `json:"assignees"`
	Milestone   milestone `json:"milestone"`
	Comments    int       `json:"comments"`
	Created     string    `json:"created_at"`
	Updated     string    `json:"updated_at"`
	Closed      string    `json:"closed_at"`
	AuthorAssoc string    `json:"author_association"`
	ActiveLock  string    `json:"active_lock_reason"`
	Draft       bool      `json:"draft"`
	PullRequest struct {
		URL      string `json:"url"`
		HtmlURL  string `json:"html_url"`
		DiffURL  string `json:"diff_url"`
		PatchURL string `json:"patch_url"`
		MergedAt string `json:"merged_at"`
	} `json:"pull_request"`
	Body        string    `json:"body"`
	Reactions   reactions `json:"reactions"`
	TimelineURL string    `json:"timeline_url"`
	Performed   string    `json:"performed_via_github_app"`
	StateReason string    `json:"state_reason"`
}

type comment struct {
	URL         string    `json:"url"`
	HtmlURL     string    `json:"html_url"`
	IssueURL    string    `json:"issue_url"`
	ID          int       `json:"id"`
	NodeID      string    `json:"node_id"`
	User        user      `json:"user"`
	Created     string    `json:"created_at"`
	Updated     string    `json:"updated_at"`
	AuthorAssoc string    `json:"author_association"`
	Body        string    `json:"body"`
	Reactions   reactions `json:"reactions"`
	Performed   string    `json:"performed_via_github_app"`
}

type review struct {
	ID             int    `json:"id"`
	NodeID         string `json:"node_id"`
	User           user   `json:"user"`
	Body           string `json:"body"`
	Commit         string `json:"commit_id"`
	Submitted      string `json:"submitted_at"`
	State          string `json:"state"`
	HtmlURL        string `json:"html_url"`
	PullRequestURL string `json:"pull_request_url"`
	AuthorAssoc    string `json:"author_association"`
	Links          struct {
		Html        href `json:"html"`
		PullRequest href `json:"pull_request"`
	} `json:"_links"`
}

type pullRequest struct {
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
}

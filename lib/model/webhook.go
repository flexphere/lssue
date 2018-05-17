package model

type WebhookIssueParam struct {
	Issue      GithubIssue
	Repository GithubRepository
}

type GithubRepository struct {
	Name  string      `json:"name"`
	Owner GithubOwner `json:"owner"`
}

type GithubOwner struct {
	Login string `json:"login"`
}

type GithubUser struct {
	Login string `json:"login"`
}

type GithubIssue struct {
	ID        int          `json:"name"`
	Title     string       `json:"title"`
	State     string       `json:"state"`
	HTMLURL   string       `json:"html_url"`
	Assignees []GithubUser `json:"assignees"`
}

package github

// PullRequestsURL is the Github API URL
const PullRequestsURL = "https://api.github.com"

// PullRequestsResult represents a collection of PullRequests
type PullRequestsResult struct {
	PullRequests []*PullRequest
}

// PullRequest represents a single pull request
type PullRequest struct {
	Number int
	State  string
	Locked bool `json:"omitempty"`
	Title  string
	Body   string // in Markdown format
}

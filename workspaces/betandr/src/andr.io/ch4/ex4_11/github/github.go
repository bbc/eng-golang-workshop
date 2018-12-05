package github

import "time"

// PullRequestsURL is the Github API URL
const PullRequestsURL = "https://api.github.com"

// PullRequestsResult represents a collection of PullRequests
type PullRequestsResult struct {
	PullRequests []*PullRequest
}

// User represents a GitHub user
type User struct {
	Login string
}

// Head is the branch the PR is on
type Head struct {
	Ref string
}

// Base is the branch this PR's branch is from
type Base struct {
	Ref string
}

// Users is collection of `User`s
type Users []User

// PullRequest represents a single pull request
type PullRequest struct {
	Number             int
	State              string
	Locked             bool `json:"omitempty"`
	Title              string
	Body               string // in Markdown format
	User               User
	CreatedAt          time.Time `json:"created_at"`
	MergedAt           time.Time `json:"merged_at"`
	MergeCommitSha     string    `json:"merge_commit_sha"`
	Head               Head
	Base               Base
	RequestedReviewers Users `json:"requested_reviewers"`
}

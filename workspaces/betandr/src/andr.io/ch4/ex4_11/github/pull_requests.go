package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// ListPullRequestStatuses returns the statuses from a pull request
func ListPullRequestStatuses(reqURL string) ([]*Status, error) {
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	acceptHeaders := []string{"application/vnd.github.symmetra-preview+json", "application/vnd.github.sailor-v-preview+json"}
	req.Header.Set("Accept", strings.Join(acceptHeaders, ", "))
	req.Header.Set("Authorization", fmt.Sprintf("token %s", os.Getenv("OAUTH_TOKEN")))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("GET PR statuses failed: %s", res.Status)
	}

	var statuses []*Status
	if err := json.NewDecoder(res.Body).Decode(&statuses); err != nil {
		res.Body.Close()
		return nil, err
	}

	res.Body.Close()
	return statuses, nil
}

// ListPullRequestComments returns comments dfor a particular pull request
func ListPullRequestComments(repo string, number int) ([]*Comment, error) {
	reqURL := fmt.Sprintf("%s/repos/%s/issues/%d/comments", PullRequestsURL, repo, number)

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	acceptHeaders := []string{"application/vnd.github.symmetra-preview+json", "application/vnd.github.sailor-v-preview+json"}
	req.Header.Set("Accept", strings.Join(acceptHeaders, ", "))
	req.Header.Set("Authorization", fmt.Sprintf("token %s", os.Getenv("OAUTH_TOKEN")))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("GET PR %d commits from %s failed: %s", number, repo, res.Status)
	}

	var comments []*Comment
	if err := json.NewDecoder(res.Body).Decode(&comments); err != nil {
		res.Body.Close()
		return nil, err
	}

	res.Body.Close()
	return comments, nil
}

// ListPullRequestCommits returns commits for a particular pull request
func ListPullRequestCommits(repo string, number int) ([]*Commit, error) {
	reqURL := fmt.Sprintf("%s/repos/%s/pulls/%d/commits", PullRequestsURL, repo, number)

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	acceptHeaders := []string{"application/vnd.github.symmetra-preview+json", "application/vnd.github.sailor-v-preview+json"}
	req.Header.Set("Accept", strings.Join(acceptHeaders, ", "))
	req.Header.Set("Authorization", fmt.Sprintf("token %s", os.Getenv("OAUTH_TOKEN")))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("GET PR %d commits from %s failed: %s", number, repo, res.Status)
	}

	var commits []*Commit
	if err := json.NewDecoder(res.Body).Decode(&commits); err != nil {
		res.Body.Close()
		return nil, err
	}

	res.Body.Close()
	return commits, nil
}

// GetPullRequest returns a single pull request given a repo and a number
func GetPullRequest(repo string, number int) (*PullRequest, error) {
	reqURL := fmt.Sprintf("%s/repos/%s/pulls/%d", PullRequestsURL, repo, number)

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	acceptHeaders := []string{"application/vnd.github.symmetra-preview+json", "application/vnd.github.sailor-v-preview+json"}
	req.Header.Set("Accept", strings.Join(acceptHeaders, ", "))
	req.Header.Set("Authorization", fmt.Sprintf("token %s", os.Getenv("OAUTH_TOKEN")))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("GET PR %d from %s failed: %s", number, repo, res.Status)
	}

	var result *PullRequest
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		res.Body.Close()
		return nil, err
	}

	res.Body.Close()
	return result, nil
}

// ListPullRequests returns all pull requests for a repo
// https://developer.github.com/v3/pulls/#list-pull-requests
func ListPullRequests(repo string, allIssues bool) (*PullRequestsResult, error) {
	state := "state=open"
	if allIssues {
		state = "state=all"
	}
	reqURL := fmt.Sprintf("%s/repos/%s/pulls?%s", PullRequestsURL, repo, state)

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	acceptHeaders := []string{"application/vnd.github.symmetra-preview+json", "application/vnd.github.sailor-v-preview+json"}
	req.Header.Set("Accept", strings.Join(acceptHeaders, ", "))
	req.Header.Set("Authorization", fmt.Sprintf("token %s", os.Getenv("OAUTH_TOKEN")))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("list PRs from %s failed: %s", repo, res.Status)
	}

	var pulls []*PullRequest
	if err := json.NewDecoder(res.Body).Decode(&pulls); err != nil {
		res.Body.Close()
		return nil, err
	}

	result := new(PullRequestsResult)
	result.PullRequests = pulls

	res.Body.Close()
	return result, nil
}

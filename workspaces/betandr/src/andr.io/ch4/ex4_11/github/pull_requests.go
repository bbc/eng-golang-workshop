package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// ListPullRequests returns all pull requests for a repo
// https://developer.github.com/v3/pulls/#list-pull-requests
func ListPullRequests(owner, repo string, allIssues bool) (*PullRequestsResult, error) {
	state := "state=open"
	if allIssues {
		state = "state=all"
	}
	reqURL := fmt.Sprintf("%s/repos/%s/%s/pulls?%s", PullRequestsURL, owner, repo, state)

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
		return nil, fmt.Errorf("list PRs from %s/%s failed: %s", owner, repo, res.Status)
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

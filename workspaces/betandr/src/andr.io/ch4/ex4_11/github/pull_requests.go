package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ListPullRequests returns all pull requests for a repo
// https://developer.github.com/v3/pulls/#list-pull-requests
func ListPullRequests(owner, repo string) (*PullRequestsResult, error) {
	reqURL := fmt.Sprintf("%s/repos/%s/%s/pulls\n", PullRequestsURL, owner, repo)

	fmt.Println(reqURL)

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("list pull requests failed: %s", res.Status)
	}

	var result PullRequestsResult
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		res.Body.Close()
		return nil, err
	}

	res.Body.Close()
	return &result, nil
}

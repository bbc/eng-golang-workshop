// Build a tool that lets users create, read, update, and close GitHub i̶s̶s̶u̶e̶s̶
// pull requests from the command line, invoking their preferred text editor
// when substantial text input is required.
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"andr.io/ch4/ex4_11/github"
)

func locked(b bool) string {
	if b {
		return "locked"
	}

	return "unlocked"
}

func daysAgo(then time.Time) int {
	now := time.Now()
	days := now.Sub(then).Hours() / 24
	return int(days)
}

func merged(at time.Time) bool {
	if at.Year() != 1 {
		return true
	}

	return false
}

func main() {
	// This is a work in progress; it's not finished yet and pretty brittle! ;)
	if len(os.Args) <= 3 {
		fmt.Println("Usage: prt {owner} {repo} {list|get|create|delete}")
		os.Exit(0)
	}
	if os.Args[3] == "list" {
		allIssues := false
		if len(os.Args) > 4 && os.Args[4] == "--all" {
			allIssues = true
		}
		result, err := github.ListPullRequests(os.Args[1], os.Args[2], allIssues)
		if err != nil {
			log.Fatal(err)
		}

		for _, pr := range result.PullRequests {
			fmt.Printf("[\x1b[31;1m%d\x1b[0m] %s (\x1b[33;1m%s\x1b[0m)\n", pr.Number, pr.Title, pr.State)
			fmt.Printf("Head: \x1b[32;1m%s\x1b[0m Base: \x1b[32;1m%s \x1b[0m\n", pr.Head.Ref, pr.Base.Ref)
			fmt.Printf("Commit SHA: \x1b[36;1m%s\x1b[0m\n", pr.MergeCommitSha)

			if merged(pr.MergedAt) {
				fmt.Printf("Raised by %s and merged %d days ago\n", pr.User.Login, daysAgo(pr.MergedAt))
			} else {
				fmt.Printf("Raised by %s %d days ago\n", pr.User.Login, daysAgo(pr.CreatedAt))
			}

			if len(pr.RequestedReviewers) > 0 {
				rs := []string{}
				for _, r := range pr.RequestedReviewers {
					rs = append(rs, r.Login)
				}
				fmt.Printf("Requested reviewers: %s\n", strings.Join(rs, ", "))
			}

			fmt.Printf("%s\n\n", pr.Body)
		}
	}
}

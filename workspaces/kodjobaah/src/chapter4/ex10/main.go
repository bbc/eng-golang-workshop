// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"chapter4/github"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	categories := make(map[string][]*github.Issue) // counts of Unicode characters
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	var monthBefore = time.Now().AddDate(0, -1, 0)
	var yearBefore = time.Now().AddDate(-1, 0, 0)
	for _, item := range result.Items {

		if item.CreatedAt.After(monthBefore) {
			categories["less-than-month"] = append(categories["less-than-month"], item)
		}

		if monthBefore.After(item.CreatedAt) && item.CreatedAt.After(yearBefore) {
			categories["less-than-year"] = append(categories["less-than-year"], item)
		}

		if monthBefore.After(item.CreatedAt) && yearBefore.After(item.CreatedAt) {
			categories["more-than-year"] = append(categories["more-than-year"], item)
		}

	}

	for section, items := range categories {

		fmt.Println(section)
		for _, item := range items {
			fmt.Printf("#%-5d %9.9s %.55s  %10s\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt)
		}

	}

}

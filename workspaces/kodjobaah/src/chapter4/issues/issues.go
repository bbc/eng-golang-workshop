// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"chapter4/github"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("here")
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s  %10s\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt)
	}

}

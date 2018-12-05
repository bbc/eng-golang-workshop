// Build a tool that lets users create, read, update, and close GitHub issues
// from the command line, invoking their preferred text editor when substantial
// text input is required.
package main

import (
	"fmt"
	"log"
	"os"

	"andr.io/ch4/ex4_11/github"
)

// prt betandr sample list

// !!WORK IN PROGRESS!!

func main() {
	//
	if len(os.Args) <= 3 {
		fmt.Println("Usage: prt {owner} {repo} list")
		os.Exit(0)
	}
	if os.Args[3] == "list" {
		result, err := github.ListPullRequests(os.Args[1], os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(result)
		// fmt.Printf("Returned %d pull requests", len(result.PullRequests))
	}

}

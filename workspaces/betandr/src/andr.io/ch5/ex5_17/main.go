// Write a variadic function `ElementsByTagName` that, given an HTML node tree
// and zero or more names, returns all the elements that match one of those
// names. Here are two example calls:
//     `func ElementsByTagName(doc *html.Node, name ...string) []*html.Node`
//     `images := ElementsByTagName(doc, "img")``
//     `headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")``
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// Run with either a list of URLs, such as `element title https://foo.com https://bar.com`
// or with standard input such as `cat index.html | element title`
func main() {
	if len(os.Args) > 1 {
		for _, url := range os.Args[2:] {
			resp, err := http.Get(url)
			if err != nil {
				log.Fatalf("get %s failed: %s", url, err)
			}
			defer resp.Body.Close()

			doc, err := html.Parse(resp.Body)
			if err != nil {
				log.Fatalf("parse failed: %s", err)
			}

			images := ElementsByTagName(doc, "img")
			for image := range images {
				fmt.Printf("found: %v\n", image)
			}

			headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
			for heading := range headings {
				fmt.Printf("found: %v\n", heading)
			}
		}
	} else {
		doc, err := html.Parse(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
			os.Exit(1)
		}
		images := ElementsByTagName(doc, "img")
		for image := range images {
			fmt.Printf("found: %v\n", image)
		}

		headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
		for heading := range headings {
			fmt.Printf("found: %v\n", heading)
		}
	}
}

// ElementsByTagName returns all the elements of an an HTML node tree that match
// one of zero or more names given.
func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	return visit([]*html.Node{}, doc, name...)
}

// TODO:
func visit(nodes []*html.Node, n *html.Node, id ...string) []*html.Node {
	if n.Type == html.ElementNode {
		if n.Data == "img" {

		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = visit(nodes, c, id...)
	}

	return nodes
}

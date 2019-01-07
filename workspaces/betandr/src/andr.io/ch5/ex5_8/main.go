// Modify `forEachNode` so that the `pre` and `post` functions return a boolean
// result indicating whether to continue the traversal. Use it to write a
// function `ElementByID` with the following signature that finds the first
// HTML element with the specific id attribute. The function should stop the
// traversal as soon as a match is found.
//      `func ElementByID(doc *html.Node, id string) *html.Node`
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[2:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not get %s", url)
		}
		defer resp.Body.Close()

		doc, err := html.Parse(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not parse body: %v", err)
		}

		element := ElementByID(doc, os.Args[1])
		fmt.Printf("%s\n", element.Data)
	}
}

// ElementByID finds an element `id` in an `html.Node` document.
func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, id, startElement, endElement)
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) *html.Node {
	if pre != nil {
		found := pre(n, id)
		if found {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, id, pre, post)
	}

	if post != nil {
		found := post(n, id)
		if found {
			return n
		}
	}

	return nil
}

var depth int

func startElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		if n.Data == id {
			return true
		}
	}
	return false
}

func endElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		if n.Data == id {
			return true
		}
	}
	return false
}

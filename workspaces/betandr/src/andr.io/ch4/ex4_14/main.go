// Create a web server that queries GitHub once and then allows navigation of
// the list of bug reports, milestones, and users.
package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"gopl.io/ch4/github"
)

func main() {
	issues := func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		terms := r.Form["terms"]
		issues(terms, w)
	}

	http.HandleFunc("/issues/", issues)
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
	return
}

func issues(terms []string, out io.Writer) {
	templ := `
    <h1>{{.TotalCount}} issues</h1>
    <table>
    <tr style='text-align: left'>
      <th>#</th>
      <th>State</th>
      <th>User</th>
      <th>Title</th>
    </tr>
    {{range .Items}}
    <tr>
      <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
      <td>{{.State}}</td>
      <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
      <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
    </tr>
    {{end}}
    </table>
    `

	var issueList = template.Must(template.New("issuelist").Parse(templ))

	result, err := github.SearchIssues(terms)
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(out, result); err != nil {
		log.Fatal(err)
	}
}

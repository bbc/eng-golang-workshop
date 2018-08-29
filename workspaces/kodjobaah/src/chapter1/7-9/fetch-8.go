// Fetch prints the content found at each specified url
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"os"
)

func main(){

	for _, url := range os.Args[1:] {
		newUrl := url
		if !(strings.HasPrefix(url, "http://")) {
			newUrl = "http://" + url
		}

		resp, err := http.Get(newUrl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:  %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)

	}
}

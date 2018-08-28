# BBC Gophers — Go Language Workshop Repository

![Gopher](resources/gopher.png)

```
func main() {
	fmt.Println("Hello, BBC Gophers!")
}
```

We are on Slack as #eng-golang

We use this book for study:
[The Go Programming Language by Alan Donovan and Brian W. Kernighan](https://www.gopl.io/)

Each week we will complete one chapter. To take part, create your own directory
in the [workspaces](workspaces) directory with whatever name you chose. Change
directory to your workspace Then un-tar the examples there.

e.g.

```
cd workspace
mkdir bobbyfoo
cd bobbyfoo
tar xzvf ../../exercises.tar.gz
```

Add any source code and stuff in the exercise directory. Commit directly to
master. Just don't break things...well, fix things if they do break! :)

## Installing Go

###  Home/Linuxbrew
It may be convenient to install the latest version of Go through the
[Homebrew](https://brew.sh/) and [Linuxbrew](http://linuxbrew.sh/) package
managers.

```
brew install go
```


### Install with Binary Distributions
The [https://golang.org/dl/](https://golang.org/dl/) page contains distros for
Windows, MacOS, Linux, and source. The
[installation instructions](https://golang.org/doc/install) explains how to
install them.

### Go Modules

_Go modules_ are an experimental opt-in feature in Go 1.11. Follow this
documentation to [install and activate modules](https://github.com/golang/go/wiki/Modules#installing-and-activating-module-support).

## Links

[A Tour of Go](https://tour.golang.org/welcome/1)

[How to Write Go Code](https://golang.org/doc/code.html)

[Effective Go](https://golang.org/doc/effective_go.html)

[Source code: The Go Programming Language](https://github.com/adonovan/gopl.io)

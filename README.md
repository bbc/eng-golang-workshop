# BBC Gophers — Go Language Workshop

<p align="center">
  <img width="225" height="300" src="https://github.com/bbc/eng-golang-workshop/raw/master/resources/gopher.gif">
</p>

```
func main() {
	fmt.Println("Hello, BBC Gophers!")
}
```

## The Workshop
We are on Slack as `#eng-golang`

Our workshop text is:
[The Go Programming Language by Alan Donovan and Brian W. Kernighan](https://www.gopl.io/)

Each week we will complete one chapter then (_optionally_) meet to chat about it.
To join in, create your own directory in the [workspaces](workspaces) directory
with whatever name you chose. Change directory to your workspace and then you
can un-tar the examples there if you wish (or refer to your book).

e.g.
```
cd workspace
mkdir bobbyfoo
cd bobbyfoo
tar xzvf ../../exercises.tar.gz
```

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
documentation to
[install and activate modules](https://github.com/golang/go/wiki/Modules#installing-and-activating-module-support).

## Running Go

To run your exercises without using _Go Modules_ you should set your `$GOPATH`
to your current directory, such as:
```
export GOPATH=/home/bobbyf/eng-golang-workshop/workspaces/bobbyfoo
```

Then when runnning something like `go get gopl.io/ch1/helloworld` the source will
be placed in your `$GOPATH/src` directory and
`go build src/gopl.io/ch1/helloworld` will be placed in your `$GOPATH/bin`
directory etc.

A project structure could then be:

```
workspaces
    bobbyfoo\
        exercises\
            ...
        bin\
            helloworld
            ...
        src\
            bobbyfoo.com\
                ch1\
                    ex1_1\
                        main.go
                ...
```

Commit directly to master and fix stuff if it breaks. :)

## Links

[A Tour of Go](https://tour.golang.org/welcome/1)

[How to Write Go Code](https://golang.org/doc/code.html)

[Effective Go](https://golang.org/doc/effective_go.html)

[Source code: The Go Programming Language](https://github.com/adonovan/gopl.io)

[YouTube: Concurrency is not Parallelism by Rob Pike](https://www.youtube.com/watch?v=oV9rvDllKEg)

[YouTube: Go Proverbs](https://www.youtube.com/watch?v=PAAkCSZUG1c)

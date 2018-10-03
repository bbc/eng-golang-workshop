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

Work through the course text at a speed that suits you, although one
chapter per week is a good target to aim for (if you have the time). We also have
meetups booked to meet to chat about Go and the book etc. So just join in any
time! :)

## Get Coding

To begin, create your own directory in the [workspaces](workspaces) directory,
work through the book exercises, and add your code there.
[Project Structure](#project-structure) contains some ideas to structure your code.

If you need the exercises from the book, they're available in
[exercises.tar.gz](exercises.tar.gz) zipped archive but these are already in the
book so you might not need them.

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

## Getting the Workshop Code

You can grab the code samples used in the book (which you update for a bunch of
the exercises) by running such as:

```
go get gopl.io/ch1/helloworld
```

(This will get the `helloworld` code, plus the other examples).

## Running Go

To run your exercises without using _Go Modules_ you should set your `$GOPATH`
to your current directory, such as:
```
export GOPATH=/home/gopherg/eng-golang-workshop/workspaces/gogopher
```

Then when runnning something like `go get gopl.io/ch1/helloworld` the source will
be placed in your `$GOPATH/src` directory and
`go build src/gopl.io/ch1/helloworld` will be placed in your `$GOPATH/bin`
directory etc.

## Go Modules

_Go modules_ are an experimental opt-in feature in Go 1.11. Follow this
documentation to
[install and activate modules](https://github.com/golang/go/wiki/Modules#installing-and-activating-module-support).

## Project Structure

A project structure could be:

```
workspaces
    gogopher\
        bin\
            helloworld
            ...
        src\
            gogopher.io\
                ch1\
                    ex1_1\
                        main.go
                ...
```
With this directory structure, and your $GOPATH set as in the
[Running Go](#running-go) section you can run the `main.go` file with:

`go run src/gogopher.com/ch1/ex1_1/main.go`

...or use `go build` to build the binary.

## Development Environments

### Delve (Debugger)

[Delve](https://github.com/derekparker/delve) is is a debugger for Go. To install run:
```
go get -u github.com/derekparker/delve/cmd/dlv
```
To see the available commands, run `dlv` then `help` at the `(dlv)` prompt.

### GoLand

_todo_

### Emacs

If you follow similar instructions to get go support for emacs (OS X) as below
http://tleyden.github.io/blog/2014/05/22/configure-emacs-as-a-go-editor-from-scratch/

And you run into the following error when trying to get auto-complete to work.

```
Error running timer ‘ac-update-greedy’: (file-missing "Searching for program" "No such file or directory" "gocode")
Error running timer ‘ac-show-menu’: (file-missing "Searching for program" "No such file or directory" "gocode")
Error running timer ‘ac-update-greedy’: (file-missing "Searching for program" "No such file or directory" "gocode"
```

Then the problem is probably down to `gocode` not being available in your path:
https://emacs.stackexchange.com/questions/10722/emacs-and-command-line-path-disagreements-on-osx

So if you edit `/etc/paths.d/go` and add the path to the bin directory of your project it should fix problem.

### Atom

Atom supports Go development with the
[go-plus](https://atom.io/packages/go-plus) package.

To use Delve inside Atom, install the
[go-debug](https://atom.io/packages/go-debug) package.

To run your Go code in Atom, install the
[atom-runner](https://atom.io/packages/atom-runner) package.

## Links

[A Tour of Go](https://tour.golang.org/welcome/1)

[How to Write Go Code](https://golang.org/doc/code.html)

[Effective Go](https://golang.org/doc/effective_go.html)

[Source code: The Go Programming Language](https://github.com/adonovan/gopl.io)

[YouTube: Concurrency is not Parallelism by Rob Pike](https://www.youtube.com/watch?v=oV9rvDllKEg)

[YouTube: Go Proverbs](https://www.youtube.com/watch?v=PAAkCSZUG1c)

## Rights

All exercises from The Go Programming Language are copyright 2016 Alan A. A. Donovan & Brian W. Kernighan and included with permission from the authors.

All submitted code is covered under [Apache License 2.0](LICENSE).

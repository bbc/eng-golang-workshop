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

Feel free to work through the text at a speed that suits you, although one
chapter per week is a good target to aim for (if you have the time). We also have
sessions booked to meet to chat about Go and the book etc. So just join in any
time! :)

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

## Development Environments

### GoLand
_todo_

### Emacs
_todo_

### Atom
Atom supports Go development with the
[go-plus](https://atom.io/packages/go-plus) package amongst other tools.

## Editors

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


## Links

[A Tour of Go](https://tour.golang.org/welcome/1)

[How to Write Go Code](https://golang.org/doc/code.html)

[Effective Go](https://golang.org/doc/effective_go.html)

[Source code: The Go Programming Language](https://github.com/adonovan/gopl.io)

[YouTube: Concurrency is not Parallelism by Rob Pike](https://www.youtube.com/watch?v=oV9rvDllKEg)

[YouTube: Go Proverbs](https://www.youtube.com/watch?v=PAAkCSZUG1c)

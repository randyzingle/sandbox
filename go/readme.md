# Golang

## Create a new project

Initialize the module-based project.

```sh
$ mkdir hello
$ cd hello
$ go mod init hello
$ go mod init hello
go: creating new go.mod: module hello
$ ls
go.mod
$ cat go.mod
module hello

go 1.15
```

Create a file that pulls in an external package:

```sh
$ cat hello.go
package main

import "fmt"
import "rsc.io/quote"

func main() {
  fmt.Println(quote.Go())
}
```

Fire it up and it will pull down the library and add it to your module dependencies:
```sh
razing hello $ go run hello.go
go: finding module for package rsc.io/quote
go: found rsc.io/quote in rsc.io/quote v1.5.2
Dont communicate by sharing memory, share memory by communicating.
razing hello $ cat go.mod
module hello

go 1.15

require rsc.io/quote v1.5.2 // indirect
```

I'll be pushing my code to: https://github.com/randyzingle/sandbox/go

## Multi-module Build

Greetings Module:

```sh
razing go $ mkdir greetings
razing go $ cd greetings
razing greetings $ go mod init github.com/randyzingle/sandbox/go/greetings
go: creating new go.mod: module github.com/randyzingle/sandbox/go/greetings
```

Hello Module Update:

```sh
razing hello $ cat hello.go
package main

import "fmt"
import "github.com/randyzingle/sandbox/go/greetings"

func main() {
  message := greetings.Hello("Mymir")
  fmt.Println(message)
}
$ go build hello.go
go: finding module for package github.com/randyzingle/sandbox/go/greetings
go: downloading github.com/randyzingle/sandbox v0.0.0-20200908201250-78df0726affc
hello.go:4:8: module github.com/randyzingle/sandbox@latest found (v0.0.0-20200908201250-78df0726affc), but does not contain package github.com/randyzingle/sandbox/go/greetings
```

We haven't pushed our greetings module to github yet. We could either reference it locally for now  or push our code to github and run as is.

Use local version:
```sh
razing hello $ cat go.mod
module hello

go 1.15

replace github.com/randyzingle/sandbox/go/greetings => ../greetings
```

Run it:
```
razing hello $ go run hello.go
Hi there Mymir, welcome to Baldursoft inc.
razing hello $ cat go.mod
module hello

go 1.15

replace github.com/randyzingle/sandbox/go/greetings => ../greetings

require github.com/randyzingle/sandbox/go/greetings v0.0.0-00010101000000-000000000000 // indirect
```

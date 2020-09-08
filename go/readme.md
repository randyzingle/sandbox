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

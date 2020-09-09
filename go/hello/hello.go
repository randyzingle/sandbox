package main

import "fmt"
import "github.com/randyzingle/sandbox/go/greetings"

func main() {
  message := greetings.Hello("Mymir")
  fmt.Println(message)
}

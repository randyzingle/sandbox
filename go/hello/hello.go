package main

import "fmt"
import "github.com/randyzingle/sandbox/tree/latest/go/greetings"

func main() {
  message := greetings.Hello("Mymir")
  fmt.Println(message)
}

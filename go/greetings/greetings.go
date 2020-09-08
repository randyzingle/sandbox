package greetings

import "fmt"

// Hello returns a greeting for the named person
func Hello(name string) string {
  message := fmt.Sprintf("Hi there %, welcome to Baldursoft inc.", name)
  return message
}

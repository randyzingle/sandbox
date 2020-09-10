package greetings

import "fmt"

// Hello returns a greeting for the named person
func Hello(name string) string {
  message := fmt.Sprintf("Baldurdash! Hi there %v, welcome to Baldursoft inc.", name)
  return message
}

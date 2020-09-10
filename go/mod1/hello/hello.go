package hello

import "rsc.io/quote/v3"
import "fmt"

func Hello() string {
  fmt.Println(quote.HelloV3())
  return "Hello there!"
}

func Proverb() string {
  fmt.Println(quote.Concurrency())
  return quote.Concurrency()
}

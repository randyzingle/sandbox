package hello

import "testing"

func TestHello(t *testing.T) {
  want := "Hello there!"
  if got := Hello(); got != want {
    t.Errorf("Hello() = %v, want %v", got, want)
  }
}

func TestProverb(t *testing.T) {
  want := "Concurrency is not parallelism."
  if got := Proverb(); got != want {
      t.Errorf("Proverb() = %q, want %q", got, want)
  }
}

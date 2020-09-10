package dogs

import "testing"

func TestBark(t *testing.T) {
  want := "bark bark"
  if got := Bark(); want != got {
    t.Errorf("Bark() = %v, want %v", got, want)
  }
}

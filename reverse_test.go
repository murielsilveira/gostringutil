package gostringutil

import "testing"

func TestReverse(t *testing.T) {
  input := "!oG olleG"
  expected := "Hello Go!"

  got := Reverse(input)

  if got != expected {
    t.Errorf("Reverse(%q) == %q, expected %q", input, got, expected)
  }
}

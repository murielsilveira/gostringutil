package gostringutil

import "testing"

func TestReverse(t *testing.T) {
  cases := []struct { input, expected string } {
    {"Hello Go!", "!oG olleH"},
    {"Hello 世界", "界世 olleH"},
    {"å", "å"},
    {"", ""},
  }

  for _, c := range cases {
    got := Reverse(c.input)
    if got != c.expected {
      t.Errorf("Reverse(%q) == %q, expected %q", c.input, got, c.expected)
    }
  }
}

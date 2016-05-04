package gostringutil

import "testing"

func TestSimpleReverse(t *testing.T) {
  input := "Hello Go!"
  expected := "!oG olleH"

  got := Reverse(input)

  if got != expected {
    t.Errorf("Reverse(%q) == %q, expected %q", input, got, expected)
  }
}

func TestWithChineseCharacters(t *testing.T) {
  input := "Hello 世界"
  expected := "界世 olleH"

  got := Reverse(input)

  if got != expected {
    t.Errorf("Reverse(%q) == %q, expected %q", input, got, expected)
  }
}

func TestWithASingleCharacter(t *testing.T) {
  input := "å"
  expected := "å"

  got := Reverse(input)

  if got != expected {
    t.Errorf("Reverse(%q) == %q, expected %q", input, got, expected)
  }
}

func TestWithEmptyString(t *testing.T) {
  input := ""

  got := Reverse(input)

  if got != input {
    t.Errorf("Reverse(%q) == %q, expected %q", input, got, input)
  }
}

func TestReverseWithAllCases(t *testing.T) {
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

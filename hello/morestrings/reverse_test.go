// go test github.com/denglitong/golang/hello/morestrings
//
// Go has a lightweight test framework composed of the `go test` command
// and the `testing` standard package.
package morestrings

import "testing"

// You write a test by creating a file with a name ending in `_test.go`
// that contains functions named `TestXxx` with signature `func(t *testing.T)`,
// the test framework runs each function.
func TestReverseRunes(t *testing.T) {
	cases := []struct {
		input, expected string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		actual := ReverseRunes(c.input)
		if actual != c.expected {
			// Ff the function calls a failure function
			// such as t.Error(), t.Errorf() or t.Fail(), the test is considered to have failed.
			t.Errorf("ReverseRunes(%q) == %q, expected %q", c.input, actual, c.expected)
		}
	}
}

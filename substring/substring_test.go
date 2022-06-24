package substring

import (
	"fmt"
	"testing"
)

type testCase struct {
	input       string // string to be redacted
	substring   string // target substring
	replacement string // replacement string
	expected    string // expected output
}

func TestRedact(t *testing.T) {
	cases := []testCase{
		{"this is a test.", "test", "[redacted]", "this is a [redacted]."},
		{"abcdefg deabcfg defgabc", "abc", "XXX", "XXXdefg deXXXfg defgXXX"},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("input=%q;substring=%q;replacement=%q;expected=%q; ",
			tc.input, tc.substring, tc.replacement, tc.expected), func(t *testing.T) {
			r := New(tc.substring, tc.replacement)
			got := r.Redact(tc.input)
			if tc.expected != got {
				t.Errorf("Expected '%s', but got '%s'", tc.expected, got)
			}
		})
	}
}

func TestString(t *testing.T) {
	redactor := New("foo", "bar")
	stringer := redactor.(fmt.Stringer)

	expected := `{substring="foo"; replacement="bar"}`
	got := stringer.String()

	if expected != got {
		t.Errorf("Expected '%s', but got '%s'", expected, got)
	}
}

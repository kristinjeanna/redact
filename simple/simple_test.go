package simple

import (
	"fmt"
	"testing"
)

type testCase struct {
	input       string // string to be redacted
	replacement string // replacement string
	expected    string // expected output
}

func TestRedact(t *testing.T) {
	cases := []testCase{
		{"this is a test.", "[redacted]", "[redacted]"},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("input=%q;replacement=%q;expected=%q; ", tc.input, tc.replacement, tc.expected), func(t *testing.T) {
			r := New(tc.replacement)
			got := r.Redact(tc.input)
			if tc.expected != got {
				t.Errorf("Expected '%s', but got '%s'", tc.expected, got)
			}
		})
	}
}

package blackout

import (
	"fmt"
	"testing"
)

type testCase struct {
	input       string // string to be redacted
	replacement string // replacement char(s)
	expected    string // expected output
}

func TestRedact(t *testing.T) {
	cases := []testCase{
		{"this is a test.", "x", "xxxx xx x xxxxx"},
		{"password", "██", "████████████████"},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("input=%q;replacement=%q;expected=%q; ", tc.input, tc.replacement, tc.expected), func(t *testing.T) {
			r := New(tc.replacement)
			got, err := r.Redact(tc.input)
			if err != nil {
				t.Error(err)
			}
			if tc.expected != got {
				t.Errorf("Expected '%s', but got '%s'", tc.expected, got)
			}
		})
	}
}

func TestString(t *testing.T) {
	redactor := New("x")
	stringer := redactor.(fmt.Stringer)

	expected := `{replacement="x"}`
	got := stringer.String()

	if expected != got {
		t.Errorf("Expected '%s', but got '%s'", expected, got)
	}
}

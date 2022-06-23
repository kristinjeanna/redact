package regex

import (
	"fmt"
	"testing"
)

var (
	pair1, pair2, pair3, pair4, pair5 *Pair
)

func setupPairs() {
	pair1, _ = NewPair("[redacted]", "test")
	pair2, _ = NewPair("X", "[is]")
	pair3, _ = NewPair("XXXX", ".*")
	pair4, _ = NewPair("${1}XXXX${3}", `(.*)(b[aA][rR])(.*)`)
	pair5, _ = NewPair("${1}zed${3}", `^(.*\s)(.*)$`)
}

type testCase struct {
	input    string // string to be redacted
	pairs    []Pair //  the regex/replacement pairs
	expected string // expected output
}

func TestRedact(t *testing.T) {
	setupPairs()
	cases := []testCase{
		{"this is a test.", []Pair{*pair1}, "this is a [redacted]."},
		{"this string contains sensitive information", []Pair{*pair2},
			"thXX XtrXng contaXnX XenXXtXve XnformatXon"},
		{"foo is bar", []Pair{*pair3}, "XXXX"},
		{"foo is bar zaz", []Pair{*pair4}, "foo is XXXX zaz"},
		{"foo is bar zaz", []Pair{*pair4, *pair5}, "foo is XXXX zed"},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("input=%q;pairs=%q;expected=%q; ", tc.input, tc.pairs, tc.expected), func(t *testing.T) {
			r, err := New(tc.pairs)
			if err != nil { // err not expected
				t.Error(err)
			}
			got := r.Redact(tc.input)
			if tc.expected != got {
				t.Errorf("Expected '%s', but got '%s'", tc.expected, got)
			}
		})
	}
}

type testCaseErr struct {
	replacement string // replacement string
	pairs       []Pair // the regex/replacement pairs
	expected    error  // the expected error
}

func TestRedact_err(t *testing.T) {
	cases := []testCaseErr{
		{"[redacted]", nil, errRePairsSliceNil},
		{"[redacted]", []Pair{}, errRePairsSliceEmpty},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("pairs=%q;err=%q", tc.pairs, tc.expected), func(t *testing.T) {
			_, err := New(tc.pairs)
			if err == nil {
				t.Errorf("expected error")
			}
			if err != tc.expected {
				t.Errorf("Expected '%v', but got '%v'", tc.expected, err)
			}
		})
	}
}

func TestString(t *testing.T) {
	pair, _ := NewPair("[redacted]", "test")
	redactor, _ := New([]Pair{*pair})
	reRedactor := redactor.(fmt.Stringer)

	expected := `{pairs=["{regex=\"test\"; replacement=\"[redacted]\"}"]}`
	got := reRedactor.String()

	if expected != got {
		t.Errorf("Expected '%s', but got '%s'", expected, got)
	}
}

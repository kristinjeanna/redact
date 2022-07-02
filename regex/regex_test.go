package regex

import (
	"fmt"
	"testing"

	"github.com/kristinjeanna/redact"
	"github.com/kristinjeanna/redact/blackout"
	"github.com/kristinjeanna/redact/middle"
	"github.com/kristinjeanna/redact/url"
)

var (
	pair1, pair2, pair3, pair4, pair5, pair6, pair7 *Pair
)

func setupPairs() {
	pair1, _ = NewPairUsingSimple("[redacted]", "test")
	pair2, _ = NewPairUsingSimple("X", "[is]")
	pair3, _ = NewPairUsingSimple("XXXX", ".*")
	pair4, _ = NewPairUsingSimple("${1}XXXX${3}", `(.*)(b[aA][rR])(.*)`)
	pair5, _ = NewPairUsingSimple("${1}zed${3}", `^(.*\s)(.*)$`)

	pair6, _ = NewPair(middle.New(), `b[aA][rRzZ]`)
	pair7, _ = NewPair(blackout.New("X"), SSNRegex)
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
		{"foo is bar baz", []Pair{*pair6}, "foo is [redacted] [redacted]"},
		{"the SSN is 123-45-6789", []Pair{*pair7}, "the SSN is XXXXXXXXXXX"},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("input=%q;pairs=%q;expected=%q; ", tc.input, tc.pairs, tc.expected), func(t *testing.T) {
			r, err := New(tc.pairs)
			if err != nil { // err not expected
				t.Error(err)
			}
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

type testCaseErr struct {
	replacement string // replacement string
	pairs       []Pair // the regex/replacement pairs
	expected    error  // the expected error
}

func TestRedact_errorsSpecific(t *testing.T) {
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

func TestRedact_errorsInfLoop(t *testing.T) {
	s := "this string contains redacted"
	m, err := middle.NewFromOptions(middle.WithReplacementText("[redacted]"))
	if err != nil {
		t.Error(err)
	}

	pair, err := NewPair(m, "redacted")
	if err != nil {
		t.Error(err)
	}

	var r redact.Redactor
	r, err = New([]Pair{*pair})
	if err != nil {
		t.Error(err)
	}

	_, err = r.Redact(s)
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}

func TestRedact_errorsRedactFail(t *testing.T) {
	s := "this string contains redacted"
	u := url.New("foo", nil)

	pair, err := NewPair(u, "[is]")
	if err != nil {
		t.Error(err)
	}

	var r redact.Redactor
	r, err = New([]Pair{*pair})
	if err != nil {
		t.Error(err)
	}

	_, err = r.Redact(s)
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}

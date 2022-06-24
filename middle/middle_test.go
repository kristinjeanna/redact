package middle

import (
	"fmt"
	"testing"

	"github.com/kristinjeanna/redact"
)

type testCase struct {
	input    string // string to be redacted
	expected string // expected output
}

func TestRedact(t *testing.T) {
	cases := []testCase{
		{"", ""},
		{"a", "[redacted]"},
		{"ab", "[redacted]"},
		{"abc", "[redacted]"},
		{"abcd", "[redacted]"},
		{"abcde", "[redacted]"},
		{"abcdef", "[redacted]"},
		{"abcdefg", "[redacted]"},
		{"abcdefgh", "[redacted]"},
		{"abcdefghi", "[redacted]"},
		{"abcdefghij", "[redacted]"},
		{"abcdefghijk", "[redacted]"},
		{"abcdefghijkl", "[redacted]"},

		{"abcdefghijklm", "abc[redacted]"},
		{"abcdefghijklmn", "abc[redacted]"},
		{"abcdefghijklmno", "abc[redacted]"},

		{"abcdefghijklmnop", "abc[redacted]nop"},
		{"abcdefghijklmnopq", "abc[redacted]opq"},
		{"abcdefghijklmnopqrs", "abc[redacted]qrs"},
		{"abcdefghijklmnopqrst", "abc[redacted]rst"},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("input=%q;expected=%q; ", tc.input, tc.expected), func(t *testing.T) {
			r := New()
			got := r.Redact(tc.input)
			if tc.expected != got {
				t.Errorf("Expected '%s', but got '%s'", tc.expected, got)
			}
		})
	}
}

func TestPrefixOnlyMode(t *testing.T) {
	cases := []testCase{
		{"", ""},
		{"a", "[redacted]"},
		{"ab", "[redacted]"},
		{"abc", "[redacted]"},
		{"abcd", "[redacted]"},
		{"abcde", "[redacted]"},
		{"abcdef", "[redacted]"},
		{"abcdefg", "[redacted]"},
		{"abcdefgh", "[redacted]"},
		{"abcdefghi", "[redacted]"},
		{"abcdefghij", "[redacted]"},
		{"abcdefghijk", "[redacted]"},
		{"abcdefghijkl", "[redacted]"},

		{"abcdefghijklm", "abc[redacted]"},
		{"abcdefghijklmn", "abc[redacted]"},
		{"abcdefghijklmno", "abc[redacted]"},
		{"abcdefghijklmnop", "abc[redacted]"},
		{"abcdefghijklmnopq", "abc[redacted]"},
		{"abcdefghijklmnopqr", "abc[redacted]"},
		{"abcdefghijklmnopqrs", "abc[redacted]"},
		{"abcdefghijklmnopqrst", "abc[redacted]"},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("input=%q;expected=%q; ", tc.input, tc.expected), func(t *testing.T) {
			r, err := NewFromOptions(WithMode(PrefixOnlyMode))
			if err != nil {
				t.Error(err)
			}

			got := r.Redact(tc.input)
			if tc.expected != got {
				t.Errorf("Expected '%s', but got '%s'", tc.expected, got)
			}
		})
	}
}

func TestSuffixOnlyMode(t *testing.T) {
	cases := []testCase{
		{"", ""},
		{"a", "[redacted]"},
		{"ab", "[redacted]"},
		{"abc", "[redacted]"},
		{"abcd", "[redacted]"},
		{"abcde", "[redacted]"},
		{"abcdef", "[redacted]"},
		{"abcdefg", "[redacted]"},
		{"abcdefgh", "[redacted]"},
		{"abcdefghi", "[redacted]"},
		{"abcdefghij", "[redacted]"},
		{"abcdefghijk", "[redacted]"},
		{"abcdefghijkl", "[redacted]"},

		{"abcdefghijklm", "[redacted]klm"},
		{"abcdefghijklmn", "[redacted]lmn"},
		{"abcdefghijklmno", "[redacted]mno"},

		{"abcdefghijklmnop", "[redacted]nop"},
		{"abcdefghijklmnopq", "[redacted]opq"},
		{"abcdefghijklmnopqr", "[redacted]pqr"},
		{"abcdefghijklmnopqrs", "[redacted]qrs"},
		{"abcdefghijklmnopqrst", "[redacted]rst"},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("input=%q;expected=%q; ", tc.input, tc.expected), func(t *testing.T) {
			r, err := NewFromOptions(WithMode(SuffixOnlyMode))
			if err != nil {
				t.Error(err)
			}

			got := r.Redact(tc.input)
			if tc.expected != got {
				t.Errorf("Expected '%s', but got '%s'", tc.expected, got)
			}
		})
	}
}

func TestWithPrefixLength(t *testing.T) {
	cases := []testCase{
		{"", ""},
		{"a", "[redacted]"},
		{"ab", "[redacted]"},
		{"abc", "[redacted]"},
		{"abcd", "[redacted]"},
		{"abcde", "[redacted]"},
		{"abcdef", "[redacted]"},
		{"abcdefg", "[redacted]"},
		{"abcdefgh", "[redacted]"},
		{"abcdefghi", "[redacted]"},
		{"abcdefghij", "[redacted]"},
		{"abcdefghijk", "[redacted]"},
		{"abcdefghijkl", "[redacted]"},
		{"abcdefghijklm", "[redacted]"},
		{"abcdefghijklmn", "[redacted]"},

		{"abcdefghijklmno", "abcde[redacted]"}, // 15
		{"abcdefghijklmnop", "abcde[redacted]"},
		{"abcdefghijklmnopq", "abcde[redacted]"},

		{"abcdefghijklmnopqr", "abcde[redacted]pqr"}, // 18
		{"abcdefghijklmnopqrs", "abcde[redacted]qrs"},
		{"abcdefghijklmnopqrst", "abcde[redacted]rst"},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("input=%q;expected=%q; ", tc.input, tc.expected), func(t *testing.T) {
			r, err := NewFromOptions(WithPrefixLength(5))
			if err != nil {
				t.Error(err)
			}

			got := r.Redact(tc.input)
			if tc.expected != got {
				t.Errorf("Expected '%s', but got '%s'", tc.expected, got)
			}
		})
	}
}

func TestWithSuffixLength(t *testing.T) {
	cases := []testCase{
		{"", ""},
		{"a", "[redacted]"},
		{"ab", "[redacted]"},
		{"abc", "[redacted]"},
		{"abcd", "[redacted]"},
		{"abcde", "[redacted]"},
		{"abcdef", "[redacted]"},
		{"abcdefg", "[redacted]"},
		{"abcdefgh", "[redacted]"},
		{"abcdefghi", "[redacted]"},
		{"abcdefghij", "[redacted]"},
		{"abcdefghijk", "[redacted]"},
		{"abcdefghijkl", "[redacted]"},

		{"abcdefghijklm", "abc[redacted]"}, // 13
		{"abcdefghijklmn", "abc[redacted]"},

		{"abcdefghijklmno", "abc[redacted]"}, // 15
		{"abcdefghijklmnop", "abc[redacted]"},
		{"abcdefghijklmnopq", "abc[redacted]"},

		{"abcdefghijklmnopqr", "abc[redacted]nopqr"}, // 18
		{"abcdefghijklmnopqrs", "abc[redacted]opqrs"},
		{"abcdefghijklmnopqrst", "abc[redacted]pqrst"},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("input=%q;expected=%q; ", tc.input, tc.expected), func(t *testing.T) {
			r, err := NewFromOptions(WithSuffixLength(5))
			if err != nil {
				t.Error(err)
			}

			got := r.Redact(tc.input)
			if tc.expected != got {
				t.Errorf("Expected '%s', but got '%s'", tc.expected, got)
			}
		})
	}
}

func TestWithReplacementText(t *testing.T) {
	cases := []testCase{
		{"", ""},
		{"a", "xxx"},
		{"ab", "xxx"},
		{"abc", "xxx"},
		{"abcd", "xxx"},
		{"abcde", "xxx"},
		{"abcdef", "abcxxx"},
		{"abcdefg", "abcxxx"},
		{"abcdefgh", "abcxxx"},

		{"abcdefghi", "abcxxxghi"}, // 9
		{"abcdefghij", "abcxxxhij"},
		{"abcdefghijk", "abcxxxijk"},
		{"abcdefghijkl", "abcxxxjkl"},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("input=%q;expected=%q; ", tc.input, tc.expected), func(t *testing.T) {
			r, err := NewFromOptions(WithReplacementText("xxx"))
			if err != nil {
				t.Error(err)
			}

			got := r.Redact(tc.input)
			if tc.expected != got {
				t.Errorf("Expected '%s', but got '%s'", tc.expected, got)
			}
		})
	}
}

func TestString(t *testing.T) {
	type testCase struct {
		r        redact.Redactor
		expected string
	}

	r1, err := NewFromOptions(WithMode(PrefixOnlyMode))
	if err != nil {
		t.Error(err)
	}

	r2, err := NewFromOptions(WithMode(SuffixOnlyMode))
	if err != nil {
		t.Error(err)
	}

	cases := []testCase{
		{New(), `{mode:"FullMode"; replacementText="[redacted]"; prefixLength=3; suffixLength=3}`},
		{r1, `{mode:"PrefixOnlyMode"; replacementText="[redacted]"; prefixLength=3; suffixLength=3}`},
		{r2, `{mode:"SuffixOnlyMode"; replacementText="[redacted]"; prefixLength=3; suffixLength=3}`},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("expected=%q; ", tc.expected), func(t *testing.T) {
			stringer := tc.r.(fmt.Stringer)
			got := stringer.String()

			if tc.expected != got {
				t.Errorf("Expected '%s', but got '%s'", tc.expected, got)
			}
		})
	}
}

func TestPrefixTooShort(t *testing.T) {
	_, err := NewFromOptions(WithPrefixLength(2))
	expected := errPrefixLengthTooShort
	if err == nil {
		t.Errorf("Expected '%s', but got '%s'", expected, err)
	}
}

func TestSuffixTooShort(t *testing.T) {
	_, err := NewFromOptions(WithSuffixLength(2))
	expected := errSuffixLengthTooShort
	if err == nil {
		t.Errorf("Expected '%s', but got '%s'", expected, err)
	}
}

func TestReplacementTextTooShort(t *testing.T) {
	_, err := NewFromOptions(WithReplacementText("x"))
	expected := errReplacementTextTooShort
	if err == nil {
		t.Errorf("Expected '%s', but got '%s'", expected, err)
	}
}

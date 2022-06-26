package url

import (
	"fmt"
	"testing"
)

func TestRedact(t *testing.T) {
	type testCase struct {
		input       string // string to be redacted
		replacement string // password replacement string
		expected    string // expected output
	}

	cases := []testCase{
		{"http://example.com", "XXXXXXXXX", "http://example.com"},
		{"http://user@example.com", "XXXXXXXXX", "http://user@example.com"},
		{"http://user:password@example.com", "XXXXXXXXX", "http://user:XXXXXXXXX@example.com"},
		{"postgresql://user:pass@host:5432/db", "XXXXXXXXX", "postgresql://user:XXXXXXXXX@host:5432/db"},
		{"mongodb://user:password@localhost", "XXXXXXXXX", "mongodb://user:XXXXXXXXX@localhost"},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("input=%q;replacement=%q;expected=%q; ", tc.input, tc.replacement, tc.expected), func(t *testing.T) {
			r := New(tc.replacement, nil)
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

func TestRedact_WithUsernameRedaction(t *testing.T) {
	type testCase struct {
		input               string // string to be redacted
		passwordReplacement string // password replacement string
		usernameReplacement string // password replacement string
		expected            string // expected output
	}

	cases := []testCase{
		{"http://example.com", "XXXXX", "ZZZZZ", "http://example.com"},
		{"http://user@example.com", "XXXXX", "ZZZZZ", "http://ZZZZZ@example.com"},
		{"http://user:password@example.com", "XXXXX", "ZZZZZ", "http://ZZZZZ:XXXXX@example.com"},
		{"postgresql://user:pass@host:5432/db", "XXXXX", "ZZZZZ", "postgresql://ZZZZZ:XXXXX@host:5432/db"},
		{"mongodb://user:password@localhost", "XXXXX", "ZZZZZ", "mongodb://ZZZZZ:XXXXX@localhost"},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("input=%q;passwordReplacement=%q;usernameReplacement=%q;expected=%q; ", tc.input, tc.passwordReplacement, tc.usernameReplacement, tc.expected), func(t *testing.T) {
			r := New(tc.passwordReplacement, &tc.usernameReplacement)
			got, err := r.Redact(tc.input)
			if err != nil {
				t.Error(err)
			}
			if tc.expected != got {
				t.Errorf("expected '%s', but got '%s'", tc.expected, got)
			}
		})
	}
}

func TestRedact_WithError(t *testing.T) {
	r := New("XXX", nil)

	_, err := r.Redact("foo.html") // invalid URL
	if err == nil {
		t.Error("expected an error, but got nil")
	}
}

func TestString(t *testing.T) {
	redactor := New("[redacted]", nil)
	stringer := redactor.(fmt.Stringer)

	expected := `{passwordReplacement="[redacted]", usernameReplacement=<nil>}`
	got := stringer.String()

	if expected != got {
		t.Errorf("expected '%s', but got '%s'", expected, got)
	}
}

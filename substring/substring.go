package substring

import (
	"fmt"
	"strings"

	"github.com/kristinjeanna/redact"
)

// SubstringRedactor is a redactor that replaces all occurrences of
// the substring in the specified string.
type SubstringRedactor struct {
	substring   string
	replacement string
}

// New returns a new SubstringRedactor.
func New(substring string, replacement string) redact.Redactor {
	return SubstringRedactor{
		substring:   substring,
		replacement: replacement,
	}
}

// Redact replaces all occurrences of the substring in the specified string.
func (r SubstringRedactor) Redact(s string) (string, error) {
	return strings.ReplaceAll(s, r.substring, r.replacement), nil
}

func (r SubstringRedactor) String() string {
	return fmt.Sprintf("{substring=%q; replacement=%q}", r.substring, r.replacement)
}

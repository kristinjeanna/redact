package blackout

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/kristinjeanna/redact"
)

// BlackoutRedactor is a redactor that replaces the characters
// of each word in a string with a specified replacement string.
type BlackoutRedactor struct {
	replacement string
}

// New returns a new BlackoutRedactor.
func New(replacement string) redact.Redactor {
	return BlackoutRedactor{replacement: replacement}
}

// Redact returns a string with each character of each word
// replaced by the redactor's replacement text.
//
// The input string is broken into a slice of strings via
// the strings.Fields function. For each word, each character is
// replaced with the replacement string. The redacted result is
// then formed by joining the blacked out words with a single space.
func (r BlackoutRedactor) Redact(s string) (string, error) {
	var b bytes.Buffer
	words := strings.Fields(s)
	count := len(words)

	for i, w := range words {
		b.WriteString(r.redactWord(w))
		if i < count-1 {
			b.WriteString(" ")
		}
	}
	return b.String(), nil
}

// redactWord replaces each character of the string with the replacement string.
func (r BlackoutRedactor) redactWord(word string) string {
	var b bytes.Buffer
	for i := 0; i < len(word); i++ {
		b.WriteString(r.replacement)
	}

	return b.String()
}

// String returns a text representation of the redactor.
func (r BlackoutRedactor) String() string {
	return fmt.Sprintf("{replacement=%q}", r.replacement)
}

package blackout

import (
	"bytes"
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
func (r BlackoutRedactor) Redact(s string) string {
	var b bytes.Buffer
	words := strings.Fields(s)
	count := len(words)

	for i, w := range words {
		b.WriteString(r.redactWord(w))
		if i < count-1 {
			b.WriteString(" ")
		}
	}
	return b.String()
}

func (r BlackoutRedactor) redactWord(word string) string {
	var b bytes.Buffer
	for i := 0; i < len(word); i++ {
		b.WriteString(r.replacement)
	}

	return b.String()
}

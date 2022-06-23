package simple

import (
	"fmt"

	"github.com/kristinjeanna/redact"
)

// SimpleRedactor is a redactor that simply replaces an
// entire string with a specified replacement string.
type SimpleRedactor struct {
	replacement string
}

// New returns a new SimpleRedactor.
func New(replacement string) redact.Redactor {
	return SimpleRedactor{replacement: replacement}
}

// Redact simply returns the replacement text for any string passed to it.
func (r SimpleRedactor) Redact(s string) string {
	return r.replacement
}

func (r SimpleRedactor) String() string {
	return fmt.Sprintf("{replacement=%q}", r.replacement)
}

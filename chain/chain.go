package chain

import (
	"fmt"

	"github.com/kristinjeanna/redact"
)

const (
	errMsgFmtRedactFail = "chain.Redact: an error occurred while redacting, %w"
)

// ChainRedactor is redactor consisting of a sequence of redactors.
type ChainRedactor []redact.Redactor

// Redact executes the redactors in the chain on the specified string.
func (r ChainRedactor) Redact(s string) (string, error) {
	redacted := s
	for _, r := range r {
		var err error
		redacted, err = r.Redact(redacted)
		if err != nil {
			return "", fmt.Errorf(errMsgFmtRedactFail, err)
		}
	}

	return redacted, nil
}

// New creates a new chain redactor from a slice of redactors.
func New(redactors []redact.Redactor) ChainRedactor {
	return ChainRedactor(redactors)
}

func (r ChainRedactor) String() string {
	return fmt.Sprintf("{redactors=%q}", []redact.Redactor(r))
}

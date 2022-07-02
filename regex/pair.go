package regex

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/kristinjeanna/redact"
	"github.com/kristinjeanna/redact/simple"
)

var (
	errRegexEmpty = errors.New("regex.NewPair: regex is required")

	errMsgFmtCompileFailure = "regex.NewPair: regex failed to compile, %w"
)

// Pair represents a regular expression and the corresponding
// expression that replaces regex matches
type Pair struct {
	redactor redact.Redactor
	regex    string
	compiled *regexp.Regexp
}

func (p Pair) String() string {
	return fmt.Sprintf("{regex=%q; redactor=%v}", p.regex, p.redactor)
}

// NewPair returns a new Pair.
func NewPair(redactor redact.Redactor, regex string) (*Pair, error) {
	if len(regex) == 0 {
		return nil, errRegexEmpty
	}
	rec, err := regexp.Compile(regex)
	if err != nil {
		return nil, fmt.Errorf(errMsgFmtCompileFailure, err)
	}

	return &Pair{
		redactor: redactor,
		regex:    regex,
		compiled: rec,
	}, nil
}

// NewPairUsingSimple returns a new Pair with a simple redactor.
func NewPairUsingSimple(replacement string, regex string) (*Pair, error) {
	return NewPair(simple.New(replacement), regex)
}

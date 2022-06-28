package regex

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/kristinjeanna/redact"
)

const (
	// Regex for an HTTP authorization header
	AuthHeaderRegex string = `(?i)(Authorization[\s]*:[\s]*).*`

	// Regex for a US Social Security Number
	SSNRegex string = `(\d{3}-?\d{2}-?\d{4})`
)

var (
	errRegexEmpty        = errors.New("regex.NewPair: regex is required")
	errRePairsSliceNil   = errors.New("regex.New: regex pairs slice must not be nil")
	errRePairsSliceEmpty = errors.New("regex.New: regex pairs slice must not be empty")

	errMsgFmtCompileFailure = "regex.NewPair: regex failed to compile, %w"
)

// Pair represents a regular expression and the corresponding
// expression that replaces regex matches
type Pair struct {
	replacement string
	regex       string
	compiled    *regexp.Regexp
}

func (p Pair) String() string {
	return fmt.Sprintf("{regex=%q; replacement=%q}", p.regex, p.replacement)
}

// NewPair returns a new Pair.
func NewPair(replacement string, regex string) (*Pair, error) {
	if len(regex) == 0 {
		return nil, errRegexEmpty
	}
	rec, err := regexp.Compile(regex)
	if err != nil {
		return nil, fmt.Errorf(errMsgFmtCompileFailure, err)
	}

	return &Pair{
		replacement: replacement,
		regex:       regex,
		compiled:    rec,
	}, nil
}

// RegexRedactor is a redactor that replaces substrings matching a given
// regular expression with a specified replacement string. Multiple pairs
// of replacement strings and regular expressions can be specified to chain
// the behavior.
type RegexRedactor struct {
	pairs []Pair
}

// New returns a new RegexRedactor.
func New(rePairs []Pair) (redact.Redactor, error) {
	if rePairs == nil {
		return nil, errRePairsSliceNil
	}

	if len(rePairs) == 0 {
		return nil, errRePairsSliceEmpty
	}

	return RegexRedactor{pairs: rePairs}, nil
}

// Redact simply returns the replacement text for any string passed to it.
func (r RegexRedactor) Redact(s string) (string, error) {
	src := s

	for _, pair := range r.pairs {
		src = pair.compiled.ReplaceAllString(src, pair.replacement)
	}

	return src, nil
}

// String returns a text representation of the redactor.
func (r RegexRedactor) String() string {
	return fmt.Sprintf("{pairs=%q}", r.pairs)
}

package regex

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/kristinjeanna/redact"
)

var (
	errRegexEmpty        = errors.New("regex is required")
	errRePairsSliceNil   = errors.New("regex pairs slice must not be nil")
	errRePairsSliceEmpty = errors.New("regex slice must not be empty")
)

// Pair represents a regular expression and the
// corresponding replacement expression.
type Pair struct {
	replacement string
	regex       string
	compiled    *regexp.Regexp
}

func (p Pair) String() string {
	return fmt.Sprintf("{regex=%q; replacement=%q}", p.regex, p.replacement)
}

// NewPair returns new Pair.
func NewPair(replacement string, regex string) (*Pair, error) {
	if len(regex) == 0 {
		return nil, errRegexEmpty
	}
	rec, err := regexp.Compile(regex)
	if err != nil {
		return nil, err
	}

	return &Pair{
		replacement: replacement,
		regex:       regex,
		compiled:    rec,
	}, nil
}

// RegexRedactor is a redactor that replaces matching
// substrings with a specified replacement string.
type RegexRedactor struct {
	pairs []Pair
}

func (r RegexRedactor) String() string {
	return fmt.Sprintf("{pairs=%q}", r.pairs)
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

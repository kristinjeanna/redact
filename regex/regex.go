package regex

import (
	"errors"
	"fmt"

	"github.com/kristinjeanna/redact"
	"github.com/kristinjeanna/redact/simple"
)

const (
	// Regex for an HTTP authorization header
	AuthHeaderRegex string = `(?i)(Authorization[\s]*:[\s]*).*`

	// Regex for a US Social Security Number
	SSNRegex string = `(\d{3}-?\d{2}-?\d{4})`
)

var (
	errRePairsSliceNil         = errors.New("regex.New: regex pairs slice must not be nil")
	errRePairsSliceEmpty       = errors.New("regex.New: regex pairs slice must not be empty")
	errRegexMatchesReplacement = errors.New("regex.RegexRedactor.Redact: regex must not match replacement text returned from the pair's redactor")

	errMsgFmtRedactFailure = "regex.RegexRedactor.Redact: error while redacting, %w"
)

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
		switch pair.redactor.(type) {
		case simple.SimpleRedactor:
			repl, _ := pair.redactor.Redact("")
			src = pair.compiled.ReplaceAllString(src, repl)
		default:
			indexes := pair.compiled.FindAllStringIndex(src, -1)
			count := len(indexes) // infinite loop prevention

			for indexes != nil {
				loc := indexes[0]
				found := src[loc[0]:loc[1]]
				repl, err := pair.redactor.Redact(found)
				if err != nil {
					return "", fmt.Errorf(errMsgFmtRedactFailure, err)
				}

				src = replaceSubstring(src, repl, loc[0], loc[1])
				indexes = pair.compiled.FindAllStringIndex(src, -1)

				// the length of indexes should only decrease with each iteration
				// if it doesn't, that means the regex matched on the replacement text;
				// return error here to prevent infinite loop
				if len(indexes) >= count {
					return "", errRegexMatchesReplacement
				}
				count = len(indexes)
			}
		}
	}

	return src, nil
}

// String returns a text representation of the redactor.
func (r RegexRedactor) String() string {
	return fmt.Sprintf("{pairs=%v}", r.pairs)
}

func replaceSubstring(in string, repl string, from int, to int) string {
	inRunes := []rune(in)
	replRunes := []rune(repl)
	outRunes := make([]rune, 0)

	if from != 0 {
		outRunes = append(outRunes, inRunes[:from]...)
	}

	outRunes = append(outRunes, replRunes...)
	outRunes = append(outRunes, inRunes[to:]...)

	return string(outRunes)
}

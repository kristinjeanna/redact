package middle

import (
	"fmt"

	"github.com/kristinjeanna/redact"
)

type Mode int8

const (
	FullMode Mode = iota
	PrefixOnlyMode
	SuffixOnlyMode
)

// String returns a text representation of the mode.
func (m Mode) String() string {
	switch m {
	case PrefixOnlyMode:
		return "PrefixOnlyMode"
	case SuffixOnlyMode:
		return "SuffixOnlyMode"
	case FullMode:
		fallthrough
	default:
		return "FullMode"
	}
}

const (
	defaultPrefixLength    uint   = 3
	defaultSuffixLength    uint   = 3
	defaultReplacementText string = "[redacted]"
	defaultMode            Mode   = FullMode

	replacementTextMinLength = 3
	prefixLengthMinimum      = 3
	suffixLengthMinimum      = 3
)

var (
	errReplacementTextTooShort = fmt.Errorf("middle.NewFromOptions: length of replacement text must not be less than %d", replacementTextMinLength)
	errPrefixLengthTooShort    = fmt.Errorf("middle.NewFromOptions: prefix length must not be less than %d", prefixLengthMinimum)
	errSuffixLengthTooShort    = fmt.Errorf("middle.NewFromOptions: suffix length must not be less than %d", suffixLengthMinimum)
)

// MiddleRedactor is a redactor that replaces the middle contents of a string with a replacement string, leaving a prefix of unredacted characters and a suffix of unredacted characters if the input string is long enough. For shorter input strings, the redactor uses only a prefix or suffix or just the replacement string itself.
type MiddleRedactor struct {
	mode            Mode
	prefixLength    uint
	suffixLength    uint
	replacementText string
}

// New creates a new MiddleRedactor with a default configuration.
func New() redact.Redactor {
	return MiddleRedactor{
		mode:            defaultMode,
		prefixLength:    defaultPrefixLength,
		suffixLength:    defaultSuffixLength,
		replacementText: defaultReplacementText,
	}
}

// NewFromOptions creates a new MiddleRedactor with the provided options.
func NewFromOptions(opts ...Option) (redact.Redactor, error) {
	m := New().(MiddleRedactor)
	for _, o := range opts {
		o(&m)
	}

	if m.prefixLength < prefixLengthMinimum {
		return nil, errPrefixLengthTooShort
	}

	if m.suffixLength < suffixLengthMinimum {
		return nil, errSuffixLengthTooShort
	}

	if len(m.replacementText) < replacementTextMinLength {
		return nil, errReplacementTextTooShort
	}

	return m, nil
}

// Redact redacts the input string and returns the result.
func (m MiddleRedactor) Redact(s string) (string, error) {
	length := uint(len(s))
	lengthReplText := uint(len(m.replacementText))
	minLength := m.prefixLength + lengthReplText + m.suffixLength

	if length == 0 {
		return s, nil
	}

	// long enough for prefix & suffix
	if length >= minLength {
		if m.mode == FullMode {
			return s[:m.prefixLength] + m.replacementText + s[length-m.suffixLength:], nil
		}
		if m.mode == PrefixOnlyMode {
			return s[:m.prefixLength] + m.replacementText, nil
		}
		if m.mode == SuffixOnlyMode {
			return m.replacementText + s[length-m.suffixLength:], nil
		}
	}

	// not long enough for both prefix & suffix
	if length > lengthReplText {
		showPrefix := m.mode == PrefixOnlyMode || m.mode == FullMode

		if showPrefix && length >= m.prefixLength+lengthReplText {
			return s[:m.prefixLength] + m.replacementText, nil // redact with prefix
		}

		if !showPrefix && length >= m.suffixLength+lengthReplText {
			return m.replacementText + s[length-m.suffixLength:], nil // redact with suffix
		}
	}

	// length <= lengthReplText but not 0
	return m.replacementText, nil
}

// String returns a text representation of the redactor.
func (m MiddleRedactor) String() string {
	return fmt.Sprintf(
		"{mode:%q; replacementText=%q; prefixLength=%d; suffixLength=%d}",
		m.mode,
		m.replacementText,
		m.prefixLength,
		m.suffixLength,
	)
}

// Option defines options for creating new middle redactors.
type Option func(*MiddleRedactor)

/*
WithMode sets the mode for the redactor. Default is "FullMode".
*/
func WithMode(mode Mode) Option {
	return func(m *MiddleRedactor) {
		m.mode = mode
	}
}

/*
WithPrefixLength sets the amount of text from the start of the input
string to allow in the redacted result. Default is 3 characters.

Must be no less than 3.
*/
func WithPrefixLength(prefixLength uint) Option {
	return func(m *MiddleRedactor) {
		m.prefixLength = prefixLength
	}
}

/*
WithSuffixLength sets the amount of text from the end of the input
string to allow in the redacted result. Default is 3 characters.

Must be no less than 3.
*/
func WithSuffixLength(suffixLength uint) Option {
	return func(m *MiddleRedactor) {
		m.suffixLength = suffixLength
	}
}

/*
WithReplacementText sets the replacement string. Default is "[redacted]".

Must be a minimum of 3 characters long.
*/
func WithReplacementText(replacementText string) Option {
	return func(m *MiddleRedactor) {
		m.replacementText = replacementText
	}
}

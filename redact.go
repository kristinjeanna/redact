package redact

// Redactor is the common interface implemented by all redactors.
type Redactor interface {

	// Redact redacts the input string and returns the result.
	Redact(s string) (string, error)
}

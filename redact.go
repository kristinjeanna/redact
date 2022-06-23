package redact

// Redactor is the common interface implemented by all redactors.
type Redactor interface {
	Redact(s string) string
}

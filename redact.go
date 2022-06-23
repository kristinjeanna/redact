package redact

type Redactor interface {
	Redact(s string) string
}

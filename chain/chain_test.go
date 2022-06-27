package chain

import (
	"testing"

	"github.com/kristinjeanna/redact"
	"github.com/kristinjeanna/redact/substring"
	"github.com/kristinjeanna/redact/url"
)

func TestNewPair_err(t *testing.T) {
	substringRedactor := substring.New("string", "XXXXX")
	urlRedactor := url.New("http", nil)
	chainRedactor := New([]redact.Redactor{substringRedactor, urlRedactor})

	sampleString := "this is a string"
	_, err := chainRedactor.Redact(sampleString)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}

func TestString(t *testing.T) {
	substringRedactor := substring.New("string", "XXXXX")
	chainRedactor := New([]redact.Redactor{substringRedactor})

	expected := `{redactors=["{substring=\"string\"; replacement=\"XXXXX\"}"]}`
	got := chainRedactor.String()

	if expected != got {
		t.Errorf("Expected '%s', but got '%s'", expected, got)
	}
}

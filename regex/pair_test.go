package regex

import (
	"fmt"
	"testing"
)

func TestNewPairUsingSimple_err(t *testing.T) {
	_, err := NewPairUsingSimple("test", "")
	if err == nil {
		t.Errorf("Expected '%v', but got nil", errRegexEmpty)
	}
	if err != errRegexEmpty {
		t.Errorf("Expected '%v', but got '%v'", errRegexEmpty, err)
	}

	_, err = NewPairUsingSimple("test", "foo+++++")
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}

func TestString(t *testing.T) {
	pair, _ := NewPairUsingSimple("[redacted]", "test")
	redactor, _ := New([]Pair{*pair})
	reRedactor := redactor.(fmt.Stringer)

	expected := `{pairs=[{regex="test"; redactor={replacement="[redacted]"}}]}`
	got := reRedactor.String()

	if expected != got {
		t.Errorf("Expected '%s', but got '%s'", expected, got)
	}
}

package middle

import (
	"fmt"
	"log"
)

const (
	sampleString = "this string contains sensitive information"
)

func ExampleMiddleRedactor() {
	redactor := New()

	result, err := redactor.Redact(sampleString)
	if err != nil {
		log.Fatalf("an error occurred while redacting: %s", err)
	}

	fmt.Println(result)
	// Output: thi[redacted]ion
}

func ExampleNewFromOptions() {
	redactor, err := NewFromOptions(
		WithMode(PrefixOnlyMode),
		WithReplacementText("XXXXX"),
		WithPrefixLength(8),
	)
	if err != nil {
		log.Fatalf("an error occurred while creating redactor: %s", err)
	}

	result, err := redactor.Redact(sampleString)
	if err != nil {
		log.Fatalf("an error occurred while redacting: %s", err)
	}

	fmt.Println(result)
	// Output: this strXXXXX
}

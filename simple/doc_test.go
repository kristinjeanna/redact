package simple

import (
	"fmt"
	"log"
)

func ExampleSimpleRedactor() {
	redactor := New("[redacted]")
	sampleString := "this string contains sensitive information"

	result, err := redactor.Redact(sampleString)
	if err != nil {
		log.Fatalf("an error occurred while redacting: %s", err)
	}

	fmt.Println(result)
	// Output: [redacted]
}

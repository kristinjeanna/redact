package url

import (
	"fmt"
	"log"
)

func ExampleURLRedactor() {
	redactor := New("REDACTED", nil)
	sampleString := "mysql://user:foobar@localhost:3306"

	result, err := redactor.Redact(sampleString)
	if err != nil {
		log.Fatalf("an error occurred while redacting: %s", err)
	}

	fmt.Println(result)
	// Output: mysql://user:REDACTED@localhost:3306
}

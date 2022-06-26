package substring

import (
	"fmt"
	"log"
)

func Example() {
	redactor := New("contains sensitive", "XXXXX")
	sampleString := "this string contains sensitive information"

	result, err := redactor.Redact(sampleString)
	if err != nil {
		log.Fatalf("an error occurred while redacting: %s", err)
	}

	fmt.Println(result)
	// Output: this string XXXXX information
}

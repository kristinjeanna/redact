package simple

import "fmt"

func Example() {
	redactor := New("[redacted]")
	sampleString := "this string contains sensitive information"

	fmt.Println(redactor.Redact(sampleString))
	// Output: [redacted]
}

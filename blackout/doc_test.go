package blackout

import "fmt"

func Example() {
	redactor := New("█")
	sampleString := "this string contains sensitive information"

	fmt.Println(redactor.Redact(sampleString))
}

package substring

import "fmt"

func Example() {
	redactor := New("contains sensitive", "XXXXX")
	sampleString := "this string contains sensitive information"

	fmt.Println(redactor.Redact(sampleString))
	// Output: this string XXXXX information
}

package middle

import "fmt"

const (
	sampleString = "this string contains sensitive information"
)

func Example() {
	redactor := New()

	fmt.Println(redactor.Redact(sampleString))
	// Output: thi[redacted]ion
}

func Example_usingOptions() {
	redactor, err := NewFromOptions(
		WithMode(PrefixOnlyMode),
		WithReplacementText("XXXXX"),
		WithPrefixLength(8),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(redactor.Redact(sampleString))
	// Output: this strXXXXX
}

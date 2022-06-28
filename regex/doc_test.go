package regex

import (
	"fmt"
	"log"
)

const sampleString = "this string contains sensitive information"

func ExampleRegexRedactor() {
	// a regex.Pair holds the regex and the replacement string for matches
	pair, err := NewPair("X", "[is]")
	if err != nil {
		log.Fatalf("an error occurred while create regex replacement pair: %s", err)
	}

	// the redactor is constructed with a slice of Pairs
	redactor, err := New([]Pair{*pair})
	if err != nil {
		log.Fatalf("an error occurred while creating redactor: %s", err)
	}

	result, err := redactor.Redact(sampleString)
	if err != nil {
		log.Fatalf("an error occurred while redacting: %s", err)
	}

	fmt.Println(result)
	// Output: thXX XtrXng contaXnX XenXXtXve XnformatXon
}

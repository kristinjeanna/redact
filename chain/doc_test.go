package chain

import (
	"fmt"
	"log"

	"github.com/kristinjeanna/redact"
	"github.com/kristinjeanna/redact/regex"
	"github.com/kristinjeanna/redact/substring"
)

func ExampleChainRedactor() {
	substringRedactor := substring.New("contains", "HIDES")

	regexPair, err := regex.NewPairUsingSimple(" [redacted] ", `(?i)\ss\w*\s`)
	if err != nil {
		log.Fatalf("an error occurred while creating regex pair: %s", err)
	}
	regexRedactor, err := regex.New([]regex.Pair{*regexPair})
	if err != nil {
		log.Fatalf("an error occurred while creating regex redactor: %s", err)
	}

	chainRedactor := New([]redact.Redactor{substringRedactor, regexRedactor})
	sampleString := "this string contains sensitive information"

	result, err := chainRedactor.Redact(sampleString)
	if err != nil {
		log.Fatalf("an error occurred while redacting: %s", err)
	}

	fmt.Println(result)
	// Output: this [redacted] HIDES [redacted] information
}

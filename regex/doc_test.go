package regex

import (
	"fmt"
	"log"
)

func ExampleRegexRedactor() {
	// a regex.Pair holds the regex and the replacement redactor for matches
	pair, err := NewPairUsingSimple("X", "[is]")
	if err != nil {
		log.Fatalf("an error occurred while creating the regex pair: %s", err)
	}

	// the redactor is constructed with a slice of Pairs
	redactor, err := New([]Pair{*pair})
	if err != nil {
		log.Fatalf("an error occurred while creating redactor: %s", err)
	}

	sampleString := "this string contains sensitive information"
	result, err := redactor.Redact(sampleString)
	if err != nil {
		log.Fatalf("an error occurred while redacting: %s", err)
	}

	fmt.Println(result)
	// Output: thXX XtrXng contaXnX XenXXtXve XnformatXon
}

func ExampleRegexRedactor_usingCaptureGroup() {
	// a regex.Pair holds the regex and the replacement redactor for matches
	pair, err := NewPairUsingSimple("${1}XXXX", "(b[aA][rRzZ])")
	if err != nil {
		log.Fatalf("an error occurred while creating the regex pair: %s", err)
	}

	// the redactor is constructed with a slice of Pairs
	redactor, err := New([]Pair{*pair})
	if err != nil {
		log.Fatalf("an error occurred while creating redactor: %s", err)
	}

	sampleString := "foo bar baz"
	result, err := redactor.Redact(sampleString)
	if err != nil {
		log.Fatalf("an error occurred while redacting: %s", err)
	}

	fmt.Println(result)
	// Output: foo barXXXX bazXXXX
}

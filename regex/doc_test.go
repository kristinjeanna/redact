package regex

import "fmt"

const sampleString = "this string contains sensitive information"

func Example() {
	pair, err := NewPair("X", "[is]")
	if err != nil {
		panic(err)
	}

	redactor, err := New([]Pair{*pair})
	if err != nil {
		panic(err)
	}

	fmt.Println(redactor.Redact(sampleString))
	// Output: thXX XtrXng contaXnX XenXXtXve XnformatXon
}

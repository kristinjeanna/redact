package regex

import "fmt"

func Example() {
	pair, err := NewPair("X", "[is]")
	if err != nil {
		panic(err)
	}

	redactor, err := New([]Pair{*pair})
	sampleString := "this string contains sensitive information"

	fmt.Println(redactor.Redact(sampleString))
	// Output: thXX XtrXng contaXnX XenXXtXve XnformatXon
}

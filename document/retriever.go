package document

import "github.com/PuerkitoBio/goquery"

// retriever is a function that implements the same interface as goquery
type retriever func(location string) (*goquery.Document, error)

// documentFn exists to make testing possible without resorting to hardcoded function.
type documentFn func(string) (*goquery.Document, error)

// genRetrieve generates a new retriever function, useful for testing.
func genRetrieve(fn documentFn) retriever {
	if fn == nil {
		fn = goquery.NewDocument
	}
	return func(location string) (*goquery.Document, error) {
		return fn(location)
	}
}

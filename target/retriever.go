package target

import "github.com/PuerkitoBio/goquery"

// retriever is a function that implements the same interface as goquery
type retriever func(Location) (*goquery.Document, error)

// documentFn exists to make testing possible without resorting to hardcoded function.
type documentFn func(string) (*goquery.Document, error)

// generates a new retriever function, useful for testing.
func GenRetrieve(fn documentFn) retriever {
	if fn == nil {
		fn = goquery.NewDocument
	}
	return func(l Location) (*goquery.Document, error) {
		return fn(l.String())
	}
}

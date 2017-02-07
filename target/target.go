package target

// go:generate gen

import (
	"github.com/PuerkitoBio/goquery"
	"net/url"
)

// retriever is a function that implements the same interface as goquery
type retriever func(Location) (*goquery.Document, error)

// Config is the global configuration object for target
var Config = struct {
	// Retrieve is the global document retriever, override for testing with GenRetrieve
	Retrieve retriever
}{
	Retrieve: GenRetrieve(nil),
}

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

// Location represents a validated url.
// +gen set slice:"DistinctBy"
type Location struct {
	*url.URL
}

// Retrieve calls the global document retriever on a location
func (l Location) Retrieve() (*goquery.Document, error) {
	return Config.Retrieve(l)
}

// Parse makes a Location from a reference Location.
func (l Location) Parse(textURL string) (Location, error) {
	u, err := l.URL.Parse(textURL)
	if err != nil {
		return Location{}, err
	}
	return Location{u}, nil
}

var baseLocation = &Location{&url.URL{}}

// New makes a Location from a raw url string.
func New(textURL string) (Location, error) {
	return baseLocation.Parse(textURL)
}

// ParseAll converts arguments into a LocationSlice of distinct values
func (l Location) ParseAll(args []string) (LocationSlice, error) {
	ls := make(LocationSlice, 0, len(args))
	for i := range args {
		if loc, err := l.Parse(args[i]); err == nil {
			ls = append(ls, loc)
		} else {
			return LocationSlice{}, err
		}
	}
	return ls.DistinctByURL(), nil
}

// ParseAll without creating a Location
func ParseAll(args []string) (LocationSlice, error) {
	return baseLocation.ParseAll(args)
}

func (ls LocationSlice) DistinctByURL() LocationSlice {
	urlPredicate := func(a, b Location) bool {
		return a.URL.String() == b.URL.String()
	}
	return ls.DistinctBy(urlPredicate)
}

package target

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

// ParseAll converts arguments into a list of URLs.
func (l Location) ParseAll(args []string) ([]Location, error) {
	ts := make([]Location, 0, len(args))
	for i := range args {
		if targ, err := l.Parse(args[i]); err == nil {
			ts = append(ts, targ)
		} else {
			return []Location{}, err
		}
	}
	return ts, nil
}

var baseLocation = &Location{&url.URL{}}

// New makes a Location from a raw url string.
func New(textURL string) (Location, error) {
	return baseLocation.Parse(textURL)
}

// ParseAll without creating a Location
func ParseAll(args []string) ([]Location, error) {
	return baseLocation.ParseAll(args)
}

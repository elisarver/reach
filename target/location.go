package target

import (
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

// Location represents a validated url.
// +gen set slice:"DistinctBy"
type Location struct {
	*url.URL
}

var emptyLocation = Location{&url.URL{}}

// NewLocation makes a Location from a raw url string.
func NewLocation(textURL string) (Location, error) {
	return emptyLocation.Parse(textURL)
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

// ParseAll without creating a Location
func ParseLocations(args ...string) (LocationSlice, error) {
	return emptyLocation.ParseLocations(args...)
}

// ParseAll converts arguments into a LocationSlice of distinct values
func (l Location) ParseLocations(args ...string) (LocationSlice, error) {
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

package target

import (
	"fmt"
	"net/url"
)

// Location represents a validated url.
// +gen set slice:"DistinctBy"
type Location interface {
	Parse(string) (Location, error)
	ParseAll(args ...string) (LocationSlice, error)
	fmt.Stringer
}

type location struct {
	u *url.URL
}

var emptyLocation = location{&url.URL{}}

// Parse makes a Location from a raw url string.
var Parse = emptyLocation.Parse

// ParseAll processes a list of strings
var ParseAll = emptyLocation.ParseAll

// Parse makes a Location from a reference Location.
func (l location) Parse(textURL string) (Location, error) {
	u, err := l.u.Parse(textURL)
	return &location{u}, err
}

// ParseAll converts arguments into a LocationSlice of distinct values
func (l location) ParseAll(args ...string) (LocationSlice, error) {
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

func (l location) String() string {
	return l.u.String()
}

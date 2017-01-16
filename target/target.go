package target

import (
	"net/url"
)

// Location represents a validated url.
type Location struct {
	*url.URL
}

var baseTarget = &Location{&url.URL{}}

// New makes a Location from a raw url string.
func New(textURL string) (Location, error) {
	return baseTarget.Parse(textURL)
}

// Parse makes a Location from a reference Location.
func (t Location) Parse(textURL string) (Location, error) {
	u, err := t.URL.Parse(textURL)
	if err != nil {
		return Location{}, err
	}
	return Location{u}, nil
}

// ParseAll without creating a Location
func ParseAll(args []string) ([]Location, error) {
	return baseTarget.ParseAll(args)
}

// ParseAll converts arguments into a list of URLs.
func (t Location) ParseAll(args []string) ([]Location, error) {
	ts := make([]Location, 0, len(args))
	for i := range args {
		if targ, err := t.Parse(args[i]); err == nil {
			ts = append(ts, targ)
		} else {
			return []Location{}, err
		}
	}
	return ts, nil
}

package target

import (
	"net/url"
)

// Target represents a validated url.
type Target struct {
	*url.URL
}

var baseTarget = &Target{&url.URL{}}

// NewTarget makes a Target from a raw url string.
func NewTarget(textURL string) (Target, error) {
	return baseTarget.Parse(textURL)
}

// Parse makes a Target from a reference Target.
func (t Target) Parse(textURL string) (Target, error) {
	u, err := t.URL.Parse(textURL)
	if err != nil {
		return Target{}, err
	}
	return Target{u}, nil
}

// ParseAll without creating a Target
func ParseAll(args []string) ([]Target, error) {
	return baseTarget.ParseAll(args)
}

// ParseAll converts arguments into a list of URLs.
func (t Target) ParseAll(args []string) ([]Target, error) {
	ts := make([]Target, 0, len(args))
	for i := range args {
		if targ, err := t.Parse(args[i]); err == nil {
			ts = append(ts, targ)
		} else {
			return []Target{}, err
		}
	}
	return ts, nil
}

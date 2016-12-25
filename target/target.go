package target

import (
	"net/url"
)

// Target represents a validated url.
type Target struct {
	*url.URL
}

// NewTarget makes a Target from a raw url string.
func NewTarget(rawurl string) (Target, error) {
	var t = &Target{&url.URL{}}
	return t.Parse(rawurl)
}

// Parse without creating a Target.
func Parse(rawurl string) (Target, error) {
	return Target{&url.URL{}}.Parse(rawurl)
}

// Parse makes a Target from a reference Target.
func (t Target) Parse(rawurl string) (Target, error) {
	u, err := t.URL.Parse(rawurl)
	if err != nil {
		return Target{}, err
	}
	return Target{u}, nil
}

// ParseAll without creating a Target
func ParseAll(args []string) ([]Target, error) {
	return Target{&url.URL{}}.ParseAll(args)
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

package main

import (
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

// A Target is a validated url for the purposes of a HTTP dial.
type Target struct {
	*url.URL
}

// NewTarget generates a Target IFF
func NewTarget(rawurl string) (Target, error) {
	var t Target = Target{&url.URL{}}
	return t.Parse(rawurl)
}

func (t Target) Parse(rawurl string) (Target, error) {
	u, err := t.URL.Parse(rawurl)
	return Target{u}, err
}

// Reach function retrieves a goquery Document for a URL
func Reach(t Target) (Response, error) {
	return goquery.NewDocument(t.String())
}

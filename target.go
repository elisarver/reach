package main

import (
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type Target struct {
	*url.URL
}

func NewTarget(s string) (Target, error) {
	if u, err := url.ParseRequestURI(s); err == nil {
		return Target{u}, nil
	} else {
		return Target{&url.URL{}}, err
	}
}

// Reach function retrieves a goquery Document for a URL
func Reach(t Target) (Response, error) {
	return goquery.NewDocument(t.String())
}

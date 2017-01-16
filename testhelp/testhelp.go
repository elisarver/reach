// Package testhelp are re-usable test helpers for all modules in reach.
package testhelp

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strings"
	"testing"
)

// Reporter is this package's name for the standard Formatter.
type Reporter func(message string, values ...interface{})

// Errmsg returns a function that works like fmt.Errorf, but logs the test information
// and runs it through t's Errorf.
func Errmsg(t *testing.T, instance string) Reporter {
	return func(message string, values ...interface{}) {
		t.Errorf(fmt.Sprintf("%s: %s", instance, message), values...)
	}
}

// NewURL creates a new url, and fails the test if it's invalid.
func NewURL(t *testing.T, textURL string) *url.URL {
	u, err := url.Parse(textURL)
	if err != nil {
		t.Error(err)
	}
	return u
}

// GenDoc generates a goquery.Document from a raw HTML string.
func GenDoc(t *testing.T, s string) *goquery.Document {
	var (
		res *goquery.Document
		err error
	)
	if res, err = goquery.NewDocumentFromReader(strings.NewReader(s)); err != nil {
		t.Error(err)
	}
	return res
}

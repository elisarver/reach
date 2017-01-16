// Package testhelp are re-usable test helpers for all modules in reach.
package testhelp

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

type R struct {
	t        *testing.T
	instance string
}

// Reporter returns an R that reports on a test instance in its logging.
func Reporter(t *testing.T, instance string) *R {
	return &R{t: t, instance: instance}
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

func (r *R) Errorf(format string, values ...interface{}) {
	r.t.Errorf("%s:%s", r.instance, fmt.Sprintf(format, values...))
}

func (r *R) Compare(expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		r.Errorf("Expected:\n\t%s\nGot:\n\t%s\n", expected, actual)
	}
}

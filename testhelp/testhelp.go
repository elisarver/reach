package testhelp

import (
	"fmt"
	"net/url"
	"testing"
)

type Reporter func(message string, values ...interface{})

func Errmsg(t *testing.T, instance string) func(message string, values ...interface{}) {
	return func(message string, values ...interface{}) {
		t.Errorf(fmt.Sprintf("%s: %s", instance, message), values...)
	}
}

func NewURL(t *testing.T, u string) *url.URL {
	url, err := url.Parse(u)
	if err != nil {
		t.Error(err)
	}
	return url
}

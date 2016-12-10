package target

import (
	"testing"

	"github.com/elisarver/reach/testhelp"
)

type InputExpected map[string]struct {
	input    string
	expected interface{}
}

func TestNewTarget(t *testing.T) {
	tests := InputExpected{
		"empty":         {input: "", expected: ""},
		"root relative": {input: "/", expected: "/"},
		"domain":        {input: "http://foo.bar/", expected: "http://foo.bar/"},
		"full path":     {input: "http://foo.bar/a/b", expected: "http://foo.bar/a/b"},
		"parse error":   {input: "http://foo bar/a/b", expected: "parse http://foo bar/a/b: invalid character \" \" in host name"},
	}
	for instance, test := range tests {
		reporter := testhelp.Errmsg(t, instance)
		fn := func() (*Target, error) {
			return NewTarget(test.input)
		}
		reportOn(reporter, fn, test.expected)
	}
}

func TestParse(t *testing.T) {
	u := testhelp.NewURL(t, "http://foo.bar/")
	tests := InputExpected{
		"/baz": {input: "/baz", expected: "http://foo.bar/baz"},
	}
	for instance, test := range tests {
		reporter := testhelp.Errmsg(t, instance)
		t := &Target{URL: u}
		fn := func() (*Target, error) {
			return t.Parse(test.input)
		}
		reportOn(reporter, fn, test.expected)
	}
}

type Stringer interface {
	String() string
}

type Targeter func() (*Target, error)

func reportOn(reporter testhelp.Reporter, fn Targeter, expected interface{}) {
	target, err := fn()
	var actual interface{}
	if err != nil {
		actual = err.Error()
	} else {
		actual = target.String()
	}
	if actual != expected {
		reporter("expeced %q, got %q", expected, actual)
	}
}

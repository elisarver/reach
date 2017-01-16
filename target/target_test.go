package target

import (
	"testing"

	"github.com/elisarver/reach/testhelp"
	"reflect"
)

type InputExpected map[string]struct {
	input    interface{}
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
		fn := func() (Location, error) {
			return New(test.input.(string))
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
		tgt := &Location{URL: u}
		fn := func() (Location, error) {
			return tgt.Parse(test.input.(string))
		}
		reportOn(reporter, fn, test.expected)
	}
}

func reportOn(reporter testhelp.Reporter, fn func() (Location, error), expected interface{}) {
	tgt, err := fn()
	var actual interface{}
	if err != nil {
		actual = err.Error()
	} else {
		actual = tgt.String()
	}
	if actual != expected {
		reporter("expeced %q, got %q", expected, actual)
	}
}

func TestParseAll(t *testing.T) {
	_, err := ParseAll([]string{})
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	result, err := ParseAll([]string{"http://google.com/"})
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	expected := []Location{{testhelp.NewURL(t, "http://google.com/")}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("results don't match. expected %v, got %v", expected, result)
	}
}

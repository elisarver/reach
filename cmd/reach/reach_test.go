package main

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/elisarver/reach/target"
)

func TestArgTargets(t *testing.T) {
	type pair struct {
		targets []target.Target
		err     error
	}

	ustr := []string{
		"http://example.com",
		"http://blog.example.com/",
		"://malformed@/",
	}

	l := len(ustr)

	urls := make([]*url.URL, l)
	errs := make([]error, l)

	for i := 0; i < l; i++ {
		urls[i], errs[i] = url.ParseRequestURI(ustr[i])
	}

	argss := [][]string{
		{},
		{ustr[0]},
		{ustr[0], ustr[1]},
		{ustr[2]},
	}

	expecteds := []pair{
		{[]target.Target{}, errors.New("please supply at least one URL")},
		{[]target.Target{{URL: urls[0]}}, nil},
		{[]target.Target{{URL: urls[0]}, {URL: urls[1]}}, nil},
		{[]target.Target{}, errs[2]},
	}

	for i := range argss {
		a, e := argTargets(argss[i])
		act := pair{a, e}
		exp := expecteds[i]

		shouldError := exp.err != nil
		didError := act.err != nil

		if shouldError != didError {
			if shouldError {
				t.Errorf("expected error %q ", exp.err)
			} else {
				t.Errorf("did not expect error %q", act.err)
			}
		}

		if shouldError && didError {
			if act.err.Error() != exp.err.Error() {
				t.Errorf("expected error %q, got %q", exp.err, act.err)
			}
		}

		if len(exp.targets) != len(act.targets) {
			t.Errorf("size of exp and act targets mismatch on:\n\texp: %q\n\tact: %q", exp, act)
		}
		for j := range exp.targets {
			expVal := exp.targets[j].String()
			actVal := act.targets[j].String()
			if expVal != actVal {
				t.Errorf("expected actual value %q to equal expected value %q",
					act.targets[j], exp.targets[j])
			}
		}
	}
}

// reachTargets(ts []Target, tagName string, reachFn func(string) (*goquery.Document, error)) []string {
func TestReachTargets(t *testing.T) {
	reachFnSuccess := func(_ string) (*goquery.Document, error) {
		r := strings.NewReader("<html><body><a href='http://foo.bar/'>site</a></body></html>")
		return goquery.NewDocumentFromReader(r)
	}
	u, _ := target.NewTarget("http://foo.bar/")
	us := []target.Target{u}
	actual, err := reachTargets(us, "a", reachFnSuccess)
	if err != nil {
		t.Errorf("test didn't expect %s", err)
	}
	expected := []string{"http://foo.bar/"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestDropEmpties(t *testing.T) {
	input := []string{"not empty", "", "also not empty"}
	expected := fmt.Sprintf("%q", []string{"not empty", "also not empty"})
	actual := fmt.Sprintf("%q", dropEmpties(input))
	if expected != actual {
		t.Errorf("dropEmpties failed!\nExpected:\n\t%s\nGot:\n\t%s\n",
			expected, actual)
	}
}

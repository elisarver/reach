package main

import (
	"fmt"
	"net/url"
	"testing"
)

func TestDropEmpties(t *testing.T) {
	input := []string{"not empty", "", "also not empty"}
	expected := fmt.Sprintf("%q", []string{"not empty", "also not empty"})
	actual := fmt.Sprintf("%q", dropEmpties(input))
	if expected != actual {
		t.Errorf("dropEmpties failed!\nExpected:\n\t%s\nGot:\n\t%s\n",
			expected, actual)
	}
}

func TestChkURL(t *testing.T) {
	input := []string{"http://www.google.com/", ""}

	type pair struct {
		shouldError bool
		url         *url.URL
	}

	expected := []pair{
		pair{false, &url.URL{Scheme: "http", Host: "www.google.com", Path: "/"}},
		pair{true, nil},
	}

	for i := range input {
		inp := input[i]
		exp := expected[i]
		act, acte := chkURL(inp)

		if exp.shouldError {
			if acte == nil {
				t.Error("expected error to not be nil.")
			}
			if act != nil {
				t.Errorf("expected url to be nil, was %q", act)
			}
		} else {
			if acte != nil {
				t.Errorf("expected error to be nil, was %q", acte)
			}
			if act == nil {
				t.Error("expected url to have value, did not")
			}
		}
	}
}

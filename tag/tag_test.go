package tag

import (
	"testing"

	"reflect"

	"github.com/elisarver/reach/testhelp"
)

// tags represent finders
func TestFromMultiSpec(t *testing.T) {
	tests := map[string]struct {
		name     string
		expected []*Description
	}{
		"empty":         {name: "", expected: []*Description{{Name: "", Attribute: "", CSSSelector: ""}}},
		"a automatic":   {name: "a", expected: []*Description{{Name: "a", Attribute: "href", CSSSelector: "a[href]"}}},
		"img automatic": {name: "img", expected: []*Description{{Name: "img", Attribute: "src", CSSSelector: "img[src]"}}},
		"full spec":     {name: "meta:name", expected: []*Description{{Name: "meta", Attribute: "name", CSSSelector: "meta[name]"}}},
		"multiple": {name: "a:href,meta:name", expected: []*Description{
			{Name: "a", Attribute: "href", CSSSelector: "a[href]"},
			{Name: "meta", Attribute: "name", CSSSelector: "meta[name]"}}},
	}
	for instance, test := range tests {
		reporter := testhelp.Errmsg(t, instance)
		actual := FromMultiSpec(test.name)
		if !reflect.DeepEqual(test.expected, actual) {
			reporter("expected %q, got %q", test.expected, actual)
		}
	}
}

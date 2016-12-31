package tag

import (
	"testing"

	"reflect"

	"github.com/elisarver/reach/testhelp"
)

// tags represent finders
func TestFromMultiSpec(t *testing.T) {
	tests := map[string]struct {
		tagname  string
		expected []*Tag
	}{
		"empty":         {tagname: "", expected: []*Tag{{Name: "", Attribute: "", CSSSelector: ""}}},
		"a automatic":   {tagname: "a", expected: []*Tag{{Name: "a", Attribute: "href", CSSSelector: "a[href]"}}},
		"img automatic": {tagname: "img", expected: []*Tag{{Name: "img", Attribute: "src", CSSSelector: "img[src]"}}},
		"full spec":     {tagname: "meta:name", expected: []*Tag{{Name: "meta", Attribute: "name", CSSSelector: "meta[name]"}}},
		"multiple": {tagname: "a:href,meta:name", expected: []*Tag{
			{Name: "a", Attribute: "href", CSSSelector: "a[href]"},
			{Name: "meta", Attribute: "name", CSSSelector: "meta[name]"}}},
	}
	for instance, test := range tests {
		reporter := testhelp.Errmsg(t, instance)
		actual := FromMultiSpec(test.tagname)
		if !reflect.DeepEqual(test.expected, actual) {
			reporter("expected %q, got %q", test.expected, actual)
		}
	}
}

package tag

import (
	"testing"

	"github.com/elisarver/reach/testhelp"
	"reflect"
)

// tags represent finders
func TestFromMultiSpec(t *testing.T) {
	tests := map[string]struct {
		tagname  string
		expected []*Tag
	}{
		"empty":     {tagname: "", expected: []*Tag{&Tag{Name: "", Attribute: "", CSSSelector: ""}}},
		"anchor":    {tagname: "a", expected: []*Tag{&Tag{Name: "a", Attribute: "href", CSSSelector: "a[href]"}}},
		"img":       {tagname: "img", expected: []*Tag{&Tag{Name: "img", Attribute: "src", CSSSelector: "img[src]"}}},
		"link":      {tagname: "link", expected: []*Tag{&Tag{Name: "link", Attribute: "href", CSSSelector: "link[href]"}}},
		"meta:name": {tagname: "meta:name", expected: []*Tag{&Tag{Name: "meta", Attribute: "name", CSSSelector: "meta[name]"}}},
		"default":   {tagname: "default", expected: []*Tag{&Tag{Name: "default", Attribute: "src", CSSSelector: "default[src]"}}},
		"multiple": {tagname: "a:href,meta:name", expected: []*Tag{
			&Tag{Name: "a", Attribute: "href", CSSSelector: "a[href]"},
			&Tag{Name: "meta", Attribute: "name", CSSSelector: "meta[name]"}}},
	}
	for instance, test := range tests {
		reporter := testhelp.Errmsg(t, instance)
		actual := FromMultiSpec(test.tagname)
		if !reflect.DeepEqual(test.expected, actual) {
			reporter("expected %q, got %q", test.expected, actual)
		}
	}
}

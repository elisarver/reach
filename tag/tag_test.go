package tag

import (
	"testing"

	"github.com/elisarver/reach/testhelp"
)

// tags represent finders
func TestFromMultiSpec(t *testing.T) {
	tests := map[string]struct {
		name     string
		expected DescriptionSlice
	}{
		"empty":         {name: "", expected: DescriptionSlice{Description{Name: "", Attribute: "", CSSSelector: ""}}},
		"a automatic":   {name: "a", expected: DescriptionSlice{Description{Name: "a", Attribute: "href", CSSSelector: "a[href]"}}},
		"img automatic": {name: "img", expected: DescriptionSlice{Description{Name: "img", Attribute: "src", CSSSelector: "img[src]"}}},
		"full spec":     {name: "meta:name", expected: DescriptionSlice{Description{Name: "meta", Attribute: "name", CSSSelector: "meta[name]"}}},
		"multiple":      {name: "a:href,meta:name", expected: DescriptionSlice{Description{Name: "a", Attribute: "href", CSSSelector: "a[href]"}, Description{Name: "meta", Attribute: "name", CSSSelector: "meta[name]"}}},
	}
	for instance, test := range tests {
		r := testhelp.Reporter(t, instance)
		actual := DescriptionSliceFromMultiSpec(test.name)
		r.Compare(test.expected, actual)
	}
}

func TestTagFinder(t *testing.T) {
	f := DescriptionFromSpec("a")

	if f.CSSSelector != f.Select() {
		t.Error("Select() should return the CSS selector")
	}
}

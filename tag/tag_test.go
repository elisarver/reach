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
		"empty":         {name: "", expected: DescriptionSlice{description{attribute: "", cssSelector: ""}}},
		"a automatic":   {name: "a", expected: DescriptionSlice{description{attribute: "href", cssSelector: "a[href]"}}},
		"img automatic": {name: "img", expected: DescriptionSlice{description{attribute: "src", cssSelector: "img[src]"}}},
		"full spec":     {name: "meta:name", expected: DescriptionSlice{description{attribute: "name", cssSelector: "meta[name]"}}},
		"multiple":      {name: "a:href,meta:name", expected: DescriptionSlice{description{attribute: "href", cssSelector: "a[href]"}, description{attribute: "name", cssSelector: "meta[name]"}}},
	}
	for instance, test := range tests {
		r := testhelp.Reporter(t, instance)
		actual := FromMultiSpec(test.name)
		r.Compare(test.expected, actual)
	}
}

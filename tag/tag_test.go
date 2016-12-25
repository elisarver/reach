package tag

import (
	"testing"

	"github.com/elisarver/reach/testhelp"
)

// tags represent finders
func TestNewTag(t *testing.T) {
	tests := map[string]struct {
		tagname  string
		expected Tag
	}{
		"empty":   {tagname: "", expected: Tag{Name: "a", Attribute: "href", CSSSelector: "a[href]"}},
		"anchor":  {tagname: "a", expected: Tag{Name: "a", Attribute: "href", CSSSelector: "a[href]"}},
		"img":     {tagname: "img", expected: Tag{Name: "img", Attribute: "src", CSSSelector: "img[src]"}},
		"link":    {tagname: "link", expected: Tag{Name: "link", Attribute: "href", CSSSelector: "link[href]"}},
		"default": {tagname: "default", expected: Tag{Name: "default", Attribute: "src", CSSSelector: "default[src]"}},
	}
	for instance, test := range tests {
		reporter := testhelp.Errmsg(t, instance)
		actual := NewTag(test.tagname)
		if test.expected != *actual {
			reporter("expected %q, got %q", test.expected, *actual)
		}
	}
}

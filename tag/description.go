package tag

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Description represents an html tag's attributes. Satisfies SelectorMapper
// +gen set
type Description struct {
	Name,
	Attribute,
	CSSSelector string
}

// Select returns a tag's CSS select string.
func (d Description) Select() string {
	return d.CSSSelector
}

// Map provides the selection function for a goquery.Map.
func (d Description) Map() func(int, *goquery.Selection) string {
	return func(_ int, sel *goquery.Selection) string {
		s, _ := sel.Attr(d.Attribute)
		return s
	}
}

// DescriptionFromSpec creates a new tag with the appropriate attributes built-in.
func DescriptionFromSpec(tagSpec string) Description {
	n, a := nameAttribute(tagSpec)
	if a == "" && n != "" {
		a = defaultAttribute(n)
	}
	return NewDescription(n, a)
}

// NewDescription creates a tag.Description with css selector
func NewDescription(name, attr string) Description {
	var s string
	if name != "" && attr != "" {
		s = fmt.Sprintf("%s[%s]", name, attr)
	}
	return Description{
		Name:        name,
		Attribute:   attr,
		CSSSelector: s,
	}
}

// nameAttribute splits a tagSpec into its name and attribute
func nameAttribute(tagSpec string) (name, attribute string) {
	s := strings.Split(tagSpec, ":")
	if len(s) > 0 && s[0] != "" {
		name = s[0]
	}
	if len(s) > 1 && s[1] != "" {
		attribute = s[1]
	}
	return name, attribute
}

// defaultAttribute provides the default link attribute for a given tag name
func defaultAttribute(name string) (attribute string) {
	switch name {
	default:
		attribute = "src"
	case "a", "link":
		attribute = "href"
	}
	return attribute
}

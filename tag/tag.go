package tag

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/elisarver/reach/collections"
	"strings"
)

// Selector provides a statement goquery can use in a Find call.
type Selector interface {
	Select() string
}

// Mapper generates an approprirate goquery map function to retrieve a tag's attribute.
type Mapper interface {
	Map() func(int, *goquery.Selection) string
}

// SelectorMapper is the intersection of something that can select results and map them over functions.
type SelectorMapper interface {
	Selector
	Mapper
}

// Description represents an html tag's attributes. Satisfies SelectorMapper
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

// SelectMap selects elements and maps them to response. Drops empty values.
func SelectMap(doc *goquery.Document, sm SelectorMapper) []string {
	return collections.DropEmpties(doc.Find(sm.Select()).Map(sm.Map()))
}

// New creates a tag.Description with css selector
func New(name, attr string) *Description {
	var s string
	if name != "" && attr != "" {
		s = fmt.Sprintf("%s[%s]", name, attr)
	}
	return &Description{
		Name:        name,
		Attribute:   attr,
		CSSSelector: s,
	}
}

// FromMultiSpec takes multiple comma-separated tag specs and turns them into a slice of tags
func FromMultiSpec(multiTagSpec string) []*Description {
	ss := strings.Split(multiTagSpec, ",")
	ds := make([]*Description, 0, len(ss))
	for _, s := range ss {
		d := FromSpec(s)
		ds = append(ds, d)
	}
	return ds
}

// FromSpec creates a new tag with the appropriate attributes built-in.
func FromSpec(tagSpec string) *Description {
	n, a := nameAttribute(tagSpec)
	if a == "" && n != "" {
		a = defaultAttribute(n)
	}
	return New(n, a)
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

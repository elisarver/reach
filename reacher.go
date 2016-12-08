package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

// Tag represents an html tag
type Tag struct {
	Name,
	Attribute,
	CSSSelector string
}

// NewTag creates a new tag with the appropriate attributes built-in.
// defaults to <a> tag.
func NewTag(name string) Tag {
	if name == "" {
		name = "a"
	}
	t := Tag{Name: name}

	switch t.Name {
	default:
		t.Attribute = "src"
	case "a", "link":
		t.Attribute = "href"
	}
	t.CSSSelector = fmt.Sprintf("%s[%s]", t.Name, t.Attribute)
	return t
}

// Selector provides a statement goquery can use in a Find call.
type Selector interface {
	Select() string
}

// Select returns a tag's CSS select string.
func (t Tag) Select() string {
	return t.CSSSelector
}

// Mapper generates an approprirate goquery map
// function to retrieve a tag's attribute.
type Mapper interface {
	Map() func(int, *goquery.Selection) string
}

// Map provides the selection function for a goquery.Map.
func (t Tag) Map() func(int, *goquery.Selection) string {
	return func(_ int, sel *goquery.Selection) string {
		s, _ := sel.Attr(t.Attribute)
		return s
	}
}

// SelectorMapper is the intersection of something that can select results and map them over functions.
type SelectorMapper interface {
	Selector
	Mapper
}

// SelectMap selects elements and maps them to response.
func SelectMap(r *goquery.Document, fm SelectorMapper) []string {
	return r.Find(fm.Select()).Map(fm.Map())
}

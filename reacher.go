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
func NewTag(name string) Tag {
	t := Tag{name, "", ""}

	switch t.Name {
	default:
		t.Attribute = "src"
	case "a", "link":
		t.Attribute = "href"
	}
	t.CSSSelector = fmt.Sprintf("%s[%s]", t.Name, t.Attribute)
	return t
}

// Finder provides a statement goquery can use in a Find call.
type Finder interface {
	Find() string
}

// Selector returns a tag's CSS selector string.
func (t Tag) Selector() string {
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

// FinderMapper is the intersection of something that can find results and map them over functions.
type FinderMapper interface {
	Finder
	Mapper
}

// FindMap finds elements and maps them to response.
func FindMap(r *goquery.Document, fm FinderMapper) []string {
	return r.Find(fm.Find()).Map(fm.Map())
}


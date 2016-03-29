package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type Tag struct {
	Name,
	Attribute,
	CSSSelector string
}

// Create a new tag with the
// appropriate attributes built-in.
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

// Finder provides a statement
// goquery can use in a Find call.
type Finder interface {
	Find() string
}

// Tags use CSS selectors for now.
func (t Tag) Find() string {
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

type FinderMapper interface {
	Finder
	Mapper
}

// Apply finder and mapper to a Response.
func FindMap(r *goquery.Document, fm FinderMapper) []string {
	return r.Find(fm.Find()).Map(fm.Map())
}

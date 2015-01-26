package main

import (
	"fmt"
	"net/url"

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

// Response is a goquery.Document
// re-typed so we can modify it.
type Response *goquery.Document

// Apply finder and mapper to a Response.
func FindMapInResponse(r Response, fm FinderMapper) []string {
	return r.Find(fm.Find()).Map(fm.Map())
}

type Target struct {
	*url.URL
}

func NewTarget(s string) (Target, error) {
	if u, err := url.ParseRequestURI(s); err == nil {
		return Target{u}, nil
	} else {
		return Target{&url.URL{}}, err
	}
}

// Reach function retrieves a goquery Document for a URL
func Reach(t Target) (Response, error) {
	return goquery.NewDocument(t.String())
}

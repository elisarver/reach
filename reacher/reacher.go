package reacher

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/elisarver/reach/tag"
)

// Selector provides a statement goquery can use in a Find call.
type Selector interface {
	Select() string
}

// Mapper generates an approprirate goquery map
// function to retrieve a tag's attribute.
type Mapper interface {
	Map() func(int, *goquery.Selection) string
}

// SelectorMapper is the intersection of something that can select results and map them over functions.
type SelectorMapper interface {
	Selector
	Mapper
}

// SelectMap selects elements and maps them to response.
func SelectMap(doc *goquery.Document, fm SelectorMapper) []string {
	return doc.Find(fm.Select()).Map(fm.Map())
}

// TagSelectorMapper applies SelectorMapper to Tag
type TagSelectorMapper struct {
	*tag.Tag
}

// Select returns a tag's CSS select string.
func (tr TagSelectorMapper) Select() string {
	return tr.CSSSelector
}

// Map provides the selection function for a goquery.Map.
func (tr TagSelectorMapper) Map() func(int, *goquery.Selection) string {
	return func(_ int, sel *goquery.Selection) string {
		s, _ := sel.Attr(tr.Attribute)
		return s
	}
}

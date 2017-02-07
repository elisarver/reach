package tag

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/elisarver/reach/collections"
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

// SelectMap selects elements and maps them to response. Drops empty values.
func SelectMap(doc *goquery.Document, sm SelectorMapper) []string {
	return collections.DropEmpties(doc.Find(sm.Select()).Map(sm.Map()))
}

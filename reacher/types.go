package reacher

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/elisarver/reach/lists"
	"github.com/elisarver/reach/target"
	"github.com/elisarver/reach/tag"
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
func SelectMap(doc *goquery.Document, fm SelectorMapper) []string {
	return lists.DropEmpties(doc.Find(fm.Select()).Map(fm.Map()))
}

// documentFn exists to make testing possible without resorting to hardcoded function.
type documentFn func(string) (*goquery.Document, error)

// TargetReacher is any function that fetches tags on targets.
type TargetReacher func([]target.Location, []*tag.Description) ([]string, error)

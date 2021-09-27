package document

import (
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/elisarver/reach/tag"
	"github.com/elisarver/reach/target"
)

// Processor is a an interface for processing a list of descriptions over a list of locations
type Processor interface {
	Process(tag.DescriptionSlice, target.LocationSlice) ([]string, error)
}

type processor struct {
	retrieve retriever
	reparent bool
}

// NewProcessor wires a retriever function and sets the reparent flag.
func NewProcessor(retriever retriever, reparent bool) Processor {
	if retriever == nil {
		retriever = genRetrieve(nil)
	}
	return processor{
		retrieve: retriever,
		reparent: reparent,
	}
}

// selectMap selects elements and maps them to response. Drops empty values.
func (p processor) selectMap(doc *goquery.Document, desc tag.Description) []string {
	return dropEmpties(doc.Find(desc.Select()).Map(p.mapGen(desc)))
}

// mapGen generates the mapping function necessary to process goquery selections
func (p processor) mapGen(desc tag.Description) func(int, *goquery.Selection) string {
	return func(_ int, sel *goquery.Selection) string {
		var s string
		if desc.Attribute() != "" {
			s, _ = sel.Attr(desc.Attribute())
		} else {
			s, _ = sel.Html()
		}
		return s
	}
}

// URLAttrs represents all attributes that have a URL-like as a value.
var URLAttrs = set{"href": nil, "link": nil, "src": nil}

// ReachTargets ranges over locations, and applies the descriptions to each document,
// in an attempt to extract values out of them. If the global Reparent config option
// is set, It also applies the URL re-parenting of relative paths to the values,
// generating more canonical site-oriented urls.
// it takes a slice of tag descriptions and a slice of target locations
func (p processor) Process(tags tag.DescriptionSlice, locations target.LocationSlice) ([]string, error) {
	reparentItem := func(s *string, fn func(string) (target.Location, error)) {
		if strings.HasPrefix(*s, "javascript") {
			return
		}
		if r, e := fn(*s); e == nil {
			*s = r.String()
		} else {
			return
		}
	}

	reparentList := func(values *[]string, fn func(string) (target.Location, error)) {
		for i := range *values {
			if strings.HasPrefix((*values)[i], "http") {
				continue
			}
			reparentItem(&(*values)[i], fn)
		}
	}

	var output []string
	for _, l := range locations {
		d, err := p.retrieve(l.String())
		if err != nil {
			return []string{}, err
		}
		for _, t := range tags {
			var values = p.selectMap(d, t)
			if p.reparent && URLAttrs.contains(t.Attribute()) {
				reparentList(&values, l.Parse)
			}

			output = append(output, values...)
		}
	}
	return output, nil
}

// dropEmpties eliminates empty values from a list of strings.
func dropEmpties(list []string) []string {
	newList := make([]string, 0, len(list))
	for i := range list {
		if list[i] != "" {
			newList = append(newList, list[i])
		}
	}
	return newList
}

// set is a map that has a membership concept.
type set map[string]interface{}

// contains checks whether a set contains a member.
func (a set) contains(attr string) bool {
	_, ok := a[attr]
	return ok
}

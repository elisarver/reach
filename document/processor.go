package document

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/elisarver/reach/tag"
	"github.com/elisarver/reach/target"
)


// Processor is a colletion of structs that are processed for values.
type Processor struct {
	Tags      tag.DescriptionSlice
	Locations target.LocationSlice
}

// NewProcessor creates a new Processor with the target locations and tag descriptions
func NewProcessor(loc target.LocationSlice, tags tag.DescriptionSlice) *Processor {
	return &Processor{
		Tags:      tags,
		Locations: loc}
}

// selectMap selects elements and maps them to response. Drops empty values.
func (p *Processor) selectMap(doc *goquery.Document, desc tag.Description) []string {
	return DropEmpties(doc.Find(desc.Select()).Map(p.mapGen(desc)))
}

//mapGen generates the mapping function necessary to process goquery selections
func (p *Processor) mapGen(desc tag.Description) func(int, *goquery.Selection) string {
	return func(_ int, sel *goquery.Selection) string {
		s, _ := sel.Attr(desc.Attribute)
		return s
	}
}

// URLAttrs represents all attributes that have a URL-like as a value.
var URLAttrs = set{"href": nil, "link": nil, "src": nil}

// ReachTargets ranges over locations, and applies the descriptions to each document,
// in an attempt to extract values out of them. If the global Reparent config option
// is set, It also applies the URL re-parenting of relative paths to the values,
// generating more canonical site-oriented urls.
func (p *Processor) ReachTargets() ([]string, error) {
	reparentItem := func(s *string, fn func(string) (target.Location, error)) {
		if strings.HasPrefix(*s, "javascript") {
			return
		}
		if r, e := fn(*s); e == nil {
			*s = r.URL.String()
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
		return
	}

	var output []string
	for _, l := range p.Locations {
		d, err := Config.Retrieve(l.String())
		if err != nil {
			return []string{}, err
		}
		for _, t := range p.Tags {
			var values = p.selectMap(d, t)
			if Config.Reparent && URLAttrs.contains(t.Attribute) {
				reparentList(&values, l.Parse)
			}

			output = append(output, values...)
		}
	}
	return output, nil
}

// DropEmpties eliminates empty values from a list of strings.
func DropEmpties(list []string) []string {
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
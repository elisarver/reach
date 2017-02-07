package reacher

import (
	"strings"

	"github.com/elisarver/reach/collections"
	"github.com/elisarver/reach/tag"
	"github.com/elisarver/reach/target"
)

type config struct {
	Reparent bool
}

var Config = config{
	// Reparent represents the re-building of a URL by the originating host's URL.
	// When set, a relative URL becomes an absolute URL with the target.Location's
	// URL parts filling in missing values.
	Reparent: false,
}

// URLAttrs represents all attributes that have a URL-like as a value.
var URLAttrs = collections.Set{"href": nil, "link": nil, "src": nil}

// ReachTargets ranges over locations, and applies the descriptions to each document,
// in an attempt to extract values out of them. If the global Reparent config option
// is set, It also applies the URL re-parenting of relative paths to the values,
// generating more canonical site-oriented urls.
func ReachTargets(ls target.LocationSlice, ds []*tag.Description) ([]string, error) {
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
	for _, l := range ls {
		resp, err := l.Retrieve()
		if err != nil {
			return []string{}, err
		}
		for _, d := range ds {
			var values = tag.SelectMap(resp, d)
			if Config.Reparent && URLAttrs.Contains(d.Attribute) {
				reparentList(&values, l.Parse)
			}
			output = append(output, values...)
		}
	}
	return output, nil
}

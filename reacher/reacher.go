package reacher

import (
	"github.com/elisarver/reach/tag"
	"github.com/elisarver/reach/target"
	"strings"
	"fmt"
	"os"
	"github.com/elisarver/reach/collections"
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

func mp(values []string, fn func(string) (target.Location, error)) []string {
	for i, v := range values {
		if !strings.HasPrefix(v, "javascript") {
			r, e := fn(v)
			if e != nil {
				fmt.Fprintf(os.Stderr, "junk value failed parse: %s", v)
				continue
			}
			nv := r.URL.String()
			if nv != v && nv != "" {
				values[i] = nv
			}
		}
	}
	return values
}

// ReachTargets ranges over locations, and applies the descriptions to each document,
// in an attempt to extract values out of them. If the global Reparent config option
// is set, It also applies the URL re-parenting of relative paths to the values,
// generating more canonical site-oriented urls.
func ReachTargets(ls []target.Location, ds []*tag.Description) ([]string, error) {
	var output []string
	for _, l := range ls {
		resp, err := l.Retrieve()
		if err != nil {
			return []string{}, err
		}
		for _, d := range ds {
			var values = tag.SelectMap(resp, d)
			if Config.Reparent && URLAttrs.Contains(d.Attribute) {
				values = mp(values, l.Parse)
			}
			output = append(output, values...)
		}
	}
	return output, nil
}

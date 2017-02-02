package reacher

import (
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

func ReachTargets(ls []target.Location, ds []*tag.Description) ([]string, error) {
	var output []string
	for _, l := range ls {
		resp, err := l.Retrieve()
		if err != nil {
			return []string{}, err
		}
		for _, d := range ds {
			output = append(output, tag.SelectMap(resp, d)...)
		}
	}
	return output, nil
}

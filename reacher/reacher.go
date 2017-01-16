package reacher

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/elisarver/reach/tag"
	"github.com/elisarver/reach/target"
)

// genReachTargets binds the appropriate function to generate a document
// to the ReachTargets func.
func genReachTargets(fn documentFn) TargetReacher {
	if fn == nil {
		fn = goquery.NewDocument
	}
	return func(ls []target.Location, ds []*tag.Description) ([]string, error) {
		var output []string
		for _, l := range ls {
			resp, err := fn(l.String())
			if err != nil {
				return []string{}, err
			}
			for _, d := range ds {
				output = append(output, SelectMap(resp, d)...)
			}
		}
		return output, nil
	}
}

var (
	// ReachTargets takes a list of targets, a list of tags, and a fetcher, fetches the targets with the
	// fetcher, and finds the tags in the document.
	ReachTargets = genReachTargets(nil)
)

package reacher

import (
	"github.com/elisarver/reach/tag"
	"github.com/elisarver/reach/target"
)

// genReachTargets binds the appropriate function to generate a document
// to the ReachTargets func.
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

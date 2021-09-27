package tag

import "strings"

// FromMultiSpec takes multiple comma-separated tag specs and turns them into a slice of tags.
func FromMultiSpec(multiTagSpec string) DescriptionSlice {
	ss := strings.Split(multiTagSpec, ",")
	ds := make(DescriptionSlice, 0, len(ss))
	for _, s := range ss {
		d := FromSpec(s)
		ds = append(ds, d)
	}

	return ds
}

func RawQuery(query string) DescriptionSlice {
	return DescriptionSlice{description{cssSelector: query}}
}

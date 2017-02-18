package tag

import "strings"

// DescriptionSliceFromMultiSpec takes multiple comma-separated tag specs and turns them into a slice of tags
func DescriptionSliceFromMultiSpec(multiTagSpec string) DescriptionSlice {
	ss := strings.Split(multiTagSpec, ",")
	ds := make(DescriptionSlice, 0, len(ss))
	for _, s := range ss {
		d := DescriptionFromSpec(s)
		ds = append(ds, d)
	}
	return ds
}

func RawQuery(query string) DescriptionSlice {
	return DescriptionSlice{Description{CSSSelector: query}}
}

package tag

import "strings"

// DescriptionSetFromMultiSpec takes multiple comma-separated tag specs and turns them into a slice of tags
func DescriptionSetFromMultiSpec(multiTagSpec string) DescriptionSet {
	ss := strings.Split(multiTagSpec, ",")
	ds := make(DescriptionSet, len(ss))
	for _, s := range ss {
		d := DescriptionFromSpec(s)
		ds.Add(d)
	}
	return ds
}

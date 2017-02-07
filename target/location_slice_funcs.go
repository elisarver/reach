package target

func (ls LocationSlice) DistinctByURL() LocationSlice {
	urlPredicate := func(a, b Location) bool {
		return a.URL.String() == b.URL.String()
	}
	return ls.DistinctBy(urlPredicate)
}

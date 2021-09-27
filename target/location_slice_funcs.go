package target

// DistinctByURL retrieves a url by its distinct URL string value.
func (ls LocationSlice) DistinctByURL() LocationSlice {
	urlPredicate := func(a, b Location) bool {
		return a.String() == b.String()
	}

	return ls.DistinctBy(urlPredicate)
}

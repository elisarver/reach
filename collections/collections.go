package collections

// DropEmpties eliminates empty values from a list of strings.
func DropEmpties(list []string) []string {
	newList := make([]string, 0, len(list))
	for i := range list {
		if list[i] != "" {
			newList = append(newList, list[i])
		}
	}
	return newList
}

// Set is a map that has a membership concept.
type Set map[string]interface{}

func (a Set) Contains(attr string) bool {
	_, ok := a[attr]
	return ok
}

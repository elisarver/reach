package tag

import "fmt"

// Tag represents an html tag
type Tag struct {
	Name,
	Attribute,
	CSSSelector string
}

// NewTag creates a new tag with the appropriate attributes built-in.
// defaults to <a> tag.
func NewTag(name string) Tag {
	if name == "" {
		name = "a"
	}
	t := Tag{Name: name}
	switch t.Name {
	default:
		t.Attribute = "src"
	case "a", "link":
		t.Attribute = "href"
	}
	t.CSSSelector = fmt.Sprintf("%s[%s]", t.Name, t.Attribute)
	return t
}

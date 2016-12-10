package tag

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

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

// Select returns a tag's CSS select string.
func (t Tag) Select() string {
	return t.CSSSelector
}

// Map provides the selection function for a goquery.Map.
func (t Tag) Map() func(int, *goquery.Selection) string {
	return func(_ int, sel *goquery.Selection) string {
		s, _ := sel.Attr(t.Attribute)
		return s
	}
}

package tag

import "fmt"
import "strings"

// Tag represents an html tag
type Tag struct {
	Name,
	Attribute,
	CSSSelector string
}

// NewTag creates a new tag with the appropriate attributes built-in.
//
// defaults to 'a' tag Name. defaults to 'src' Attribute, unless a or link, in which case it is 'href'
//
// tagSpec is a string of the form [tag]:[attr]
//
// Call NewTag if you have no tag, or have no Attribute to give.
func NewTag(tagSpec string) *Tag {
	spec := strings.Split(tagSpec, ":")
	var tag, attr string
	if len(spec) > 0 && spec[0] != "" {
		tag = spec[0]
	}
	if len(spec) > 1 && spec[1] != "" {
		attr = spec[1]
	}
	if tag == "" {
		tag = "a"
	}
	if attr == "" {
		switch tag {
		default:
			attr = "src"
		case "a", "link":
			attr = "href"
		}
	}
	t := &Tag{
		Name:        tag,
		Attribute:   attr,
		CSSSelector: fmt.Sprintf("%s[%s]", tag, attr)}
	return t
}

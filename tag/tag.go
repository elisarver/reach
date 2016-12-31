package tag

import "fmt"
import "strings"

// Tag represents an html tag
type Tag struct {
	Name,
	Attribute,
	CSSSelector string
}

// New creates a tag with css selector
func New(name, attr string) *Tag {
	var selector string
	if name != "" && attr != "" {
		selector = fmt.Sprintf("%s[%s]", name, attr)
	}
	return &Tag{
		Name:        name,
		Attribute:   attr,
		CSSSelector: selector,
	}
}

// FromMultiSpec takes multiple comma-separated tag specs and turns them into a slice of tags
func FromMultiSpec(multiTagSpec string) []*Tag {
	specs := strings.Split(multiTagSpec, ",")
	tags := make([]*Tag, 0, len(specs))
	for _, spec := range specs {
		tag := FromSpec(spec)
		tags = append(tags, tag)
	}
	return tags
}

// FromSpec creates a new tag with the appropriate attributes built-in.
func FromSpec(tagSpec string) *Tag {
	name, attr := NameAttribute(tagSpec)
	if attr == "" && name != "" {
		attr = DefaultAttribute(name)
	}
	return New(name, attr)
}

// NameAttribute splits a tagSpec into its name and attribute
func NameAttribute(tagSpec string) (name, attribute string) {
	spec := strings.Split(tagSpec, ":")
	if len(spec) > 0 && spec[0] != "" {
		name = spec[0]
	}
	if len(spec) > 1 && spec[1] != "" {
		attribute = spec[1]
	}
	return name, attribute
}

// DefaultAttribute provides the default link attribute for a given tag name
func DefaultAttribute(name string) (attribute string) {
	switch name {
	default:
		attribute = "src"
	case "a", "link":
		attribute = "href"
	}
	return attribute
}

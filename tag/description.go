package tag

import (
	"fmt"
	"strings"
)

// description represents an html tag's attributes. Satisfies SelectorMapper
// +gen slice
type Description interface {
	Attribute() string
	Select() string
}

type description struct {
	attribute   string
	cssSelector string
}

// Select returns a tag's CSS select string.
func (d description) Select() string {
	return d.cssSelector
}

func (d description) Attribute() string {
	return d.attribute
}

// FromSpec creates a new tag with the appropriate attributes built-in.
func FromSpec(tagSpec string) Description {
	n, a := nameAttribute(tagSpec)
	if a == "" && n != "" {
		a = defaultAttribute(n)
	}
	return NewDescription(n, a)
}

// NewDescription creates a tag.description with css selector
func NewDescription(name, attr string) Description {
	var selector string
	if name != "" && attr != "" {
		selector = fmt.Sprintf("%s[%s]", name, attr)
	}
	return description{
		cssSelector: selector,
		attribute:   attr,
	}
}

// nameAttribute splits a tagSpec into its name and attribute
func nameAttribute(tagSpec string) (name, attribute string) {
L:
	for i, v := range strings.Split(tagSpec, ":") {
		switch i {
		case 0:
			name = v
		case 1:
			attribute = v
		default:
			// we don't parse beyond two values
			break L
		}
	}
	return name, attribute
}

// defaultAttribute provides the default reference attribute for a given tag name
func defaultAttribute(tag string) (attribute string) {
	switch tag {
	case "a", "link":
		return "href"
	default:
		return "src"
	}
}

package main

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

// tags represent finders
func TestNewTag(t *testing.T) {

	names := [4]string{
		"a", "img", "link", "dontcare",
	}

	expecteds := [4]Tag{
		{"a", "href", "a[href]"},
		{"img", "src", "img[src]"},
		{"link", "href", "link[href]"},
		{"dontcare", "src", "dontcare[src]"},
	}

	for i := range names {
		name := names[i]
		act := NewTag(name)
		exp := expecteds[i]

		if act != exp {
			t.Errorf("expected %q, got %q", exp, act)
		}
	}
}

func TestTagFinder(t *testing.T) {
	doc, _ := goquery.NewDocumentFromReader(
		strings.NewReader("<a href='http://www.example.com/'/>"))
	var f Finder = NewTag("a")

	if f.(Tag).CSSSelector != f.Find() {
		t.Error("Find() should return the CSS selector")
	}

	act := doc.Find(f.Find())
	if act.Size() != 1 {
		t.Error("expected only one a[href] match.")
	}
}

func TestTagMapper(t *testing.T) {
	doc, _ := goquery.NewDocumentFromReader(
		strings.NewReader("<a href='http://www.example.com/'/><link href=''/><dontcare/>"))

	var m Mapper = NewTag("a")

	act := doc.Find("a").Map(m.Map())
	if act[0] != "http://www.example.com/" {
		t.Error("Map should have resulted in extracting the url.")
	}

	// negative
	act = doc.Find("img").Map(m.Map())
	if len(act) != 0 {
		t.Error("Map should have 0 entries.")
	}

	m = NewTag("dontcare")
	act = doc.Find("dontcare").Map(m.Map())
	if len(act) != 1 {
		t.Error("Map should have 1 entry")
	}
	if act[0] != "" {
		t.Error("first value should be empty")
	}
}

func TestFindMap(t *testing.T) {
	var res *goquery.Document
	res, _ = goquery.NewDocumentFromReader(
		strings.NewReader("<a href='http://www.example.com/'/><link href=''/><dontcare/>"))

	var fm FinderMapper = NewTag("a")
	exp := []string{"http://www.example.com/"}
	act := FindMap(res, fm)

	if len(act) != 1 {
		t.Error("Map should have 1 entry")
	}

	if act[0] != exp[0] {
		t.Errorf("expected %q, got %q", exp[0], act[0])
	}
}

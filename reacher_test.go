package main

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

// tags represent finders
func TestNewTag(t *testing.T) {
	for _, pair := range []struct {
		tagname  string
		expected Tag
	}{
		{"", Tag{"a", "href", "a[href]"}},
		{"a", Tag{"a", "href", "a[href]"}},
		{"img", Tag{"img", "src", "img[src]"}},
		{"link", Tag{"link", "href", "link[href]"}},
		{"dontcare", Tag{"dontcare", "src", "dontcare[src]"}},
	} {
		actual := NewTag(pair.tagname)
		if pair.expected != actual {
			t.Errorf("expected %q, got %q", pair.expected, actual)
		}
	}
}

func TestTagFinder(t *testing.T) {
	doc := genDoc(t, "<a href='http://www.example.com/'/>")
	var f Selector = NewTag("a")

	if f.(Tag).CSSSelector != f.Select() {
		t.Error("Find() should return the CSS selector")
	}

	act := doc.Find(f.Select())
	if act.Size() != 1 {
		t.Error("expected only one a[href] match.")
	}
}

func TestTagMapper(t *testing.T) {
	doc := genDoc(t, "<a href='http://www.example.com/'/><link href=''/><dontcare/>")
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

func TestSelectMap(t *testing.T) {
	doc := genDoc(t, "<a href='http://www.example.com/'/><link href=''/><dontcare/>")
	var fm SelectorMapper = NewTag("a")
	exp := []string{"http://www.example.com/"}
	act := SelectMap(doc, fm)

	if len(act) != 1 {
		t.Error("Map should have 1 entry")
	}

	if act[0] != exp[0] {
		t.Errorf("expected %q, got %q", exp[0], act[0])
	}
}

func genDoc(t *testing.T, s string) *goquery.Document {
	var (
		res *goquery.Document
		err error
	)
	if res, err = goquery.NewDocumentFromReader(strings.NewReader(s)); err == nil {
		t.Error(err)
	}
	return res
}

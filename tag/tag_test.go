package tag

import (
	"testing"

	"reflect"

	"github.com/elisarver/reach/testhelp"
)

// tags represent finders
func TestFromMultiSpec(t *testing.T) {
	tests := map[string]struct {
		name     string
		expected []*Description
	}{
		"empty":         {name: "", expected: []*Description{{Name: "", Attribute: "", CSSSelector: ""}}},
		"a automatic":   {name: "a", expected: []*Description{{Name: "a", Attribute: "href", CSSSelector: "a[href]"}}},
		"img automatic": {name: "img", expected: []*Description{{Name: "img", Attribute: "src", CSSSelector: "img[src]"}}},
		"full spec":     {name: "meta:name", expected: []*Description{{Name: "meta", Attribute: "name", CSSSelector: "meta[name]"}}},
		"multiple": {name: "a:href,meta:name", expected: []*Description{
			{Name: "a", Attribute: "href", CSSSelector: "a[href]"},
			{Name: "meta", Attribute: "name", CSSSelector: "meta[name]"}}},
	}
	for instance, test := range tests {
		reporter := testhelp.Errmsg(t, instance)
		actual := FromMultiSpec(test.name)
		if !reflect.DeepEqual(test.expected, actual) {
			reporter("expected %q, got %q", test.expected, actual)
		}
	}
}

func TestTagFinder(t *testing.T) {
	doc := testhelp.GenDoc(t, "<a href='http://www.example.com/'/>")
	f := FromSpec("a")

	if f.CSSSelector != f.Select() {
		t.Error("Select() should return the CSS selector")
	}

	act := doc.Find(f.Select())
	if act.Size() != 1 {
		t.Error("expected only one a[href] match.")
	}
}

func TestTagMapper(t *testing.T) {
	doc := testhelp.GenDoc(t, "<a href='http://www.example.com/'/><link href=''/><dontcare/>")
	tsm := FromSpec("a")

	act := doc.Find("a").Map(tsm.Map())
	if act[0] != "http://www.example.com/" {
		t.Error("Map should have resulted in extracting the url.")
	}

	// negative
	act = doc.Find("img").Map(tsm.Map())
	if len(act) != 0 {
		t.Error("Map should have 0 entries.")
	}

	tsm = FromSpec("dontcare")
	act = doc.Find("dontcare").Map(tsm.Map())
	if len(act) != 1 {
		t.Error("Map should have 1 entry")
	}
	if act[0] != "" {
		t.Error("first value should be empty")
	}
}

func TestSelectMap(t *testing.T) {
	doc := testhelp.GenDoc(t, "<a href='http://www.example.com/'/><a href=''/><link href=''/><dontcare/>")
	fm := FromSpec("a")
	exp := []string{"http://www.example.com/"}
	act := SelectMap(doc, fm)

	if len(act) != 1 {
		t.Error("Map should have 1 entry")
	}

	if act[0] != exp[0] {
		t.Errorf("expected %q, got %q", exp[0], act[0])
	}
}

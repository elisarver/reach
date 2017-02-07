package tag

import (
	"testing"

	"github.com/elisarver/reach/testhelp"
)

// tags represent finders
func TestFromMultiSpec(t *testing.T) {
	tests := map[string]struct {
		name     string
		expected DescriptionSet
	}{
		"empty":         {name: "", expected: NewDescriptionSet(Description{Name: "", Attribute: "", CSSSelector: ""})},
		"a automatic":   {name: "a", expected: NewDescriptionSet(Description{Name: "a", Attribute: "href", CSSSelector: "a[href]"})},
		"img automatic": {name: "img", expected: NewDescriptionSet(Description{Name: "img", Attribute: "src", CSSSelector: "img[src]"})},
		"full spec":     {name: "meta:name", expected: NewDescriptionSet(Description{Name: "meta", Attribute: "name", CSSSelector: "meta[name]"})},
		"multiple":      {name: "a:href,meta:name", expected: NewDescriptionSet(Description{Name: "a", Attribute: "href", CSSSelector: "a[href]"}, Description{Name: "meta", Attribute: "name", CSSSelector: "meta[name]"})},
	}
	for instance, test := range tests {
		r := testhelp.Reporter(t, instance)
		actual := DescriptionSetFromMultiSpec(test.name)
		r.Compare(test.expected, actual)
	}
}

func TestTagFinder(t *testing.T) {
	doc := testhelp.GenDoc(t, "<a href='http://www.example.com/'/>")
	f := DescriptionFromSpec("a")

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
	tsm := DescriptionFromSpec("a")

	act := doc.Find("a").Map(tsm.Map())
	if act[0] != "http://www.example.com/" {
		t.Error("Map should have resulted in extracting the url.")
	}

	// negative
	act = doc.Find("img").Map(tsm.Map())
	if len(act) != 0 {
		t.Error("Map should have 0 entries.")
	}

	tsm = DescriptionFromSpec("dontcare")
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
	fm := DescriptionFromSpec("a")
	exp := []string{"http://www.example.com/"}
	act := SelectMap(doc, fm)

	if len(act) != 1 {
		t.Error("Map should have 1 entry")
	}

	if act[0] != exp[0] {
		t.Errorf("expected %q, got %q", exp[0], act[0])
	}
}

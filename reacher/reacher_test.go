package reacher

import (
	"reflect"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/elisarver/reach/tag"
	"github.com/elisarver/reach/target"
)

func TestTagFinder(t *testing.T) {
	doc := genDoc(t, "<a href='http://www.example.com/'/>")
	f := TagSelectorMapper{tag.FromSpec("a")}

	if f.CSSSelector != f.Select() {
		t.Error("Select() should return the CSS selector")
	}

	act := doc.Find(f.Select())
	if act.Size() != 1 {
		t.Error("expected only one a[href] match.")
	}
}

func TestTagMapper(t *testing.T) {
	doc := genDoc(t, "<a href='http://www.example.com/'/><link href=''/><dontcare/>")
	m := TagSelectorMapper{tag.FromSpec("a")}

	act := doc.Find("a").Map(m.Map())
	if act[0] != "http://www.example.com/" {
		t.Error("Map should have resulted in extracting the url.")
	}

	// negative
	act = doc.Find("img").Map(m.Map())
	if len(act) != 0 {
		t.Error("Map should have 0 entries.")
	}

	m = TagSelectorMapper{tag.FromSpec("dontcare")}
	act = doc.Find("dontcare").Map(m.Map())
	if len(act) != 1 {
		t.Error("Map should have 1 entry")
	}
	if act[0] != "" {
		t.Error("first value should be empty")
	}
}

func TestSelectMap(t *testing.T) {
	doc := genDoc(t, "<a href='http://www.example.com/'/><a href=''/><link href=''/><dontcare/>")
	fm := TagSelectorMapper{tag.FromSpec("a")}
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
	if res, err = goquery.NewDocumentFromReader(strings.NewReader(s)); err != nil {
		t.Error(err)
	}
	return res
}

func TestReachTargets(t *testing.T) {
	reachFnSuccess := func(_ string) (*goquery.Document, error) {
		r := strings.NewReader("<html><body><a href='http://foo.bar/'>site</a><img src='/logo.png'/></body></html>")
		return goquery.NewDocumentFromReader(r)
	}
	// use the internal generator to separate defaults from
	// runtime version
	rf := genReachTargets(reachFnSuccess)
	u, _ := target.NewTarget("http://foo.bar/")
	us := []target.Target{u}
	tags := []*tag.Tag{tag.FromSpec("a"), tag.FromSpec("img")}
	actual, err := rf(us, tags)
	if err != nil {
		t.Errorf("test didn't expect %s", err)
	}
	expected := []string{"http://foo.bar/", "/logo.png"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

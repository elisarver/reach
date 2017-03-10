package document

import (
	"reflect"
	"strings"
	"testing"

	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/elisarver/reach/tag"
	"github.com/elisarver/reach/target"
	"github.com/elisarver/reach/testhelp"
)

func TestReachTargets(t *testing.T) {
	Config.Reparent = true
	defer func() { Config.Reparent = false }()

	reachFnSuccess := func(_ string) (*goquery.Document, error) {
		r := strings.NewReader("<html><body><a href='http://foo.bar/'>site</a><img src='/logo.png'/></body></html>" +
			"<img src='javascript:'/>")
		return goquery.NewDocumentFromReader(r)
	}
	// use the internal generator to separate defaults from
	// runtime version
	Config.Retrieve = genRetrieve(reachFnSuccess)
	l, _ := target.Parse("http://foo.bar/")
	ls := target.LocationSlice{l}
	ds := tag.DescriptionSlice{tag.FromSpec("a"), tag.FromSpec("img")}
	processor := NewProcessor(ls, ds)
	actual, err := processor.ReachTargets()
	if err != nil {
		t.Errorf("test didn't expect %s", err)
	}
	expected := []string{"http://foo.bar/", "http://foo.bar/logo.png", "javascript:"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestRawQuery(t *testing.T) {
	reachFnSuccess := func(_ string) (*goquery.Document, error) {
		r := strings.NewReader("<html><body><p>body1</p><p>body2</p></body></html>")
		return goquery.NewDocumentFromReader(r)
	}
	// use the internal generator to separate defaults from
	// runtime version
	Config.Retrieve = genRetrieve(reachFnSuccess)
	l, _ := target.Parse("http://foo.bar/")
	ls := target.LocationSlice{l}
	ds := tag.RawQuery("body")
	processor := NewProcessor(ls, ds)
	actual, err := processor.ReachTargets()
	if err != nil {
		t.Errorf("test didn't expect %s", err)
	}
	expected := []string{"<p>body1</p><p>body2</p>"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestDropEmpties(t *testing.T) {
	input := []string{"not empty", "", "also not empty"}
	expected := fmt.Sprintf("%q", []string{"not empty", "also not empty"})
	actual := fmt.Sprintf("%q", DropEmpties(input))
	r := testhelp.Reporter(t, "dropEmpties")
	r.Compare(expected, actual)
}

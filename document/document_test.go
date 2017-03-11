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
	reachFnSuccess := func(_ string) (*goquery.Document, error) {
		r := strings.NewReader("<html><body><a href='http://foo.bar/'>site</a><img src='/logo.png'/></body></html>" +
			"<img src='javascript:'/>")
		return goquery.NewDocumentFromReader(r)
	}
	processor := NewProcessor(genRetrieve(reachFnSuccess), true)
	l, _ := target.Parse("http://foo.bar/")
	ls := target.LocationSlice{l}
	ds := tag.DescriptionSlice{tag.FromSpec("a"), tag.FromSpec("img")}
	actual, err := processor.Process(ds, ls)
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

	processor := NewProcessor(genRetrieve(reachFnSuccess), false)
	l, _ := target.Parse("http://foo.bar/")
	ls := target.LocationSlice{l}
	ds := tag.RawQuery("body")
	actual, err := processor.Process(ds, ls)
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
	actual := fmt.Sprintf("%q", dropEmpties(input))
	r := testhelp.Reporter(t, "dropEmpties")
	r.Compare(expected, actual)
}

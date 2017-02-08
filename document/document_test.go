package document

import (
	"reflect"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/elisarver/reach/tag"
	"github.com/elisarver/reach/target"
)

func TestReachTargets(t *testing.T) {
	Config.Reparent = true
	reachFnSuccess := func(_ string) (*goquery.Document, error) {
		r := strings.NewReader("<html><body><a href='http://foo.bar/'>site</a><img src='/logo.png'/></body></html>" +
			"<img src='javascript:'/>")
		return goquery.NewDocumentFromReader(r)
	}
	// use the internal generator to separate defaults from
	// runtime version
	Config.Retrieve = genRetrieve(reachFnSuccess)
	l, _ := target.NewLocation("http://foo.bar/")
	ls := target.LocationSlice{l}
	ds := tag.DescriptionSlice{tag.DescriptionFromSpec("a"), tag.DescriptionFromSpec("img")}
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

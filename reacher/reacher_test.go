package reacher

import (
	"reflect"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/elisarver/reach/tag"
	"github.com/elisarver/reach/target"
)

func TestReachTargets(t *testing.T) {
	reachFnSuccess := func(_ string) (*goquery.Document, error) {
		r := strings.NewReader("<html><body><a href='http://foo.bar/'>site</a><img src='/logo.png'/></body></html>")
		return goquery.NewDocumentFromReader(r)
	}
	// use the internal generator to separate defaults from
	// runtime version
	target.Config.Retrieve = target.GenRetrieve(reachFnSuccess)
	l, _ := target.New("http://foo.bar/")
	ls := []target.Location{l}
	ds := []*tag.Description{tag.FromSpec("a"), tag.FromSpec("img")}
	actual, err := ReachTargets(ls, ds)
	if err != nil {
		t.Errorf("test didn't expect %s", err)
	}
	expected := []string{"http://foo.bar/", "/logo.png"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

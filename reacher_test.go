package reach

import (
	"net/url"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

type beforeAfter struct {
	subject,
	expected Reacher
	results []string
}

var (
	bu, _   = url.Parse("http://example.com/")
	testSet = [...]beforeAfter{
		beforeAfter{
			Reacher{Element: "a", Local: false, BaseURL: bu},
			Reacher{Element: "a", Attribute: "href", Selector: "a[href]"},
			[]string{"http://example.com/a"},
		},
		beforeAfter{
			Reacher{Element: "img", Local: false, BaseURL: bu},
			Reacher{Element: "img", Attribute: "src", Selector: "img[src]"},
			[]string{"http://example.com/img"},
		},
		beforeAfter{
			Reacher{Element: "foo", Local: false, BaseURL: bu},
			Reacher{Element: "foo", Attribute: "src", Selector: "foo[src]"},
			[]string{"http://somewhereelse.com/bar"},
		},
		beforeAfter{
			Reacher{Element: "foo", Local: true, BaseURL: bu},
			Reacher{Element: "foo", Attribute: "src", Selector: "foo[src]"},
			[]string{""},
		},
	}

	testHtml = "<a href='/a'/><img src='/img'/><foo src='http://somewhereelse.com/bar'/>"
)

func TestGenAttribute(t *testing.T) {
	for _, tt := range testSet {
		tt.subject.genAttribute()
		if tt.subject.Attribute != tt.expected.Attribute {
			t.Errorf("Expected Reacher.Attribute to have value %q, but it had %q",
				tt.expected.Attribute, tt.subject.Attribute)
		}
	}
}

func TestGenSelector(t *testing.T) {
	for _, tt := range testSet {
		tt.subject.genSelector()
		if tt.subject.Selector != tt.expected.Selector {
			t.Errorf("Expected Reacher.Selector to have value %q, but it had %q",
				tt.expected.Selector, tt.subject.Selector)
		}
	}
}

func TestGenMapper(t *testing.T) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(testHtml))
	if err != nil {
		t.Errorf("Error parsing testHtml %q; error:%q", testHtml, err)
	}
	for _, tt := range testSet {
		tt.subject.genAll()

		if tt.subject.mapper == nil {
			t.Error("Expected Reacher.Mapper to contain an anonymous function.")
		} else {
			results := tt.subject.fromDocument(doc)
			if len(results) != len(tt.results) {
				t.Errorf("Got different results in expected than actual;\nactual: %q\nexpected%q\n", results, tt.results)
			}
			for i := range tt.results {
				if results[i] != tt.results[i] {
					t.Errorf("Result %d for selector %s: Expected %q, got %q", i, tt.expected.Selector, tt.results[0], results[0])
				}
			}
		}
	}
}

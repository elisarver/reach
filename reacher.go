package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type Reacher struct {
	Tag,
	Selector,
	Attribute string
	BaseURL *url.URL
	Local   bool
	mapper  func(int, *goquery.Selection) string
}

func (r *Reacher) fromDocument(doc *goquery.Document) []string {
	return doc.Find(r.Selector).Map(r.mapper)
}

func (r *Reacher) Reach() ([]string, error) {
	if r.BaseURL == nil {
		return nil, fmt.Errorf("BaseURL was not set")
	}
	r.genAll()

	resp, err := http.Get(r.BaseURL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}

	return r.fromDocument(doc), nil
}

func (r *Reacher) genAll() {
	r.genAttribute()
	r.genSelector()
	r.genMapper()
}

func (r *Reacher) genAttribute() {
	switch r.Tag {
	case "a", "link":
		r.Attribute = "href"
	default:
		r.Attribute = "src"
	}
}

func (r *Reacher) genSelector() {
	if r.Attribute == "" {
		r.genAttribute()
	}
	r.Selector = fmt.Sprintf("%s[%s]", r.Tag, r.Attribute)
}

func (r *Reacher) genMapper() {
	r.mapper = func(_ int, sel *goquery.Selection) string {
		s, ok := sel.Attr(r.Attribute)
		if !ok {
			return ""
		}

		u, err := url.Parse(s)
		if err != nil {
			return ""
		}

		u = r.BaseURL.ResolveReference(u)

		if r.Local && u.Host != r.BaseURL.Host {
			return ""
		}

		return u.String()
	}
}

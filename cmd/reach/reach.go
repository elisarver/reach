package main

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/elisarver/reach/reacher"
	"github.com/elisarver/reach/tag"
	"github.com/elisarver/reach/target"
)

var (
	pTag string
)

const (
	examples = `
Examples:

  Get all img src from a web page:

  > reach -t img http://blog.golang.org
  http://blog.golang.org/gophergala/fancygopher.jpg

  Get all unique links on a page:

  > reach http://example.com/ | sort | uniq
  http://example.com/blog
  http://example.com/about

`
)

func init() {
	const (
		defaultTag = "a"
		tagUsage   = "Tag to search for."
	)

	flag.Usage = func() {
		cmd := filepath.Base(os.Args[0])
		fmt.Fprint(os.Stderr, "Reach gathers urls from a website.\n\n")
		fmt.Fprintf(os.Stderr,
			"Usage:\n\n  %s [-t=\"a\" | -tag=\"img\"] URLs...\n", cmd)
		fmt.Fprint(os.Stderr, examples)
	}

	flag.StringVar(&pTag, "tag", defaultTag, tagUsage)
	flag.StringVar(&pTag, "t", defaultTag, tagUsage+" (Shorthand)")
}

func main() {
	flag.Parse()

	ts, err := argTargets(flag.Args())
	trap(err)

	output := reachTargets(ts, pTag, Reach)
	fmt.Print(strings.Join(output, "\n"))
	fmt.Println()
}

type rf func(*target.Target) (*goquery.Document, error)

func reachTargets(ts []*target.Target, tagName string, reachFn rf) []string {
	var (
		output = make([]string, len(ts))
		tag    = tag.NewTag(tagName)
	)
	for i, t := range ts {
		resp, err := reachFn(t)
		trap(err)

		URLs := reacher.SelectMap(resp, reacher.TagSelectorMapper(tag))

		output[i] = strings.Join(dropEmpties(URLs), "\n")
	}
	return output
}

// Reach function retrieves a goquery Document for a URL
func Reach(t *target.Target) (*goquery.Document, error) {
	return goquery.NewDocument(t.String())
}

// argTargets filters the incoming argument array,
// verifying that the parameters are able to provide
// request URIs. Returns a list of targets or an error.
// Always check error! In order to return a good default
// type, we pass back a slice of Target. This is not a
// pointer, so a nil check is unnecessary. Target aliases
// to strings verified by this function.
func argTargets(args []string) ([]*target.Target, error) {
	numArgs := len(args)
	if numArgs == 0 {
		return []*target.Target{}, errors.New("please supply at least one URL")
	}
	t := target.Target{URL: &url.URL{}}
	return t.ParseAll(args)
}

// dropEmpties eliminates empty values from a list of strings.
func dropEmpties(list []string) []string {
	newList := make([]string, 0, len(list))
	for i := range list {
		if list[i] != "" {
			newList = append(newList, list[i])
		}
	}
	return newList
}

func trap(err error) {
	if err != nil {
		flag.Usage()
		fmt.Fprintln(os.Stderr, fmt.Errorf("%s", err.Error()))
		os.Exit(1)
	}
}

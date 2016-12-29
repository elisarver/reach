package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/elisarver/reach/reacher"
	"github.com/elisarver/reach/tag"
	"github.com/elisarver/reach/target"
)

var (
	ErrOneURL    = errors.New("please supply at least one URL")
	ErrEmptyTags = errors.New("Empty list of tags")
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

  Get mutiple types:

  > reach -t img,a http://example.com/
  /example.png
  /about.html

  Or attributes:

  > reach -t img:class http://example.com/
  logo
`
)

func main() {
	flag.Usage = func() {
		cmd := filepath.Base(os.Args[0])
		fmt.Fprint(os.Stderr, "Reach gathers urls from a website.\n\n")
		fmt.Fprintf(os.Stderr,
			"Usage:\n\n  %s [-t=\"a\" | -tag=\"img\"] URLs...\n", cmd)
		fmt.Fprint(os.Stderr, examples)
	}

	var pTag string
	tagUsage := "Tag to search for."
	flag.StringVar(&pTag, "tag", "a", tagUsage)
	flag.StringVar(&pTag, "t", "a", tagUsage+" (Shorthand)")
	flag.Parse()

	if pTag == "" {
		pTag = "a"
	}

	tagsList := strings.Split(pTag, ",")
	tags := make([]*tag.Tag, 0, len(tagsList))

	for _, v := range tagsList {
		tags = append(tags, tag.NewTag(v))
	}

	if len(flag.Args()) == 0 {
		fmt.Fprintln(os.Stderr, ErrOneURL.Error())
	}

	targets, err := target.ParseAll(flag.Args())
	if err != nil {
		fmt.Fprintln(os.Stderr, ErrOneURL.Error())
		os.Exit(1)
	}

	output, err := reacher.ReachTargets(targets, tags, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("unexpected error %s", err.Error()))
		os.Exit(1)
	}
	fmt.Print(strings.Join(output, "\n"))
	fmt.Println()
}

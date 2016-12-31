package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/elisarver/reach/reacher"
	"github.com/elisarver/reach/tag"
	"github.com/elisarver/reach/target"
)

var (
	ErrOneURL    = errors.New("please supply at least one URL")
	ErrEmptyTags = errors.New("empty tag arguments")
)

func main() {
	var pTag string
	flag.StringVar(&pTag, "tag", "a", "comma-separated list of `tags` to search for.\n\tformat: name1[:attribute1][,name2[:attribute2]]")
	var help bool
	flag.BoolVar(&help, "h", false, "print help")
	flag.Parse()

	if help {
		fmt.Fprint(os.Stderr, 
`Description:
	reach visits URLs and returns tags/attributes of interest.
`)
		flag.Usage()
		fmt.Fprint(os.Stderr, 
`Positional arguments:
  URL [URL...]
  	one or more URLs to dial.

`)
		os.Exit(0)
	}

	tagsList := strings.Split(pTag, ",")
	tags := make([]*tag.Tag, 0, len(tagsList))

	for _, v := range tagsList {
		tags = append(tags, tag.NewTag(v))
	}

	if len(flag.Args()) == 0 {
		fmt.Fprintln(os.Stderr, ErrOneURL)
		os.Exit(1)
	}

	targets, err := target.ParseAll(flag.Args())
	if err != nil {
		fmt.Fprintln(os.Stderr, ErrEmptyTags)
		os.Exit(1)
	}

	output, err := reacher.ReachTargets(targets, tags, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("unexpected error %s", err))
		os.Exit(1)
	}
	fmt.Print(strings.Join(output, "\n"))
	fmt.Println()
}

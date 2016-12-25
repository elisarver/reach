package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/elisarver/reach/reacher"
	"github.com/elisarver/reach/target"
)

var (
	ErrOneURL = errors.New("please supply at least one URL")
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

	targets, err := argTargets(flag.Args())
	if err != nil {
		flag.Usage()
		fmt.Fprintln(os.Stderr, fmt.Errorf("%s", err.Error()))
		os.Exit(1)
	}

	output, err := reacher.ReachTargets(targets, pTag, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("unexpected error %s", err.Error()))
	}
	fmt.Print(strings.Join(output, "\n"))
	fmt.Println()
}

// argTargets filters the incoming argument array,
// verifying that the parameters are able to provide
// request URIs. Returns a list of targets or an error.
// Always check error! In order to return a good default
// type, we pass back a slice of Target. This is not a
// pointer, so a nil check is unnecessary. Target aliases
// to strings verified by this function.
func argTargets(args []string) ([]target.Target, error) {
	if len(args) == 0 {
		return []target.Target{}, ErrOneURL
	}
	return target.ParseAll(args)
}

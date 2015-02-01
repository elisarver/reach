package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
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
	errorHeader = "\nERROR:\n\n  "
)

func init() {
	const (
		defaultTag = "a"
		tagUsage   = "Tag to search for."
	)

	flag.Usage = func() {
		cmd := filepath.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "Reach is a tool to gather urls from a website.\n\n")
		fmt.Fprintf(os.Stderr,
			"Usage:\n\n  %s [-t=\"a\" | -tag=\"img\"] URLs...\n", cmd)
		fmt.Fprintf(os.Stderr, examples)
	}

	flag.StringVar(&pTag, "tag", defaultTag, tagUsage)
	flag.StringVar(&pTag, "t", defaultTag, tagUsage+" (Shorthand)")
}

func main() {
	flag.Parse()

	ts, err := argTargets(flag.Args())
	trap(err)

	var (
		output = make([]string, len(ts))
		tag    = NewTag(pTag)
	)

	for i, t := range ts {
		resp, err := Reach(t)
		trap(err)

		URLs := FindMap(resp, tag)

		output[i] = strings.Join(dropEmpties(URLs), "\n")
	}
	fmt.Print(strings.Join(output, "\n"))
	fmt.Println()
}

// argTargets filters the incoming argument array,
// verifying that the parameters are able to provide
// request URIs. Returns a list of targets or an error.
// Always check error! In order to return a good default
// type, we pass back a slice of Target. This is not a
// pointer, so a nil check is unnecessary. A target is
// an alias to strings verified by this function.
func argTargets(args []string) ([]Target, error) {
	numArgs := len(args)
	if numArgs == 0 {
		return []Target{}, fmt.Errorf("Please supply at least one URL.")
	}
	t := Target{&url.URL{}}
	return t.ParseAll(args)
}

func (t Target) ParseAll(args []string) ([]Target, error) {
	ts := make([]Target, 0, len(args))
	for i, _ := range args {
		if targ, err := t.Parse(args[i]); err != nil {
			return []Target{}, err
		} else {
			ts = append(ts, targ)
		}
	}
	return ts, nil
}

// dropEmpties simply eliminates empty values from a list of strings.
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
		fmt.Fprint(os.Stderr, fmt.Errorf(errorHeader+err.Error()+"\n\n"))
		os.Exit(1)
	}
}

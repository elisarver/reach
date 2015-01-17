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
	pTag   string
	pLocal bool
)

const (
	examples = `
Examples:

  Get all img src from a web page:

  > reach -l -t img http://blog.golang.org
  http://blog.golang.org/gophergala/fancygopher.jpg

  Get all unique local links on a page:

  > reach -l http://example.com/ | sort | uniq
  http://example.com/blog
  http://example.com/about

`
	errorHeader = "\nERROR:\n\n  "
)

func init() {
	const (
		defaultTag   = "a"
		tagUsage     = "Tag to search for."
		defaultLocal = false
		localUsage   = "Only display local links."
	)

	flag.Usage = func() {
		cmd := filepath.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "Reach is a tool to gather urls from a website.\n\n")
		fmt.Fprintf(os.Stderr,
			"Usage:\n\n  %s [-l | -local] [-t=\"a\" | -tag=\"img\"] URLs...\n", cmd)
		fmt.Fprintf(os.Stderr, examples)
	}

	flag.StringVar(&pTag, "tag", defaultTag, tagUsage)
	flag.StringVar(&pTag, "t", defaultTag, tagUsage+" (Shorthand)")
	flag.BoolVar(&pLocal, "local", defaultLocal, localUsage)
	flag.BoolVar(&pLocal, "l", defaultLocal, localUsage+" (Shorthand)")
}

func main() {
	flag.Parse()

	numArgs, err := chkArgs(flag.Args())
	trap(err)

	var output = make([]string, 0, numArgs)
	var rr Reacher
	for _, arg := range flag.Args() {
		rr = Reacher{
			Local: pLocal,
			Tag:   pTag,
		}

		var err error

		// checked in chkArgs, ignoring error
		rr.BaseURL, _ = url.ParseRequestURI(arg)

		URLs, err := rr.Reach()
		trap(err)

		output = append(output, strings.Join(dropEmpties(URLs), "\n"))
	}
	fmt.Print(strings.Join(output, "\n"))
	fmt.Println()
}

func chkArgs(args []string) (int, error) {
	numArgs := len(args)
	if numArgs == 0 {
		return 0, fmt.Errorf("Please supply at least one URL.")
	}
	for _, a := range args {
		_, err := chkURL(a)
		if err != nil {
			return 0, err
		}
	}
	return numArgs, nil
}

func chkURL(u string) (*url.URL, error) {
	var result *url.URL
	result, err := url.ParseRequestURI(u)
	if err != nil {
		return nil, fmt.Errorf("URL %q is mal-formed %q.", u, err)
	}
	return result, nil
}

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

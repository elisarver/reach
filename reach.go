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
	conf Reacher
)

const (
	examples = `
EXAMPLES:

  Get all img src from a web page:

  > reach -l -t img http://blog.golang.org
  http://blog.golang.org/gophergala/fancygopher.jpg

  Get all unique local links on a page:

  > reach -l http://example.com/ | sort | uniq
  http://example.com/blog
  http://example.com/about

`
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
		fmt.Fprintf(os.Stderr,
			"%s: usage: %s [-l | -local] [-t=\"a\" | -tag=\"img\"] URLs...\n", cmd, cmd)
		fmt.Fprintf(os.Stderr, examples)
	}

	var pTag string
	var pLocal bool

	flag.StringVar(&pTag, "tag", defaultTag, tagUsage)
	flag.StringVar(&pTag, "t", defaultTag, tagUsage+" (Shorthand)")
	flag.BoolVar(&pLocal, "local", defaultLocal, localUsage)
	flag.BoolVar(&pLocal, "l", defaultLocal, localUsage+" (Shorthand)")
	flag.Parse()

	conf = Reacher{
		Local: pLocal,
		Tag:   pTag,
	}
}

func main() {
	numArgs := chkArgs(flag.Args())
	var output = make([]string, 0, numArgs)
	for _, a := range flag.Args() {
		c := conf // defensive copy; done several times.
		var err error
		if c.BaseURL, err = chkURL(a); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}

		URLs, err := c.Reach()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		cleaned := dropEmpties(URLs)
		output = append(output, strings.Join(cleaned, "\n"))
	}
	fmt.Print(strings.Join(output, "\n"))
	fmt.Println()
}

func chkArgs(args []string) int {
	numArgs := len(args)
	if numArgs == 0 {
		flag.Usage()
		fmt.Fprint(os.Stderr, "\nERROR:\n\n  Please supply at least one URL.\n\n")
		os.Exit(1)
	}
	return numArgs
}

func chkURL(u string) (*url.URL, error) {
	url, err := url.ParseRequestURI(u)
	if err != nil {
		err = fmt.Errorf("URL '%s': is mal-formed, stopping.\n", u)
		return nil, err
	}
	return url, nil
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

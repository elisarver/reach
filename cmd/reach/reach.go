package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/elisarver/reach/document"
	"github.com/elisarver/reach/tag"
	"github.com/elisarver/reach/target"
)

func main() {
	var pTag string
	flag.StringVar(&pTag, "tag", "a", "comma-separated list of `tags` to search for.\n\tformat: name1[:attribute1][,name2[:attribute2]]")

	var pQuery string
	flag.StringVar(&pQuery, "query", "", "enter a raw query, such as 'ul>li'")

	var help bool
	flag.BoolVar(&help, "h", false, "print help")

	var reparent bool
	flag.BoolVar(&reparent, "p", false, "reparent relative URIs to request domain")

	flag.Parse()

	if help {
		fmt.Fprint(os.Stderr,
			"Description:\n\treach visits URLs and returns tags/attributes of interest.\n")
		flag.Usage()
		fmt.Fprint(os.Stderr,
			"Positional arguments:\n  URL [URL...]\n\tone or more URLs to dial.\n\n")
		os.Exit(0)
	}

	if len(flag.Args()) == 0 {
		exitErr(errors.New("please supply at least one URL"))
	}

	targets, err := target.ParseAll(flag.Args()...)
	exitErr(err)

	var tags tag.DescriptionSlice
	if pQuery != "" {
		tags = tag.RawQuery(pQuery)
	} else {
		tags = tag.FromMultiSpec(pTag)
	}
	processor := document.NewProcessor(nil, reparent)
	output, err := processor.Process(tags, targets)
	exitErr(err)
	fmt.Print(strings.Join(output, "\n"))

	fmt.Println()
}

func exitErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, "try `reach -h` for usage.")
		os.Exit(1)
	}
}

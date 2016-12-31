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
	errOneURL    = errors.New("please supply at least one URL")
	errEmptyTags = errors.New("empty tag arguments")
)

func main() {
	var pTag string
	flag.StringVar(&pTag, "tag", "a", "comma-separated list of `tags` to search for.\n\tformat: name1[:attribute1][,name2[:attribute2]]")
	var help bool
	flag.BoolVar(&help, "h", false, "print help")
	flag.Parse()

	if help {
		fmt.Fprint(os.Stderr,
			"Description:\n\treach visits URLs and returns tags/attributes of interest.\n")
		flag.Usage()
		fmt.Fprint(os.Stderr,
			"Positional arguments:\n  URL [URL...]\n\tone or more URLs to dial.\n\n")
		os.Exit(0)
	}

	tags := tag.FromMultiSpec(pTag)

	if len(flag.Args()) == 0 {
		exitErr(errOneURL)
	}

	targets, err := target.ParseAll(flag.Args())
	exitErr(err)

	output, err := reacher.ReachTargets(targets, tags, nil)
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

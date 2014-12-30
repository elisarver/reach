package reach

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

var (
	conf Reacher
)

func init() {
	var pElement string
	var pLocal bool
	flag.StringVar(&pElement, "t", "a", "Tag to search for.")
	flag.StringVar(&pElement, "tag", "a", "Tag to search for. (long)")
	flag.BoolVar(&pLocal, "l", false, "Only display local links.")
	flag.BoolVar(&pLocal, "local", false, "Only display local links (long)")
	flag.Parse()

	conf = Reacher{
		Local:   pLocal,
		Element: pElement,
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
		fmt.Fprintln(os.Stderr, "Please supply at least one URL.")
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

Reach is a tool to gather urls from a website.

Usage:

  reach [-t="a" | -tag="img"] URLs...

Examples:

  Get all img src from a web page:

````
  > reach -t img http://blog.golang.org
  http://blog.golang.org/gophergala/fancygopher.jpg
````

  Get all unique links on a page:

````
  > reach http://example.com/ | sort | uniq
  http://example.com/blog
  http://example.com/about
````

Installing: go get github.com/elisarver/reach

Contact: eli@elisarver.com

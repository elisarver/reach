Reach gathers urls from a website.

Installation:

I suggest you install [glide](https://github.com/Masterminds/glide), and run glide install.

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

Contact: eli.sarver@gmail.com

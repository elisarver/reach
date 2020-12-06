Reach gathers urls from a website.

Installation:


You can use `go get -u github.com/elisarver/reach/cmd/reach` if you're only interested in the executable.

We build with `go.mod` support under go v1.11.  

The application displays full usage with the -h flag.

- Depends on [goquery](https://github.com/PuerkitoBio/goquery) for querying html
- Uses [gen](https://github.com/clipperhouse/gen) to generate custom sets and slices

Example of default use (gets `a\[href]` by default):

````
> reach http://google.com/
http://google.com/intl/en/policies/privacy/
http://google.com/intl/en/policies/terms/
/intl/en/ads/
/services/
````

Reparenting urls:
````
> reach -p http://google.com/
http://google.com/intl/en/policies/privacy/
http://google.com/intl/en/policies/terms/
http://google.com/intl/en/ads/
http://google.com/services/
````
Specify tag type:
````
> reach -tag meta:name http://google.com/
description
robots
````

Use a raw query:
````
> reach -query "div.all_external_links" http://elisarver.com/another-week-of-accomplishment

// Document body
````

Contact: eli.sarver@gmail.com

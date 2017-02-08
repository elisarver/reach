Reach gathers urls from a website.

Installation:

You can install [glide](https://github.com/Masterminds/glide), and run glide install, if you like automatic dependency fulfillment.

You can use `go get -u github.com/elisarver/reach/cmd/reach` if you're only interested in the executable.

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

Contact: eli.sarver@gmail.com

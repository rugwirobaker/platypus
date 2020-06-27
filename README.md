[![Build Status](https://cloud.drone.io/api/badges/rugwirobaker/platypus/status.svg)](https://cloud.drone.io/rugwirobaker/platypus)
[![codebeat badge](https://codebeat.co/badges/aeb375cb-8061-44de-8ac1-6ce1a5e8500f)](https://codebeat.co/projects/github-com-rugwirobaker-platypus-master)
[![codecov](https://codecov.io/gh/rugwirobaker/platypus/branch/master/graph/badge.svg)](https://codecov.io/gh/rugwirobaker/platypus)
## Platypus

platypus is a router but for ussd input. It follows the style of http routers in go.


## Roadmap

1. support subrouting.
2. improve godoc.
3. improve test coverage.
4. add a queueing server?
5. Suggestions are welcome(open an issue) ....

## Usage
Note that to endicate that a submenu is the last on it's chain you must register it `#`.
```
....
 mux := platypus.New(prefix, platypus.HandlerFunc(notFoundHandler))
```
### Note

It's still work in progress but PRs are so welcome.
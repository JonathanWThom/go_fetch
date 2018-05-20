# Go Fetch

Send a url, and receive the time it took to fetch it back. For example:

## Implementation

`git clone && cd go_fetch`
`go run`

Visit: `http://localhost:8000/speed/www.google.com`
You'll see a number!

# Todo

- Add error handling
- Return something other than microseconds (or nanoseconds).
- Determine what actually should be recorded (the fetch is very fast, but should it be looking for something
else, like when the DOM is loaded?)
- Build in multiple runs - like 5+ tests of an endpoint, and send back and average.
- Handle authentication?? Could someone use this for a site of theirs that has auth?

# Go Fetch

Send a url, and receive the time it took to fetch it back.

## Implementation

`git clone && cd go_fetch`

`go run`

Visit: `http://localhost:8000/speed/www.google.com`
You'll see a number!

# Todo

- Handle multiple endpoint hits better - should run concurrently so that it's not a slow api call.
- Handle authentication?? Could someone use this for a site of theirs that has auth?
- Deploy and try it live.
- Build a front end?
- Schedule runs and return to an app as a webhook.
- Add tests.

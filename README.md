# S3 Proxy

This is the simplest AWS S3 proxy. It proxies the contents for an specific
bucket. If the request path ends with a slash (/), it will try to get
`index.html`, or the specified index page.

It doesn't have a cache, nor anything fancy on top of it. It was initially
intended to be used internally and not exposed to the Internet. If you want to
expose a full bucket, using the built-in S3 website would work better. Be aware
that the contents of the buckets will be accessible through this server.

## Configuration

It can be configured either with environment variables or passing arguments.

* `HOST`, or `-host <host>`: The host to where listen to (default: `":8080"`)
* `INDEX_PAGE` or `-index <index>`: The name of the index pages
  (default: `"index.html"`)
* `S3_BUCKET`, or `-bucket <bucket>`: The bucket to proxy, mandatory


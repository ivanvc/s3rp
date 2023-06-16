# S3 Reverse Proxy

Simplest AWS S3 proxy. It proxies the contents for an specific bucket. If the
request path ends with a slash (/), it will try to get `index.html`, or the
specified index page.

It doesn't have a cache, nor anything fancy on top of it. Internally it just
signs the request using the provided AWS credentials, and Go's reverse
proxy does the rest. It was initially intended for internal use and not
exposed to the Internet. If you want to expose a full bucket, using the built-in
S3 website would work better. Be aware that the contents of the buckets will be
accessible through this server.

## Usage

It can be used by using the Docker image [`ivan/s3-reverse-proxy`][docker], or
can be installed using Go, by doing:

```bash
go install github.com/ivanvc/s3rp
```

[docker]: https://hub.docker.com/r/ivan/s3-reverse-proxy

### Configuration

It can be configured either with environment variables or passing arguments.

* `HOST`, or `-host <host>`: The host to where listen to (default: `":8080"`)
* `INDEX_PAGE` or `-index <index>`: The name of the index pages
  (default: `"index.html"`)
* `BUCKET`, or `-bucket <bucket>`: The bucket to proxy, mandatory

Authentication with AWS is done by the AWS SDK. It can read environment
variables, or your aws-cli credentials, instance role, etc.

### Running

If using docker, you can run it with (adjust values as required):

```bash
docker run -e BUCKET=my-bucket -e AWS_REGION=us-west-2 AWS_ACCESS_KEY_ID=xx \
AWS_SECRET_ACCESS_KEY=xx ivan/s3-reverse-proxy
```

Or directly with Go, by doing:

```bash
s3rp -bucket=my-bucket
```

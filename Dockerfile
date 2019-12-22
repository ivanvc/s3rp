FROM golang:1.13-alpine AS build

WORKDIR /app
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN apk add -U --no-cache ca-certificates
COPY go.mod go.sum /app/
RUN go mod download

COPY cmd/ /app/cmd/
COPY pkg/ /app/pkg/
RUN go install -gcflags "all=-N -l" ./...

FROM scratch
ENTRYPOINT ["/usr/bin/s3rp"]
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/bin/s3rp /usr/bin/s3rp

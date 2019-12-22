FROM golang:1.13 AS build

WORKDIR /app
ENV CGO_ENABLED=0 GOOS=linux
COPY go.mod go.sum /app/
RUN go mod download

COPY cmd/ /app/cmd/
COPY pkg/ /app/pkg/
RUN go install ./...

FROM scratch
ENTRYPOINT /usr/bin/s3rp
COPY --from=build /go/bin/s3rp /usr/bin/s3rp

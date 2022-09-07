FROM golang:1.19.1-alpine
RUN apk add git
ENV CGO_ENABLED=0
WORKDIR /
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go install -ldflags "-X main.version=$(git describe --tags)"

FROM alpine:3.16.2
ENV ICALPROXY_LISTEN_ADDR=0.0.0.0:8080
COPY --from=0 /go/bin/ical-proxy /usr/local/bin/
ENTRYPOINT [ "ical-proxy" ]

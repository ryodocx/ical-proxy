FROM golang:1.18.4-alpine
ENV CGO_ENABLED=0
WORKDIR /
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build .

FROM alpine:3.16.0
ENV ICALPROXY_LISTEN_ADDR=0.0.0.0:8080
COPY --from=0 /ical-proxy .
ENTRYPOINT [ "/ical-proxy" ]

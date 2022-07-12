FROM golang:1.18.3-alpine
ENV CGO_ENABLED=0
WORKDIR /var/build
COPY . .
RUN go build .

FROM alpine:3.16.0
ENV LISTEN_ADDR=0.0.0.0:8080
COPY --from=0 /var/build/ical-proxy .
CMD [ "/ical-proxy" ]

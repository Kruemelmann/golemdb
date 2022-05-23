FROM golang:alpine as builder

COPY . /app
WORKDIR /app/cmd/golemdb
RUN go build .

FROM alpine as prod

COPY --from=builder /app/cmd/golemdb/golemdb /app
CMD ["/app"]


FROM golang:1.15 as builder
WORKDIR /go/src/github.com/kruemelmann/golemdb
COPY . /app
WORKDIR /app/cmd/golemdb
# `skaffold debug` sets SKAFFOLD_GO_GCFLAGS to disable compiler optimizations
ARG SKAFFOLD_GO_GCFLAGS
RUN go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -o /golemdb .


FROM alpine:3
# Define GOTRACEBACK to mark this container as using the Go language runtime
# for `skaffold debug` (https://skaffold.dev/docs/workflows/debug/).
ENV GOTRACEBACK=single
COPY --from=builder /golemdb /golemdb
CMD ["/golemdb"]

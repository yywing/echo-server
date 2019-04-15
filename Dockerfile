FROM golang:1.11.9 as builder
WORKDIR /go/src/echo-server
COPY . /go/src/echo-server
RUN  GO111MODULE=on go build -mod=vendor -o echo-server

FROM golang:1.11.9 as prod
WORKDIR /app
COPY --from=builder /go/src/echo-server/echo-server .
ENTRYPOINT ["/app/echo-server"]

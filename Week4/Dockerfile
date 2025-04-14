FROM golang:1.20.1-alpine3.17 AS builder
WORKDIR /work
COPY . /work/
RUN go build sleep.go

FROM alpine:3.14
COPY --from=builder /work/sleep /work/sleep
ENTRYPOINT ["/work/sleep"]

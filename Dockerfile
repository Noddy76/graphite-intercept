FROM golang:1.14-alpine AS builder

RUN apk update && apk add --no-cache git build-base

RUN mkdir /build 
ADD . /build/
WORKDIR /build
RUN GOOS=linux GOARCH=amd64 go build -ldflags "-linkmode external -extldflags -static" -o graphite-intercept graphite-intercept.go

FROM scratch
COPY --from=builder /build/graphite-intercept /graphite-intercept
WORKDIR /
ENTRYPOINT ["/graphite-intercept"]
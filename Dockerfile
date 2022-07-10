FROM golang:1.17.11-alpine3.16 AS build

ENV GO111MODULE=on

WORKDIR /

COPY . /go/src/github.com/ryokky59/go-layered-architecture-sample

RUN apk update && apk add --no-cache git
RUN cd /go/src/github.com/ryokky59/go-layered-architecture-sample/api && go build -o bin/sample main.go

FROM alpine:3.16

COPY --from=build /go/src/github.com/ryokky59/go-layered-architecture-sample/api/bin/sample /usr/local/bin/

CMD ["sample"]

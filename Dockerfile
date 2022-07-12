FROM golang:1.17.11-alpine3.16 AS build

ENV GO111MODULE=on

WORKDIR /

COPY . /go/src/github.com/katsukiniwa/practical-go-programming

RUN apk update && apk add --no-cache git
RUN cd /go/src/github.com/katsukiniwa/practical-go-programming && go build -o bin/practical-go-programming main.go

FROM alpine:3.16

COPY --from=build /go/src/github.com/katsukiniwa/practical-go-programming/api/bin/practical-go-programming /usr/local/bin/

CMD ["practical-go-programming"]

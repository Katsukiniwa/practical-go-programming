FROM golang:1.17.1-alpine as builder

RUN apk update && apk upgrade && \
  apk --update add git make

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o engine server.go

FROM alpine:latest

RUN apk update && apk upgrade && \
  apk --update --no-cache add tzdata && \
  mkdir /app

WORKDIR /app

COPY --from=builder /app/engine /app

CMD /app/engine

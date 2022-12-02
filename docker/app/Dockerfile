FROM golang:1.19-alpine as builder

RUN apk update && apk upgrade && \
  apk --update add git make vim

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 go build -o engine server.go

FROM alpine:latest

RUN apk update && apk upgrade && \
  apk --update --no-cache add tzdata

WORKDIR /app

COPY --from=builder /app/engine /app

CMD /app/engine

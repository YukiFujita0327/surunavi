FROM golang:1.15.6-alpine3.12

WORKDIR /go/src/app

COPY ./ /go/src/app

RUN apk update && apk add git

CMD ["go", "run", "server.go"]

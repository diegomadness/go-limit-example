ARG GO_VERSION="1.15.8"
FROM golang:$GO_VERSION

ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on

WORKDIR /go/src/github.com/diegomadness/limit
ADD . .

CMD go run github.com/diegomadness/limit

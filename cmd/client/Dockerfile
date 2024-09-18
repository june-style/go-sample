FROM golang:1.21

ENV GO111MODULE=on

ENV GOPATH /go
ENV GOSRC ${GOPATH}/src

WORKDIR ${GOSRC}/github.com/june-style/go-sample

CMD [ "go", "run", "./cmd/client/main.go" ]

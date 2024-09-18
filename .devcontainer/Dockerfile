FROM golang:1.21

ENV GOPATH /go
ENV GOBIN /go/bin
ENV APP_PATH ${GOPATH}/src/github.com/june-style/go-sample
ENV GO111MODULE=on

WORKDIR ${APP_PATH}

RUN apt update \
    && apt install -y protobuf-compiler

RUN apt-get update \
    && apt-get install -y bash-completion \
    && apt-get install -y vim

RUN curl -L "https://github.com/docker/compose/releases/download/v2.23.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose \
    && chmod +x /usr/local/bin/docker-compose

CMD [ "/bin/bash" ]

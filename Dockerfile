FROM golang:1.15

ENV GOPATH=/go
ENV GO111MODULE=on

ENV APP=/mutants
ENV PROJECT_ROOT=$APP

WORKDIR $APP
COPY . $APP

RUN go get ./...
RUN go get github.com/githubnemo/CompileDaemon

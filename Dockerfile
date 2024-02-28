FROM golang:latest

RUN apt update && \
    apt install sqlite3


WORKDIR /go/src/
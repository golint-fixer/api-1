# FROM golang:1.6.2-wheezy
FROM golang:1.6.2-alpine

MAINTAINER dodd.anthonyjosiah@gmail.com

EXPOSE 3000

# Ensure godep is in place for dependency management.
RUN apk add --update git wget tar
RUN go get github.com/codegangsta/gin
RUN go get github.com/tools/godep

# Copy over needed files.
WORKDIR /go/src/github.com/thedodd/api
COPY main.go main.go
COPY common common
COPY elasticsearch elasticsearch
COPY users users
COPY Godeps Godeps

# Build our API.
RUN godep get && go install .

# Use a CMD here, instead of ENTRYPOINT, for easy overwrite in docker ecosystem.
CMD api

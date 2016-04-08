# The build API.
FROM golang:1.6.0-alpine

MAINTAINER dodd.anthonyjosiah@gmail.com

EXPOSE 3000
ENV GO15VENDOREXPERIMENT=1

# Ensure glide is in place for dependency management.
RUN apk update && apk add git
RUN go get github.com/Masterminds/glide
RUN go get github.com/codegangsta/gin

# Copy over needed files.
WORKDIR /go/src/github.com/thedodd/buildAPI
COPY ./main.go main.go
COPY ./elasticsearch elasticsearch
COPY ./common common
COPY ./glide.yaml glide.yaml
COPY ./glide.lock glide.lock

# Build our API.
RUN glide install && go install github.com/thedodd/buildAPI

# Use a CMD here, instead of ENTRYPOINT, for easy overwrite in docker ecosystem.
CMD buildAPI

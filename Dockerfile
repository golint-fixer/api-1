# The build API.
FROM golang:1.5.2-wheezy

MAINTAINER dodd.anthonyjosiah@gmail.com

ENV GO15VENDOREXPERIMENT=1

# Copy over needed files.
WORKDIR /go/src/github.com/thedodd/buildAPI
COPY ./main.go main.go
COPY ./elasticsearch elasticsearch
COPY ./glide.yaml glide.yaml
COPY ./glide.lock glide.lock

# Ensure glide is in place for dependency management.
RUN go get github.com/Masterminds/glide

# Build our API.
RUN glide install && go install github.com/thedodd/buildAPI

EXPOSE 3000

# Use a CMD here, instead of ENTRYPOINT, for easy overwrite in docker ecosystem.
CMD buildAPI

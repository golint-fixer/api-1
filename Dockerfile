FROM golang:1.6.0-alpine

MAINTAINER dodd.anthonyjosiah@gmail.com

EXPOSE 3000
ENV GO15VENDOREXPERIMENT=1

# Ensure glide is in place for dependency management.
RUN apk update && apk add git
RUN go get github.com/Masterminds/glide
RUN go get github.com/codegangsta/gin

# Copy over needed files. Filter with .dockerignore.
WORKDIR /go/src/github.com/thedodd/api
COPY . .

# Build our API.
RUN glide install && go install .

# Use a CMD here, instead of ENTRYPOINT, for easy overwrite in docker ecosystem.
CMD api

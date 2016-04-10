api
===
[![Build Status](https://travis-ci.org/thedodd/api.svg?branch=master)](https://travis-ci.org/thedodd/api)
[![Code Climate](https://codeclimate.com/github/TheDodd/buildAPI/badges/gpa.svg)](https://codeclimate.com/github/TheDodd/buildAPI)

Just a Golang API to hack on.

### development
Docker is used for all aspects of this projects development and deployment. It is assumed that you have a docker daemon available to work with. To get this API up and running, simply execute the following:

```bash
# Stand up services & tail container logs.
docker-compose up -d && docker-compose logs
```

Now you can interface with the API at port `8080` on your docker daemon's host.

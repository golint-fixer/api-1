[![Circle CI](https://circleci.com/gh/TheDodd/buildAPI.svg?style=svg)](https://circleci.com/gh/TheDodd/buildAPI)
[![Code Climate](https://codeclimate.com/github/TheDodd/buildAPI/badges/gpa.svg)](https://codeclimate.com/github/TheDodd/buildAPI)

api
===
Just a Golang API to hack on.

### development
Docker is used for all aspects of this projects development and deployment. It is assumed that you have a docker daemon available to work with. To get this API up and running, simply execute the following:

```bash
docker-compose up -d
```

Now you can interface with the API at port `8080` on your docker daemon's host.

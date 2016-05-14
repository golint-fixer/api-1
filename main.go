// main - API entry point.
package main

import (
	"github.com/gin-gonic/gin"

	"github.com/thedodd/api/common"
	"github.com/thedodd/api/elasticsearch"
	"github.com/thedodd/api/users"
)

func main() {
	api := gin.Default()
	api.Use(common.JSONOnlyAPI)

	// V1 of the API.
	v1 := api.Group("")
	v1.Use(common.RequestID)
	users.BindRoutes(v1, "/users/")
	elasticsearch.BindRoutes(v1, "/elasticsearch/")

	// Fire this bad boy up.
	api.Run("0.0.0.0:3000")
}

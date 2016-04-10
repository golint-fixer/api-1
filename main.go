package main

import (
	"github.com/gin-gonic/gin"

	"github.com/thedodd/api/common"
	"github.com/thedodd/api/elasticsearch"
	"github.com/thedodd/api/users"
)

func main() {
	router := gin.Default()

	// V1 of the API.
	v1 := router.Group("")
	v1.Use(common.RequestID)
	v1.Use(common.BasicAuthRequired)

	// Register Elasticsearch builds resource handlers.
	usersRouter := v1.Group("/users")
	{
		usersRouter.POST("/", common.ValidateInboundJSON(&users.User{}), users.CreateUser)
	}

	// Register Elasticsearch builds resource handlers.
	esBuildsRouter := v1.Group("/elasticsearch/builds")
	{
		esBuildsRouter.GET("/", elasticsearch.GetElasticsearchBuilds)
		esBuildsRouter.POST("/", common.ValidateInboundJSON(&elasticsearch.BuildModel{}), elasticsearch.CreateElasticsearchBuild)
		esBuildsRouter.GET("/:id", elasticsearch.GetElasticsearchBuildByID)
	}

	// Fire this bad boy up.
	router.Run("0.0.0.0:3000")
}

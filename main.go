package main

import (
	"github.com/gin-gonic/gin"

	"github.com/thedodd/api/common"
	"github.com/thedodd/api/elasticsearch"
)

func main() {
	router := gin.Default()

	// V1 of the API.
	v1 := router.Group("")
	v1.Use(common.RequestID)
	v1.Use(common.BasicAuthRequired)

	// Register Elasticsearch builds resource handlers.
	elasticsearchBuilds := v1.Group("/elasticsearch/builds")
	{
		elasticsearchBuilds.GET("/", elasticsearch.GetElasticsearchBuilds)
		elasticsearchBuilds.POST("/", common.ValidateInboundJSON(&elasticsearch.BuildModel{}), elasticsearch.CreateElasticsearchBuild)
		elasticsearchBuilds.GET("/:id", elasticsearch.GetElasticsearchBuildByID)
	}

	// Fire this bad boy up.
	router.Run("0.0.0.0:3000")
}

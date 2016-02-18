/**
 * Just a build API.
 */

package main

import (
	"github.com/gin-gonic/gin"

	"github.com/thedodd/buildAPI/elasticsearch"
)

func main() {
	router := gin.Default()

	// Register Elasticsearch builds resource handlers.
	elasticsearchBuilds := router.Group("/elasticsearch/builds")
	{
		elasticsearchBuilds.GET("/", elasticsearch.GetElasticsearchBuilds)
		elasticsearchBuilds.POST("/", elasticsearch.CreateElasticsearchBuild)
		elasticsearchBuilds.GET("/:id", elasticsearch.GetElasticsearchBuildByID)
	}

	// Fire this bad boy up.
	router.Run("0.0.0.0:3000")
}

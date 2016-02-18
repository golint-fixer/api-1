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
	router.GET("/elasticsearch/builds/", elasticsearch.GetElasticsearchBuilds)
	router.POST("/elasticsearch/builds/", elasticsearch.CreateElasticsearchBuild)
	router.GET("/elasticsearch/builds/:id", elasticsearch.GetElasticsearchBuildByID)

	// Fire this bad boy up.
	router.Run("0.0.0.0:3000")
}

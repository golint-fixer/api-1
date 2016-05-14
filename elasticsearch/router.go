package elasticsearch

import (
	"github.com/gin-gonic/gin"

	"github.com/thedodd/api/common"
)

// BindRoutes bind this package's routes to the given gin.RouterGroup at the given base URI.
func BindRoutes(baseRouter *gin.RouterGroup, baseURI string) {
	router := baseRouter.Group(baseURI)

	// Register routes.
	router.Use(common.BasicAuthRequired) // Protect these resources with basic auth.
	router.GET("/builds/", GetElasticsearchBuilds)
	router.POST("/builds/", common.BindJSON(&BuildModel{}), CreateElasticsearchBuild)
	router.GET("/builds/:id", GetElasticsearchBuildByID)
}

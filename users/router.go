package users

import (
	"github.com/gin-gonic/gin"

	"github.com/thedodd/api/common"
)

// BindRoutes bind this package's routes to the given gin.RouterGroup at the given base URI.
func BindRoutes(baseRouter *gin.RouterGroup, baseURI string) {
	router := baseRouter.Group(baseURI)

	// Register routes.
	router.POST("/", common.BindJSON(&User{}), CreateUser)
	router.POST("/:username/verify", common.BindJSON(&VerificationBody{}), VerifyUser)
}

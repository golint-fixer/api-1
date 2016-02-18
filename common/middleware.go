package common

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BasicAuthRequired enforce basic auth on resources.
func BasicAuthRequired(context *gin.Context) {
	username, password, credentialsProvided := context.Request.BasicAuth()
	if !credentialsProvided {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "No credentials provided."})
		context.Abort()
		return
	}

	// TODO(TheDodd): build out a real authN/Z system here.
	log.Println(username, password)
	context.Set("username", username)
	context.Set("id", username)

	// Before request.
	context.Next()
	// After request.
}

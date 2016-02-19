package common

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"gopkg.in/go-playground/validator.v8"
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
	log.Println("DELETE THIS LOG STATEMENT", username, password)
	context.Set("username", username)
	context.Set("id", username)

	// Yield to other middleware handlers.
	context.Next()
}

// RequestID tag the current request with an ID & add a response X-Request-Id header.
func RequestID(context *gin.Context) {
	// Generate an ID for this request.
	id := uuid.NewV4().String()

	// Bind request ID to context.
	context.Set("request_id", id)
	context.Writer.Header().Set("X-Request-ID", id)

	// Yield to other middleware handlers.
	context.Next()
}

// ValidateInboundJSON validate any inbound JSON against the given model pointer.
func ValidateInboundJSON(model ModelInterface) gin.HandlerFunc {
	return func(context *gin.Context) {
		if err := context.BindJSON(model); err != nil {
			validationErrors := err.(validator.ValidationErrors) // Will always be this type.
			model.HandleValidationErrors(context, validationErrors)
			context.Abort()
			return
		}

		// Bind validated model pointer to request context.
		context.Set("data", model)

		// Yield to other middleware handlers.
		context.Next()
	}
}

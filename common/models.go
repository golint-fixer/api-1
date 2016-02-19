package common

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
)

// ModelInterface defines the interface for data models.
type ModelInterface interface {
	// Build up and serialize a JSON response of errors, but do not call context.Abort yet.
	HandleValidationErrors(*gin.Context, validator.ValidationErrors)
}

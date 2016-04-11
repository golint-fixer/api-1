package common

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
	"gopkg.in/mgo.v2"
)

// ModelInterface - the interface definition for data models.
type ModelInterface interface {
	// The collection of the model.
	Collection() *mgo.Collection

	// Ensure indices needed by the model are in place.
	EnsureIndices()

	// Build up and serialize a JSON response of errors, but do not call context.Abort yet.
	HandleValidationErrors(*gin.Context, validator.ValidationErrors)
}

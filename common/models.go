package common

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
	"gopkg.in/mgo.v2"
)

// ModelInterface defines the interface for data models.
type ModelInterface interface {
	Collection() *mgo.Collection

	EnsureIndices()

	// Build up and serialize a JSON response of errors, but do not call context.Abort yet.
	HandleValidationErrors(*gin.Context, validator.ValidationErrors)
}

// SerializeValidationErrors default serialization for model validation errors.
func SerializeValidationErrors(model ModelInterface, errors validator.ValidationErrors) (collector []map[string]string) {
	collector = make([]map[string]string, 0, len(errors))
	reflectTypeElem := reflect.TypeOf(model).Elem()
	for _, fieldError := range errors {
		reflectField, _ := reflectTypeElem.FieldByName(fieldError.Field)
		jsonFieldName := reflectField.Tag.Get("json")
		validators := reflectField.Tag.Get("validate")
		err := map[string]string{
			"field":      jsonFieldName,
			"error":      "Error encountered during validation.",
			"message":    fmt.Sprintf("Validation failed for validator '%s'.", fieldError.Tag),
			"validators": validators,
		}
		collector = append(collector, err)
	}
	return collector
}

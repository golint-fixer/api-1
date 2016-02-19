package common

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
)

// ModelInterface defines the interface for data models.
type ModelInterface interface {
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
		validators := reflectField.Tag.Get("binding")
		err := map[string]string{
			"field":      jsonFieldName,
			"type":       fieldError.Type.Name(), // Name of type as string.
			"error":      fieldError.Tag,
			"validators": validators,
		}
		collector = append(collector, err)
	}
	return collector
}

// SerializeInboundTypeErrors serialize error related to inbound type mismatch.
func SerializeInboundTypeErrors(context *gin.Context, typeError *json.UnmarshalTypeError) {
	collector := []map[string]string{
		map[string]string{
			"error":        "Incorrect data type provided.",
			"expectedType": typeError.Type.Name(),
			"givenType":    typeError.Value,
		},
	}
	context.JSON(http.StatusBadRequest, gin.H{"errors": collector, "numErrors": 1})
}

package common

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/go-playground/validator.v8"
)

// Validator - the application wide struct validator. Uses struct tag "validate" for validation.
var Validator = validator.New(&validator.Config{TagName: "validate"})

// BindJSON - unmarshal & validate inbound JSON against the given serializer.
func BindJSON(model interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		// Unmarshall inbound JSON dynamically to avoid unclear type errors.
		rawUnmarshal, err := ensureJSONProvidedElseError(context)
		if err != nil {
			context.Abort()
			return
		}

		// Map raw JSON onto model.
		decoder := getDecoder(model)
		if err := decoder.Decode(rawUnmarshal); err != nil {
			serializeDecodeErrors(context, err)
			context.Abort()
			return
		}

		// Validate the populated model.
		if err := Validator.Struct(model); err != nil {
			errs, _ := err.(validator.ValidationErrors)
			serializeValidationErrors(context, model, errs)
			context.Abort()
			return
		}

		// Bind validated model pointer to request context.
		context.Set("data", model)

		// Yield to other middleware handlers.
		context.Next()
	}
}

func getDecoder(model interface{}) *mapstructure.Decoder {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "json",
		Result:  model,
	})
	if err != nil {
		panic(err)
	}
	return decoder
}

func ensureJSONProvidedElseError(context *gin.Context) (rawUnmarshal *map[string]interface{}, err error) {
	rawUnmarshal = &map[string]interface{}{}
	context.BindJSON(rawUnmarshal)

	if len(*rawUnmarshal) == 0 {
		collector := []map[string]string{
			map[string]string{
				"error":   "No JSON body found.",
				"message": "Expected JSON body to be provided.",
			},
		}
		context.JSON(http.StatusBadRequest, gin.H{"errors": collector, "numErrors": 1})
		err = errors.New("No JSON provided.")
	}
	return rawUnmarshal, err
}

func serializeDecodeErrors(context *gin.Context, err error) {
	collector := []map[string]string{}
	numErrors := 0

	switch eType := err.(type) {
	case *mapstructure.Error:
		for _, er := range eType.WrappedErrors() {
			numErrors++
			collector = append(collector, map[string]string{
				"error":   "Error encountered during decode process.",
				"message": er.Error(),
			})
		}

	default:
		numErrors++
		collector = append(collector, map[string]string{
			"error":   "Error encountered during decode process.",
			"message": err.Error(),
		})
	}

	context.JSON(http.StatusBadRequest, gin.H{"errors": collector, "numErrors": numErrors})
}

func serializeValidationErrors(context *gin.Context, model interface{}, errors validator.ValidationErrors) {
	// TODO(TheDodd): look into defining an explicit mapping for each validator.V8 validator error.
	numErrors := len(errors)
	collector := make([]map[string]string, 0, numErrors)
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

	context.JSON(http.StatusBadRequest, gin.H{"errors": collector, "numErrors": numErrors})
}

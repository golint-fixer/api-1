package common

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/go-playground/validator.v8"
)

// Validator - the application wide struct validator. Uses struct tag "validate" for validation.
var Validator = validator.New(&validator.Config{TagName: "validate"})

// ValidateInboundJSON validate any inbound JSON against the given model pointer.
func ValidateInboundJSON(model ModelInterface) gin.HandlerFunc {
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
			model.HandleValidationErrors(context, errs) // FIXME(TheDodd): this interface seems redundant.
			context.Abort()
			return
		}

		// Bind validated model pointer to request context.
		context.Set("data", model)

		// Yield to other middleware handlers.
		context.Next()
	}
}

func getDecoder(model ModelInterface) *mapstructure.Decoder {
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

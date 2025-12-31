package validator

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// func init() {
// 	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
// 		_ = v
// 	}
// }

// var v = validator.New()

// Body reads and validates the JSON body from a Gin context.
func Body(ctx *gin.Context, ptr interface{}) (map[string]string, bool) {
	if err := ctx.ShouldBindJSON(ptr); err != nil {
		errs := make(map[string]string)

		// 1. validation-tag failures
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			for _, e := range ve {
				errs["message"] = e.Field() + " " + e.Tag()
			}
			return errs, false
		}

		// 2. wrong JSON type
		var unmarshalTypeError *json.UnmarshalTypeError
		if errors.As(err, &unmarshalTypeError) {
			errs[unmarshalTypeError.Field] = "must be " + unmarshalTypeError.Type.String()
			return errs, false
		}

		// 3. anything else
		errs["body"] = "invalid JSON"
		return errs, false
	}
	return nil, true
}

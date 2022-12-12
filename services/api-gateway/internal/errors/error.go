package errors

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type response struct {
	Error any `json:"error"`
}

type validationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ErrorResponse(ctx *gin.Context, code int, msg string) {
	ctx.AbortWithStatusJSON(code, response{msg})
}

func ValidationErrorResponse(ctx *gin.Context, code int, err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]validationError, len(ve))
		for i, fe := range ve {
			out[i] = validationError{fe.Field(), msgForTag(fe.Tag())}
		}
		ctx.AbortWithStatusJSON(code, response{out})
	}
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "url":
		return "Invalid url"
	}

	return "Invalid " + tag
}

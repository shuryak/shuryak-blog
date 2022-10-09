package v1

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

func errorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, response{msg})
}

func validationErrorResponse(c *gin.Context, code int, err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]validationError, len(ve))
		for i, fe := range ve {
			out[i] = validationError{fe.Field(), msgForTag(fe.Tag())}
		}
		c.AbortWithStatusJSON(code, response{out})
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

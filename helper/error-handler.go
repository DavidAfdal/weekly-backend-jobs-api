package helper

import (
	"errors"
	"weekly/go/gin/data/response"

	"github.com/go-playground/validator/v10"
)

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func BindErrorHandler(err error) []response.ApiError {
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]response.ApiError, len(ve))
			for i, fe := range ve {
				out[i] = response.ApiError{Field: fe.Field(), Message: msgForTag(fe.Tag())}
			}
			return out
		}
	}
	return nil
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "min": 
		return "This field must be at least 20"
    case "exists":
		return "This field is required"
	}
	return ""
}
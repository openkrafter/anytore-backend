package customerror

import (
	"fmt"
	"net/http"
)

type Error404 struct {
	ErrorCode int
	Body      map[string]string
}

func (e *Error404) Error() string {
	return fmt.Sprintf("404 Not Found: %s", e.Body["message"])
}

func NewError404() Error404 {
	return Error404{
		ErrorCode: http.StatusNotFound,
		Body: map[string]string{
			"status":  "error",
			"message": "Resource not found",
		},
	}
}

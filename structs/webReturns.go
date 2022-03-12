package structs

import (
	"net/http"
)

func QuickWebReturn(statusCode uint32, errorMessage string) (*WebReturn, error) {
	return &WebReturn{
		StatusCode:   statusCode,
		ErrorMessage: errorMessage,
	}, nil
}

func ForbiddenWebReturn() (*WebReturn, error) {
	return &WebReturn{
		StatusCode:   http.StatusForbidden,
		ErrorMessage: "user not authorised to view this page",
	}, nil
}

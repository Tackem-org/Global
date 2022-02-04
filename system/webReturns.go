package system

import (
	"net/http"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/structs"
)

func QuickWebReturn(statusCode uint32, errorMessage string) (*structs.WebReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[system.ForbiddenWebReturn() (*system.WebReturn, error)]")
	return &structs.WebReturn{
		StatusCode:   statusCode,
		ErrorMessage: errorMessage,
	}, nil
}

func ForbiddenWebReturn() (*structs.WebReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[system.ForbiddenWebReturn() (*system.WebReturn, error)]")
	return &structs.WebReturn{
		StatusCode:   http.StatusForbidden,
		ErrorMessage: "user not authorised to view this page",
	}, nil
}

package structs

import (
	"net/http"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
)

func QuickWebReturn(statusCode uint32, errorMessage string) (*WebReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.structs.QuickWebReturn")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] statusCode=%d, errorMessage=%s", statusCode, errorMessage)
	return &WebReturn{
		StatusCode:   statusCode,
		ErrorMessage: errorMessage,
	}, nil
}

func ForbiddenWebReturn() (*WebReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.structs.ForbiddenWebReturn")
	return &WebReturn{
		StatusCode:   http.StatusForbidden,
		ErrorMessage: "user not authorised to view this page",
	}, nil
}

package system

import (
	"net/http"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
)

func ForbiddenWebReturn() (*WebReturn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[system.ForbiddenWebReturn() (*system.WebReturn, error)]")
	return &WebReturn{
		StatusCode:   http.StatusForbidden,
		ErrorMessage: "user not authorised to view this page",
	}, nil
}

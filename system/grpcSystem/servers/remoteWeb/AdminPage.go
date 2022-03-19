package remoteWeb

import (
	"context"
	"net/http"

	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/setupData"
	pb "github.com/Tackem-org/Proto/pb/remoteweb"
)

func (r *RemoteWebServer) AdminPage(ctx context.Context, in *pb.PageRequest) (*pb.PageResponse, error) {
	if _, err := helpers.GRPCAccessChecker(ctx, func(baseID string) helpers.ServiceKeyCheckInterface {
		return &masterData.ConnectionInfo
	}, "GRPC Add Dependent"); err != "" {
		return &pb.PageResponse{StatusCode: http.StatusInternalServerError, HideErrorFromUser: true, ErrorMessage: err}, nil
	}
	webRequest := MakeWebRequest(in)
	path := setupData.Data.GetAdminPath(webRequest.BasePath)
	if path == nil {
		logging.Warning("[GRPC Remote Web System Admin Page Request] %s: Not found", in.Path)
		return &pb.PageResponse{
			StatusCode:   http.StatusNotFound,
			ErrorMessage: "Page Not Found (NIL)",
		}, nil
	}
	response, err := path.Call(webRequest)
	pageResponse := MakePageResponse(in, response, err)
	return pageResponse, nil
}

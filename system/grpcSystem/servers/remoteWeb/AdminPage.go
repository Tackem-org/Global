package remoteWeb

import (
	"context"
	"net/http"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system/setupData"
	pb "github.com/Tackem-org/Proto/pb/remoteweb"
)

func (r *RemoteWebServer) AdminPage(ctx context.Context, in *pb.PageRequest) (*pb.PageResponse, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.servers.remoteWeb.RemoteWebServer{}.AdminPage")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] ctx in=%+v", in)
	webRequest := makeWebRequest(in)
	path := setupData.Data.GetAdminPath(webRequest.BasePath)
	if path == nil {
		logging.Warning("[GRPC Remote Web System Admin Page Request] %s: Not found", in.GetPath())
		return &pb.PageResponse{
			StatusCode:   http.StatusNotFound,
			ErrorMessage: "Page Not Found",
		}, nil
	}
	response, err := path.Call(webRequest)
	pageResponse := makePageResponse(in, response, err)
	return pageResponse, nil
}

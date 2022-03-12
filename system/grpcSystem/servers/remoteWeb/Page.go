package remoteWeb

import (
	"context"
	"net/http"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/system/setupData"
	pb "github.com/Tackem-org/Proto/pb/remoteweb"
)

func (r *RemoteWebServer) Page(ctx context.Context, in *pb.PageRequest) (*pb.PageResponse, error) {
	webRequest := makeWebRequest(in)
	path := setupData.Data.GetPath(webRequest.BasePath)
	if path == nil {
		logging.Warning("[GRPC Remote Web System Page Request] %s: Not found", in.GetPath())
		return &pb.PageResponse{
			StatusCode:   http.StatusNotFound,
			ErrorMessage: "Page Not Found",
		}, nil
	}
	response, err := path.Call(webRequest)
	pageResponse := makePageResponse(in, response, err)
	return pageResponse, nil
}

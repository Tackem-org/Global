package remoteWeb

import (
	"context"
	iofs "io/fs"
	"net/http"

	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/setupData"
	pb "github.com/Tackem-org/Proto/pb/remoteweb"
)

func (r *RemoteWebServer) File(ctx context.Context, in *pb.FileRequest) (*pb.FileResponse, error) {
	if _, err := helpers.GRPCAccessChecker(ctx, func(baseID string) helpers.ServiceKeyCheckInterface {
		return &masterData.ConnectionInfo
	}, "GRPC Add Dependent"); err != "" {
		return &pb.FileResponse{StatusCode: http.StatusInternalServerError, ErrorMessage: err}, nil
	}
	data, err := setupData.Data.StaticFS.ReadFile(in.Path)
	if err != nil {
		sc := http.StatusInternalServerError
		em := "Internal Error"
		switch err.(type) {
		case *iofs.PathError:
			if in.Path[len(in.Path)-1:] == "/" {
				sc = http.StatusForbidden
				em = "Path is a Directory Access Forbidden"
			} else {
				sc = http.StatusNotFound
				em = "File Not Found"
			}
		default:
			logging.Error("[GRPC Remote Web System File Request] %s:%s", in.Path, err.Error())
		}

		return &pb.FileResponse{
			StatusCode:   uint32(sc),
			ErrorMessage: em,
		}, nil
	}
	return &pb.FileResponse{
		StatusCode:   http.StatusOK,
		ErrorMessage: "",
		File:         data,
	}, nil
}

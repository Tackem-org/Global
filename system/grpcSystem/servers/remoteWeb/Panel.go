package remoteWeb

import (
	"context"
	"net/http"

	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/logging"
	pb "github.com/Tackem-org/Global/pb/remoteweb"
	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/setupData"
)

func (r *RemoteWebServer) Panel(ctx context.Context, in *pb.PanelRequest) (*pb.PanelResponse, error) {
	if _, err := helpers.GRPCAccessChecker(ctx, func(baseID string) helpers.ServiceKeyCheckInterface {
		return &masterData.ConnectionInfo
	}, "GRPC Add Dependent"); err != "" {
		return &pb.PanelResponse{StatusCode: http.StatusInternalServerError, HideErrorFromUser: true, ErrorMessage: err}, nil
	}

	Panel := setupData.Data.GetPanel(in.Name)

	if Panel == nil {
		logging.Warning("[GRPC Remote Web System Pop Up Panel Request] %s: Not found", in.Name)
		return &pb.PanelResponse{
			StatusCode:   http.StatusNotFound,
			ErrorMessage: "page not found",
		}, nil
	}
	response, err := Panel.HTMLCall(&structs.PanelRequest{
		Name:      in.Name,
		User:      structs.GetUserData(in.User),
		Variables: in.Variables,
	})

	pageResponse := MakePanelResponse(in, response, err)
	return pageResponse, nil
}

func MakePanelResponse(in *pb.PanelRequest, PanelReturn *structs.PanelReturn, err error) *pb.PanelResponse {
	if err != nil {
		logging.Error("[GRPC Remote Web System Pop Up Panel Request] %s:%s", in.GetName(), err.Error())
		return &pb.PanelResponse{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: "error with the system",
		}
	}

	return &pb.PanelResponse{
		StatusCode:        PanelReturn.StatusCode,
		HideErrorFromUser: false,
		ErrorMessage:      PanelReturn.ErrorMessage,
		Html:              PanelReturn.HTML,
	}
}

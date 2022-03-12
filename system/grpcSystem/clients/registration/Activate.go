package registration

import (
	"github.com/Tackem-org/Global/system/grpcSystem/connections"

	pb "github.com/Tackem-org/Proto/pb/registration"
	"google.golang.org/grpc"
)

func Activate(request *pb.ActivateRequest) (*pb.ActivateResponse, error) {
	conn, err := connections.Master()
	if err != nil {
		return &pb.ActivateResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, err
	}
	defer conn.Close()
	client := pb.NewRegistrationClient(conn)
	header, ctx, cancel := connections.MasterHeader()
	defer cancel()
	return client.Activate(ctx, request, grpc.Header(&header))
}

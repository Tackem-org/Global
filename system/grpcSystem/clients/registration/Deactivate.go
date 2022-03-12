package registration

import (
	"github.com/Tackem-org/Global/system/grpcSystem/connections"

	pb "github.com/Tackem-org/Proto/pb/registration"
	"google.golang.org/grpc"
)

func Deactivate(request *pb.DeactivateRequest) (*pb.DeactivateResponse, error) {
	conn, err := connections.Master()
	if err != nil {
		return &pb.DeactivateResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, err
	}
	defer conn.Close()
	client := pb.NewRegistrationClient(conn)
	header, ctx, cancel := connections.MasterHeader()
	defer cancel()
	return client.Deactivate(ctx, request, grpc.Header(&header))
}

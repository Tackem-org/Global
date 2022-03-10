package registration

import (
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system/grpcSystem/connections"

	pb "github.com/Tackem-org/Proto/pb/registration"
	"google.golang.org/grpc"
)

func Register(request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.client.registration.Register")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] request=%+v", request)
	conn, err := connections.MasterForce()
	if err != nil {
		return &pb.RegisterResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, err
	}
	defer conn.Close()
	client := pb.NewRegistrationClient(conn)
	header, ctx, cancel := connections.RegistrationHeader()
	defer cancel()
	return client.Register(ctx, request, grpc.Header(&header))
}

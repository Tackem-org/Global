package registration

import (
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system/grpcSystem/connections"

	pb "github.com/Tackem-org/Proto/pb/registration"
	"google.golang.org/grpc"
)

func Deregister(request *pb.DeregisterRequest) (*pb.DeregisterResponse, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.client.registration.Deregister")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] request=%+v", request)
	conn, err := connections.Master()
	if err != nil {
		return &pb.DeregisterResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, err
	}
	defer conn.Close()
	client := pb.NewRegistrationClient(conn)
	header, ctx, cancel := connections.MasterHeader()
	defer cancel()
	return client.Deregister(ctx, request, grpc.Header(&header))
}

package registration

import (
	pb "github.com/Tackem-org/Global/pb/registration"
	"github.com/Tackem-org/Global/system/grpcSystem/connections"
	"github.com/Tackem-org/Global/system/grpcSystem/headers"
	"google.golang.org/grpc"
)

type RegistrationClientInterface interface {
	Activate(request *pb.ActivateRequest) (*pb.ActivateResponse, error)
	Deactivate(request *pb.DeactivateRequest) (*pb.DeactivateResponse, error)
	Deregister(request *pb.DeregisterRequest) (*pb.DeregisterResponse, error)
	Disconnect(request *pb.DisconnectRequest) (*pb.DisconnectResponse, error)
	Register(request *pb.RegisterRequest) (*pb.RegisterResponse, error)
}

var I RegistrationClientInterface = &RegistrationClient{}

func Activate(request *pb.ActivateRequest) (*pb.ActivateResponse, error) {
	return I.Activate(request)
}

func Deactivate(request *pb.DeactivateRequest) (*pb.DeactivateResponse, error) {
	return I.Deactivate(request)
}

func Deregister(request *pb.DeregisterRequest) (*pb.DeregisterResponse, error) {
	return I.Deregister(request)
}

func Disconnect(request *pb.DisconnectRequest) (*pb.DisconnectResponse, error) {
	return I.Disconnect(request)
}

func Register(request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return I.Register(request)
}

type RegistrationClient struct{}

func (rc *RegistrationClient) Activate(request *pb.ActivateRequest) (*pb.ActivateResponse, error) {
	conn, err := connections.Master()
	if err != nil {
		return &pb.ActivateResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, err
	}
	defer conn.Close()
	client := pb.NewRegistrationClient(conn)
	header, ctx, cancel := headers.MasterHeader()
	defer cancel()
	return client.Activate(ctx, request, grpc.Header(&header))
}

func (rc *RegistrationClient) Deactivate(request *pb.DeactivateRequest) (*pb.DeactivateResponse, error) {
	conn, err := connections.Master()
	if err != nil {
		return &pb.DeactivateResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, err
	}
	defer conn.Close()
	client := pb.NewRegistrationClient(conn)
	header, ctx, cancel := headers.MasterHeader()
	defer cancel()
	return client.Deactivate(ctx, request, grpc.Header(&header))
}

func (rc *RegistrationClient) Deregister(request *pb.DeregisterRequest) (*pb.DeregisterResponse, error) {
	conn, err := connections.Master()
	if err != nil {
		return &pb.DeregisterResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, err
	}
	defer conn.Close()
	client := pb.NewRegistrationClient(conn)
	header, ctx, cancel := headers.MasterHeader()
	defer cancel()
	return client.Deregister(ctx, request, grpc.Header(&header))
}

func (rc *RegistrationClient) Disconnect(request *pb.DisconnectRequest) (*pb.DisconnectResponse, error) {
	conn, err := connections.Master()
	if err != nil {
		return &pb.DisconnectResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, err
	}
	defer conn.Close()
	client := pb.NewRegistrationClient(conn)
	header, ctx, cancel := headers.MasterHeader()
	defer cancel()
	return client.Disconnect(ctx, request, grpc.Header(&header))
}

func (rc *RegistrationClient) Register(request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	conn, err := connections.MasterForce()
	if err != nil {
		return &pb.RegisterResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, err
	}
	defer conn.Close()
	client := pb.NewRegistrationClient(conn)
	header, ctx, cancel := headers.RegistrationHeader()
	defer cancel()
	return client.Register(ctx, request, grpc.Header(&header))
}

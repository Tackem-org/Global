package registration

import pb "github.com/Tackem-org/Global/pb/registration"

type RegistrationClientInterface interface {
	Activate(request *pb.ActivateRequest) (*pb.ActivateResponse, error)
	Deactivate(request *pb.DeactivateRequest) (*pb.DeactivateResponse, error)
	Deregister(request *pb.DeregisterRequest) (*pb.DeregisterResponse, error)
	Disconnect(request *pb.DisconnectRequest) (*pb.DisconnectResponse, error)
	Register(request *pb.RegisterRequest) (*pb.RegisterResponse, error)
}

var I RegistrationClientInterface

func Activate(request *pb.ActivateRequest) (*pb.ActivateResponse, error) {
	if I == nil {
		I = &RegistrationClient{}
	}
	return I.Activate(request)
}

func Deactivate(request *pb.DeactivateRequest) (*pb.DeactivateResponse, error) {
	if I == nil {
		I = &RegistrationClient{}
	}
	return I.Deactivate(request)
}

func Deregister(request *pb.DeregisterRequest) (*pb.DeregisterResponse, error) {
	if I == nil {
		I = &RegistrationClient{}
	}
	return I.Deregister(request)
}

func Disconnect(request *pb.DisconnectRequest) (*pb.DisconnectResponse, error) {
	if I == nil {
		I = &RegistrationClient{}
	}
	return I.Disconnect(request)
}

func Register(request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if I == nil {
		I = &RegistrationClient{}
	}
	return I.Register(request)
}

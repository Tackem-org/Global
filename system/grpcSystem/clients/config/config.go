package config

import pb "github.com/Tackem-org/Proto/pb/config"

type ConfigClientInterface interface {
	Get(request *pb.GetConfigRequest) (*pb.GetConfigResponse, error)
	Set(request *pb.SetConfigRequest) (*pb.SetConfigResponse, error)
}

var (
	I ConfigClientInterface
)

func Get(request *pb.GetConfigRequest) (*pb.GetConfigResponse, error) {
	if I == nil {
		I = &ConfigClient{}
	}
	return I.Get(request)
}

func Set(request *pb.SetConfigRequest) (*pb.SetConfigResponse, error) {
	if I == nil {
		I = &ConfigClient{}
	}
	return I.Set(request)
}

package config

import pb "github.com/Tackem-org/Global/pb/config"

type ConfigClientInterface interface {
	Get(request *pb.GetConfigRequest) (*pb.GetConfigResponse, error)
	Set(request *pb.SetConfigRequest) (*pb.SetConfigResponse, error)
}

var I ConfigClientInterface = &ConfigClient{}

func Get(request *pb.GetConfigRequest) (*pb.GetConfigResponse, error) {
	return I.Get(request)
}

func Set(request *pb.SetConfigRequest) (*pb.SetConfigResponse, error) {
	return I.Set(request)
}

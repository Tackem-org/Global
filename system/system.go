package system

import (
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func GRPCServer() *grpc.Server {
	if grpcServer == nil {
		grpcServer = grpc.NewServer()
	}
	return grpcServer
}

func GetMasterConnection(force bool) (*grpc.ClientConn, error) {
	if !force {
		MUp.Wait()
	}
	url := masterUrl + ":" + masterPort
	return grpc.Dial(url, grpc.WithInsecure(), grpc.WithBlock())
}

func GetHeader() metadata.MD {
	return metadata.New(map[string]string{
		"baseID": regData.baseID,
		"key":    regData.key,
	})

}

func GetFirstHeader() metadata.MD {
	var key string
	if val, present := os.LookupEnv("MASTERKEY"); present {
		key = val
	}

	return metadata.New(map[string]string{
		"masterkey": key,
	})
}

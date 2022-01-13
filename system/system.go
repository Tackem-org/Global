package system

import (
	"context"
	"os"
	"time"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func GRPCServer() *grpc.Server {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[system.GRPCServer() *grpc.Server]")
	if grpcServer == nil {
		grpcServer = grpc.NewServer()
	}
	return grpcServer
}

func GetMasterConnection(force bool) (*grpc.ClientConn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[system.GetMasterConnection(force bool) (*grpc.ClientConn, error)]")
	if !force {
		MUp.Wait()
	}
	url := masterUrl + ":" + masterPort
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return grpc.DialContext(ctx, url, grpc.WithInsecure(), grpc.WithBlock())
}

func GetHeader() metadata.MD {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[system.GetHeader() metadata.MD]")
	return metadata.New(map[string]string{
		"baseID": regData.baseID,
		"key":    regData.key,
	})

}

func GetFirstHeader() metadata.MD {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[system.GetFirstHeader() metadata.MD]")
	var key string
	if val, present := os.LookupEnv("MASTERKEY"); present {
		key = val
	}

	return metadata.New(map[string]string{
		"masterkey": key,
	})
}

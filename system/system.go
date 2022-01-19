package system

import (
	"context"
	"fmt"
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

func getConnection(url string) (*grpc.ClientConn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[system.getConnection(url string) (*grpc.ClientConn, error)]")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return grpc.DialContext(ctx, url, grpc.WithInsecure(), grpc.WithBlock())
}

func GetMasterConnection(force bool) (*grpc.ClientConn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[system.GetMasterConnection(force bool) (*grpc.ClientConn, error)]")
	if !force {
		MUp.Wait()
	}
	return getConnection(masterUrl + ":" + masterPort)
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
	if val, ok := os.LookupEnv("REGKEY"); ok {
		key = val
	}

	return metadata.New(map[string]string{
		"registrationkey": key,
	})
}

func getRequiredService(serviceName string, serviceType string) *RequiredService {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[system.getRequiredService(serviceName string, serviceType string) *RequiredService]")
	for _, s := range requiredServices {
		if s.ServiceName == serviceName && s.ServiceType == serviceType {
			return &s
		}
	}
	return nil
}

func GetRequiredHeader(serviceName string, serviceType string) metadata.MD {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[system.GetRequiredHeader(serviceName string, serviceType string) metadata.MD]")
	r := getRequiredService(serviceName, serviceType)
	if r == nil {
		return metadata.New(map[string]string{})
	}
	return metadata.New(map[string]string{
		"baseID": regData.baseID,
		"key":    r.Key,
	})
}

func GetConnection(serviceName string, serviceType string) (*grpc.ClientConn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[system.GetConnection(serviceName string, serviceType string) (*grpc.ClientConn, error)]")
	MUp.Wait()
	r := getRequiredService(serviceName, serviceType)
	if r == nil {
		return nil, &ServiceDownError{}
	}
	return getConnection(fmt.Sprintf("%s:%d", r.Hostname, r.Hostport))

}

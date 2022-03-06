package connections

import (
	"context"
	"fmt"
	"time"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/requiredServices"
	"github.com/Tackem-org/Global/system/setupData"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func getConnection(url string) (*grpc.ClientConn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.connections.getConnection")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] url=%s", url)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return grpc.DialContext(ctx, url, grpc.WithInsecure(), grpc.WithBlock())
}

func MasterForce() (*grpc.ClientConn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.connections.MasterForce")
	return getConnection(fmt.Sprintf("%s:%d", masterData.URL, masterData.Port))
}

func Master() (*grpc.ClientConn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.connections.Master")
	if masterData.UP.Wait(5) {
		return MasterForce()
	}
	return nil, nil
}

func MasterHeader() (metadata.MD, context.Context, context.CancelFunc) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.connections.MasterHeader")
	header := metadata.New(map[string]string{
		"baseID": setupData.BaseID,
		"key":    setupData.Key,
	})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	ctx = metadata.NewOutgoingContext(ctx, header)
	return header, ctx, cancel
}

func RegistrationHeader() (metadata.MD, context.Context, context.CancelFunc) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.connections.RegistrationHeader")
	header := metadata.New(map[string]string{
		"registrationkey": masterData.RegistrationKey,
	})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	ctx = metadata.NewOutgoingContext(ctx, header)
	return header, ctx, cancel
}

func ServiceHeader(requiredService *requiredServices.RequiredService) metadata.MD {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.connections.ServiceHeader")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] requiredService=%+v", requiredService)
	return metadata.New(map[string]string{
		"baseID": setupData.BaseID,
		"key":    requiredService.Key,
	})
}

func ServiceConnection(requiredService *requiredServices.RequiredService) (*grpc.ClientConn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.connections.ServiceConnection")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] requiredService=%+v", requiredService)
	return getConnection(fmt.Sprintf("%s:%d", requiredService.IPAddress, requiredService.Port))
}

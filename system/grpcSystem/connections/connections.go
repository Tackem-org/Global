package connections

import (
	"context"
	"fmt"
	"time"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/sysErrors"
	"github.com/Tackem-org/Global/system/dependentServices"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/requiredServices"
	"github.com/Tackem-org/Global/system/setupData"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func getConnection(url string) (*grpc.ClientConn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.connections.getConnection")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] url=%s", url)
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

func RequiredServiceHeader(requiredService *requiredServices.RequiredService) metadata.MD {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.connections.RequiredServiceHeader")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] requiredService=%+v", requiredService)
	return metadata.New(map[string]string{
		"baseID": setupData.BaseID,
		"key":    requiredService.Key,
	})
}

func RequiredServiceConnection(requiredService *requiredServices.RequiredService) (*grpc.ClientConn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.connections.RequiredServiceConnection")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] requiredService=%+v", requiredService)
	if requiredService.UP.Check() {
		return getConnection(fmt.Sprintf("%s:%d", requiredService.IPAddress, requiredService.Port))
	}
	return nil, &sysErrors.ServiceDownError{}
}

func DependentServiceHeader(dependentService *dependentServices.DependentService) metadata.MD {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.connections.DependentServiceHeader")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] dependentService=%+v", dependentService)
	return metadata.New(map[string]string{
		"baseID": setupData.BaseID,
		"key":    dependentService.Key,
	})
}

func DependentServiceConnection(dependentService *dependentServices.DependentService) (*grpc.ClientConn, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.grpcSystem.connections.DependentServiceConnection")
	logging.Debug(debug.FUNCTIONARGS, "[FUNCTIONARGS] dependentService=%+v", dependentService)
	if dependentService.UP.Check() {
		return getConnection(fmt.Sprintf("%s:%d", dependentService.IPAddress, dependentService.Port))
	}
	return nil, &sysErrors.ServiceDownError{}
}

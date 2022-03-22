package connections

import (
	"context"
	"time"

	"github.com/Tackem-org/Global/sysErrors"
	"github.com/Tackem-org/Global/system/dependentServices"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/requiredServices"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ExtraOptions []grpc.DialOption

func Get(address string) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	opts := append(ExtraOptions, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	return grpc.DialContext(ctx, address, opts...)
}

func MasterForce() (*grpc.ClientConn, error) {
	return Get(masterData.Info.Address())
}

func Master() (*grpc.ClientConn, error) {
	if masterData.UP.Wait(5) {
		return Get(masterData.Info.Address())
	}
	return nil, sysErrors.MasterDownError
}

func RequiredService(requiredService *requiredServices.RequiredService) (*grpc.ClientConn, error) {
	if requiredService.UP.Check() {
		return Get(requiredService.Address())
	}
	return nil, sysErrors.ServiceDownError
}

func DependentService(dependentService *dependentServices.DependentService) (*grpc.ClientConn, error) {
	if dependentService.UP.Check() {
		return Get(dependentService.Address())
	}
	return nil, sysErrors.ServiceDownError
}

package connections

import (
	"context"
	"fmt"
	"time"

	"github.com/Tackem-org/Global/sysErrors"
	"github.com/Tackem-org/Global/system/dependentServices"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/requiredServices"
	"github.com/Tackem-org/Global/system/setupData"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func getConnection(url string) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return grpc.DialContext(ctx, url, grpc.WithInsecure(), grpc.WithBlock())
}

func MasterForce() (*grpc.ClientConn, error) {
	return getConnection(fmt.Sprintf("%s:%d", masterData.URL, masterData.Port))
}

func Master() (*grpc.ClientConn, error) {
	if masterData.UP.Wait(5) {
		return MasterForce()
	}
	return nil, nil
}

func MasterHeader() (metadata.MD, context.Context, context.CancelFunc) {
	header := metadata.New(map[string]string{
		"baseID": setupData.BaseID,
		"key":    setupData.Key,
	})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	ctx = metadata.NewOutgoingContext(ctx, header)
	return header, ctx, cancel
}

func RegistrationHeader() (metadata.MD, context.Context, context.CancelFunc) {
	header := metadata.New(map[string]string{
		"registrationkey": masterData.RegistrationKey,
	})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	ctx = metadata.NewOutgoingContext(ctx, header)
	return header, ctx, cancel
}

func RequiredServiceHeader(requiredService *requiredServices.RequiredService) metadata.MD {
	return metadata.New(map[string]string{
		"baseID": setupData.BaseID,
		"key":    requiredService.Key,
	})
}

func RequiredServiceConnection(requiredService *requiredServices.RequiredService) (*grpc.ClientConn, error) {
	if requiredService.UP.Check() {
		return getConnection(fmt.Sprintf("%s:%d", requiredService.IPAddress, requiredService.Port))
	}
	return nil, &sysErrors.ServiceDownError{}
}

func DependentServiceHeader(dependentService *dependentServices.DependentService) metadata.MD {
	return metadata.New(map[string]string{
		"baseID": setupData.BaseID,
		"key":    dependentService.Key,
	})
}

func DependentServiceConnection(dependentService *dependentServices.DependentService) (*grpc.ClientConn, error) {
	if dependentService.UP.Check() {
		return getConnection(fmt.Sprintf("%s:%d", dependentService.IPAddress, dependentService.Port))
	}
	return nil, &sysErrors.ServiceDownError{}
}

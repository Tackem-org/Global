package connections

import (
	"fmt"

	"github.com/Tackem-org/Global/sysErrors"
	"github.com/Tackem-org/Global/system/dependentServices"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/requiredServices"
	"google.golang.org/grpc"
)

func MasterForce() (*grpc.ClientConn, error) {
	if I == nil {
		I = &Connection{}
	}
	return I.get(fmt.Sprintf("%s:%d", masterData.Info.URL, masterData.Info.Port))
}

func Master() (*grpc.ClientConn, error) {
	if masterData.UP.Wait(5) {
		return MasterForce()
	}
	return nil, nil
}

func RequiredServiceConnection(requiredService *requiredServices.RequiredService) (*grpc.ClientConn, error) {
	if requiredService.UP.Check() {
		return I.get(fmt.Sprintf("%s:%d", requiredService.IPAddress, requiredService.Port))
	}
	return nil, &sysErrors.ServiceDownError{}
}

func DependentServiceConnection(dependentService *dependentServices.DependentService) (*grpc.ClientConn, error) {
	if dependentService.UP.Check() {
		return I.get(fmt.Sprintf("%s:%d", dependentService.IPAddress, dependentService.Port))
	}
	return nil, &sysErrors.ServiceDownError{}
}

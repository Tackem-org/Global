package headers

import (
	"context"
	"time"

	"github.com/Tackem-org/Global/system/dependentServices"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/requiredServices"
	"github.com/Tackem-org/Global/system/setupData"
	"google.golang.org/grpc/metadata"
)

func Header(data map[string]string) (metadata.MD, context.Context, context.CancelFunc) {
	header := metadata.New(data)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	ctx = metadata.NewOutgoingContext(ctx, header)
	return header, ctx, cancel
}

func MasterHeader() (metadata.MD, context.Context, context.CancelFunc) {
	return Header(map[string]string{"baseID": setupData.BaseID, "key": masterData.ConnectionInfo.Key})
}

func RegistrationHeader() (metadata.MD, context.Context, context.CancelFunc) {
	return Header(map[string]string{"registrationkey": masterData.Info.RegistrationKey})
}

func RequiredServiceHeader(requiredService *requiredServices.RequiredService) (metadata.MD, context.Context, context.CancelFunc) {
	return Header(map[string]string{"baseID": setupData.BaseID, "key": requiredService.Key})
}

func DependentServiceHeader(dependentService *dependentServices.DependentService) (metadata.MD, context.Context, context.CancelFunc) {
	return Header(map[string]string{"baseID": setupData.BaseID, "key": dependentService.Key})
}

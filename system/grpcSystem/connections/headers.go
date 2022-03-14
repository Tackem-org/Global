package connections

import (
	"context"
	"time"

	"github.com/Tackem-org/Global/system/dependentServices"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/requiredServices"
	"github.com/Tackem-org/Global/system/setupData"
	"google.golang.org/grpc/metadata"
)

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
		"registrationkey": masterData.Info.RegistrationKey,
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

func DependentServiceHeader(dependentService *dependentServices.DependentService) metadata.MD {
	return metadata.New(map[string]string{
		"baseID": setupData.BaseID,
		"key":    dependentService.Key,
	})
}

package global

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Proto/pb/registration"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func RegisterSystem(url string, data *registration.RegisterRequest) (BaseID string, ServiceID uint32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := registration.NewRegistrationClient(conn)

	md := metadata.New(map[string]string{})
	ctx = metadata.NewOutgoingContext(ctx, md)

	var header metadata.MD

	response, err := client.Register(ctx, data, grpc.Header(&header))
	if err != nil {
		logging.Fatal(err)
	}
	if response.GetSuccess() {
		BaseID = response.GetBaseId()
		ServiceID = response.GetServiceId()
		return
	}
	logging.Fatal(errors.New(response.GetErrorMessage()))
	return
}

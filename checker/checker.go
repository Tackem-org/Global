package checker

import (
	"context"
	"time"

	pb "github.com/Tackem-org/Proto/pb/checker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type HealthCheckReturn struct {
	running bool
	healthy bool
	errors  *[]string
}

func HealthCheck(address string, port string) HealthCheckReturn {
	url := address + ":" + port
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return HealthCheckReturn{
			running: false,
			healthy: false,
			errors: &[]string{
				"Not Running",
			},
		}
	}
	defer conn.Close()

	client := pb.NewCheckerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, metadata.MD{})

	response, err := client.HealthCheck(ctx, &pb.Empty{}, grpc.Header(&metadata.MD{}))
	if err != nil {
		return HealthCheckReturn{
			running: true,
			healthy: response.GetHealthy(),
			errors: &[]string{
				"Issue With Health Check function on Server End",
				err.Error(),
			},
		}
	}
	return HealthCheckReturn{
		running: true,
		healthy: response.GetHealthy(),
		errors:  &[]string{},
	}
}

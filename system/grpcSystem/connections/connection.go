package connections

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

type ConnectionInterface interface {
	get(url string) (*grpc.ClientConn, error)
}

var I ConnectionInterface

type Connection struct{}

func (c *Connection) get(url string) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return grpc.DialContext(ctx, url, grpc.WithInsecure(), grpc.WithBlock())
}

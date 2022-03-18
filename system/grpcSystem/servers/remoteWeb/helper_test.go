package remoteWeb_test

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type MockAdder struct {
	ip string
}

func (ma *MockAdder) Network() string {
	return ""
}

func (ma *MockAdder) String() string {
	return ma.ip
}

func MakeTestHeader(baseID string, key string, ip string) context.Context {
	header := metadata.New(map[string]string{"baseID": baseID, "key": key})
	ctx := metadata.NewIncomingContext(context.Background(), header)
	return peer.NewContext(ctx, &peer.Peer{Addr: &MockAdder{ip: ip}})
}

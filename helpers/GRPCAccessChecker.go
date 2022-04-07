package helpers

import (
	"context"
	"reflect"
	"strings"

	"github.com/Tackem-org/Global/logging"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type ServiceKeyCheckInterface interface {
	CheckKey(key string) bool
	CheckIP(ipAddress string) bool
}

func GRPCAccessChecker(ctx context.Context, getByBaseID func(baseID string) ServiceKeyCheckInterface, systemLabel string) (ServiceKeyCheckInterface, string) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		logging.Custom("SECURITY", "[%s] Bad GRPC Access Missing Header Info", systemLabel)
		return nil, "missing header info"
	}
	values := md.Get("baseID")
	if len(values) == 0 {
		logging.Custom("SECURITY", "[%s] Bad GRPC Access Missing BaseID In Header", systemLabel)
		return nil, "missing baseID in header"
	}
	baseID := values[0]
	if baseID == "" {
		logging.Custom("SECURITY", "[%s] Bad GRPC Access BaseID In Header Blank", systemLabel)
		return nil, "baseID in header blank"
	}
	s := getByBaseID(baseID)
	if s == nil || reflect.ValueOf(s).IsNil() {
		logging.Custom("SECURITY", "[%s] Bad GRPC Access BaseID Not Found", systemLabel)
		return nil, "baseID not found"
	}

	values = md.Get("key")
	if len(values) == 0 {
		logging.Custom("SECURITY", "[%s] Bad GRPC Access Missing Key In Header", systemLabel)
		return nil, "missing key in header"
	}
	key := values[0]
	if key == "" {
		logging.Custom("SECURITY", "[%s] Bad GRPC Access Key In Header Blank", systemLabel)
		return nil, "key in header blank"
	}

	if !s.CheckKey(key) {
		logging.Custom("SECURITY", "[%s] Bad GRPC Access Wrong Key", systemLabel)
		return nil, "wrong key"
	}

	client, ok := peer.FromContext(ctx)
	if !ok {
		logging.Custom("SECURITY", "[%s] Bad GRPC Access No IP Address", systemLabel)
		return nil, "no IP"
	}
	ip := strings.Split(client.Addr.String(), ":")[0]

	if !s.CheckIP(ip) {
		logging.Custom("SECURITY", "[%s] Bad GRPC Access IP Address %s NOT CORRECT IP Address For Service it is claiming to be", systemLabel, ip)
		return nil, "ip address not valid"
	}

	return s, ""
}

package regClient

import (
	"context"

	"github.com/Tackem-org/Global/system/setupData"
	"google.golang.org/grpc/metadata"
)

func checkKey(ctx context.Context) (bool, string) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false, "error retrieving header"
	}
	baseIDvalues := md.Get("baseID")
	if len(baseIDvalues) != 1 {
		return false, "baseID not found in header"
	}
	baseID := baseIDvalues[0]
	if baseID == "" {
		return false, "base id is blank"
	}
	keyvalues := md.Get("key")
	if len(keyvalues) != 1 {
		return false, "key not found in header"
	}
	key := keyvalues[0]
	if key == "" {
		return false, "key is blank"
	}

	if setupData.BaseID == baseID && setupData.Key == key {

		return true, ""
	}
	return false, "values not correct"
}

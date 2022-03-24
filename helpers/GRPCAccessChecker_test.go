package helpers_test

import (
	"context"
	"testing"

	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/logging"
	"github.com/stretchr/testify/assert"
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

type MockDataStore struct {
	baseID string
	key    string
	ip     string
}

func (m *MockDataStore) CheckKey(key string) bool {
	return key == m.key
}
func (m *MockDataStore) CheckIP(ipAddress string) bool {
	return ipAddress == m.ip
}

func TestGRPCAccessCheckerNoMetadata(t *testing.T) {
	logging.I = &MockLogging{}
	mds := &MockDataStore{
		baseID: "test1",
		key:    "",
		ip:     "",
	}
	ctx := context.Background()
	s, err := helpers.GRPCAccessChecker(ctx, func(baseID string) helpers.ServiceKeyCheckInterface {
		if baseID == mds.baseID {
			return mds
		}
		return nil
	}, "TESTING")
	assert.Nil(t, s)
	assert.NotEmpty(t, err)
	assert.Equal(t, "missing header info", err)

}

func TestGRPCAccessCheckerMissingBaseID(t *testing.T) {
	mds := &MockDataStore{
		baseID: "Test2",
		key:    "",
		ip:     "",
	}
	header := metadata.New(map[string]string{"HEAD": "TEST"})

	ctx := metadata.NewIncomingContext(context.Background(), header)
	s, err := helpers.GRPCAccessChecker(ctx, func(baseID string) helpers.ServiceKeyCheckInterface {
		if baseID == mds.baseID {
			return mds
		}
		return nil
	}, "TESTING")
	assert.Nil(t, s)
	assert.NotEmpty(t, err)
	assert.Equal(t, "missing baseID in header", err)
}

func TestGRPCAccessCheckerBlankBaseID(t *testing.T) {
	mds := &MockDataStore{
		baseID: "Test3",
		key:    "",
		ip:     "",
	}
	header := metadata.New(map[string]string{"baseID": ""})

	ctx := metadata.NewIncomingContext(context.Background(), header)
	s, err := helpers.GRPCAccessChecker(ctx, func(baseID string) helpers.ServiceKeyCheckInterface {
		if baseID == mds.baseID {
			return mds
		}
		return nil
	}, "TESTING")
	assert.Nil(t, s)
	assert.NotEmpty(t, err)
	assert.Equal(t, "baseID in header blank", err)
}

func TestGRPCAccessCheckerBaseIDNotFound(t *testing.T) {
	mds := &MockDataStore{
		baseID: "Test4",
		key:    "",
		ip:     "",
	}
	header := metadata.New(map[string]string{"baseID": "Test3"})

	ctx := metadata.NewIncomingContext(context.Background(), header)
	s, err := helpers.GRPCAccessChecker(ctx, func(baseID string) helpers.ServiceKeyCheckInterface {
		if baseID == mds.baseID {
			return mds
		}
		return nil
	}, "TESTING")
	assert.Nil(t, s)
	assert.NotEmpty(t, err)
	assert.Equal(t, "baseID not found", err)
}

func TestGRPCAccessCheckerMissingKey(t *testing.T) {
	mds := &MockDataStore{
		baseID: "Test5",
		key:    "key5",
		ip:     "",
	}
	header := metadata.New(map[string]string{"baseID": mds.baseID})

	ctx := metadata.NewIncomingContext(context.Background(), header)
	s, err := helpers.GRPCAccessChecker(ctx, func(baseID string) helpers.ServiceKeyCheckInterface {
		if baseID == mds.baseID {
			return mds
		}
		return nil
	}, "TESTING")
	assert.Nil(t, s)
	assert.NotEmpty(t, err)
	assert.Equal(t, "missing key in header", err)
}

func TestGRPCAccessCheckerBlankKey(t *testing.T) {
	mds := &MockDataStore{
		baseID: "Test6",
		key:    "key6",
		ip:     "",
	}
	header := metadata.New(map[string]string{"baseID": mds.baseID, "key": ""})

	ctx := metadata.NewIncomingContext(context.Background(), header)
	s, err := helpers.GRPCAccessChecker(ctx, func(baseID string) helpers.ServiceKeyCheckInterface {
		if baseID == mds.baseID {
			return mds
		}
		return nil
	}, "TESTING")
	assert.Nil(t, s)
	assert.NotEmpty(t, err)
	assert.Equal(t, "key in header blank", err)
}

func TestGRPCAccessCheckerKeyNotFound(t *testing.T) {
	mds := &MockDataStore{
		baseID: "Test7",
		key:    "key7",
		ip:     "",
	}
	header := metadata.New(map[string]string{"baseID": mds.baseID, "key": "key6"})

	ctx := metadata.NewIncomingContext(context.Background(), header)
	s, err := helpers.GRPCAccessChecker(ctx, func(baseID string) helpers.ServiceKeyCheckInterface {
		if baseID == mds.baseID {
			return mds
		}
		return nil
	}, "TESTING")
	assert.Nil(t, s)
	assert.NotEmpty(t, err)
	assert.Equal(t, "wrong key", err)
}

func TestGRPCAccessCheckerNoIP(t *testing.T) {
	mds := &MockDataStore{
		baseID: "Test8",
		key:    "key8",
		ip:     "",
	}
	header := metadata.New(map[string]string{"baseID": mds.baseID, "key": mds.key})

	ctx := metadata.NewIncomingContext(context.Background(), header)
	s, err := helpers.GRPCAccessChecker(ctx, func(baseID string) helpers.ServiceKeyCheckInterface {
		if baseID == mds.baseID {
			return mds
		}
		return nil
	}, "TESTING")
	assert.Nil(t, s)
	assert.NotEmpty(t, err)
	assert.Equal(t, "no IP", err)
}

func TestGRPCAccessCheckerWrongIP(t *testing.T) {
	mds := &MockDataStore{
		baseID: "Test9",
		key:    "key9",
		ip:     "127.0.0.2",
	}
	header := metadata.New(map[string]string{"baseID": mds.baseID, "key": mds.key})

	ctx := metadata.NewIncomingContext(context.Background(), header)
	ctx = peer.NewContext(ctx, &peer.Peer{Addr: &MockAdder{ip: "127.0.0.1"}})
	s, err := helpers.GRPCAccessChecker(ctx, func(baseID string) helpers.ServiceKeyCheckInterface {
		if baseID == mds.baseID {
			return mds
		}
		return nil
	}, "TESTING")
	assert.Nil(t, s)
	assert.NotEmpty(t, err)
	assert.Equal(t, "ip address not valid", err)
}

func TestGRPCAccessPass(t *testing.T) {
	mds := &MockDataStore{
		baseID: "Test9",
		key:    "key9",
		ip:     "127.0.0.1",
	}
	header := metadata.New(map[string]string{"baseID": mds.baseID, "key": mds.key})

	ctx := metadata.NewIncomingContext(context.Background(), header)
	ctx = peer.NewContext(ctx, &peer.Peer{Addr: &MockAdder{ip: mds.ip}})
	s, err := helpers.GRPCAccessChecker(ctx, func(baseID string) helpers.ServiceKeyCheckInterface {
		if baseID == mds.baseID {
			return mds
		}
		return nil
	}, "TESTING")
	assert.NotNil(t, s)
	assert.Empty(t, err)

}

package connections_test

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	"github.com/Tackem-org/Global/sysErrors"
	"github.com/Tackem-org/Global/system/dependentServices"
	"github.com/Tackem-org/Global/system/grpcSystem/connections"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/requiredServices"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func startGRPCServer() (*grpc.Server, *bufconn.Listener) {
	bufferSize := 1024 * 1024
	listener := bufconn.Listen(bufferSize)
	srv := grpc.NewServer()
	go func() {
		if err := srv.Serve(listener); err != nil {
			log.Fatalf("failed to start grpc server: %v", err)
		}
	}()

	getBufDialer := func(listener *bufconn.Listener) func(context.Context, string) (net.Conn, error) {
		return func(ctx context.Context, url string) (net.Conn, error) {
			return listener.Dial()
		}
	}

	connections.ExtraOptions = append(connections.ExtraOptions, grpc.WithContextDialer(getBufDialer(listener)))
	return srv, listener
}

func TestGet(t *testing.T) {
	srv, listener := startGRPCServer()
	assert.NotNil(t, srv, "Test GRPC SERVER not running")
	assert.NotNil(t, listener, "Test GRPC SERVER Listner Not Running")
	defer func() { time.Sleep(10 * time.Millisecond) }()
	defer srv.Stop()
	conn, err := connections.Get("")
	assert.Nil(t, err, "An Error has happened with the Connection")
	assert.NotNil(t, conn, "Connection Not Connected")
	if conn != nil {
		conn.Close()
	}
}

func TestMasterForce(t *testing.T) {
	srv, listener := startGRPCServer()
	assert.NotNil(t, srv, "Test GRPC SERVER not running")
	assert.NotNil(t, listener, "Test GRPC SERVER Listner Not Running")
	defer func() { time.Sleep(10 * time.Millisecond) }()
	defer srv.Stop()
	conn, err := connections.MasterForce()
	assert.Nil(t, err, "An Error has happened with the Connection")
	assert.NotNil(t, conn, "Connection Not Connected")
	if conn != nil {
		conn.Close()
	}
}

func TestMaster(t *testing.T) {
	srv, listener := startGRPCServer()
	assert.NotNil(t, srv, "Test GRPC SERVER not running")
	assert.NotNil(t, listener, "Test GRPC SERVER Listner Not Running")
	defer func() { time.Sleep(10 * time.Millisecond) }()
	defer srv.Stop()

	//if Master Is Up
	conn1, err1 := connections.Master()
	assert.Nil(t, err1, "An Error has happened with the Connection")
	assert.NotNil(t, conn1, "Connection Not Connected")
	if conn1 != nil {
		conn1.Close()
	}

	//After Master Goes Down
	masterData.UP.Down()
	conn2, err2 := connections.Master()
	assert.NotNil(t, err2, "An Error has happened with the Connection")
	assert.ErrorIs(t, err2, sysErrors.MasterDownError)
	assert.Nil(t, conn2, "Connection Connected when should have failed")
}

func TestRequiredServiceConnection(t *testing.T) {
	srv, listener := startGRPCServer()
	assert.NotNil(t, srv, "Test GRPC SERVER not running")
	assert.NotNil(t, listener, "Test GRPC SERVER Listner Not Running")
	defer func() { time.Sleep(10 * time.Millisecond) }()
	defer srv.Stop()

	r := &requiredServices.RequiredService{}
	//if Required Service Is Up
	conn1, err1 := connections.RequiredService(r)
	assert.Nil(t, err1, "An Error has happened with the Connection")
	assert.NotNil(t, conn1, "Connection Not Connected")
	if conn1 != nil {
		conn1.Close()
	}

	//After Required Service Goes Down
	r.UP.Down()
	conn2, err2 := connections.RequiredService(r)
	assert.NotNil(t, err2, "An Error has happened with the Connection")
	assert.ErrorIs(t, err2, sysErrors.ServiceDownError)
	assert.Nil(t, conn2, "Connection Connected when should have failed")
}

func TestDependentServiceConnection(t *testing.T) {
	srv, listener := startGRPCServer()
	assert.NotNil(t, srv, "Test GRPC SERVER not running")
	assert.NotNil(t, listener, "Test GRPC SERVER Listner Not Running")
	defer func() { time.Sleep(10 * time.Millisecond) }()
	defer srv.Stop()

	d := &dependentServices.DependentService{}
	//if Dependent Service Is Up
	conn1, err1 := connections.DependentService(d)
	assert.Nil(t, err1, "An Error has happened with the Connection")
	assert.NotNil(t, conn1, "Connection Not Connected")
	if conn1 != nil {
		conn1.Close()
	}

	//After Dependent Service Goes Down
	d.UP.Down()
	conn2, err2 := connections.DependentService(d)
	assert.NotNil(t, err2, "An Error has happened with the Connection")
	assert.Equal(t, err2, sysErrors.ServiceDownError)
	assert.Nil(t, conn2, "Connection Connected when should have failed")
}

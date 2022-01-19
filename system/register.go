package system

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"regexp"
	"time"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"

	pb "github.com/Tackem-org/Proto/pb/registration"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type Register struct {
	baseID    string
	serviceID uint32
	key       string

	data pb.RegisterRequest
}

func RegData() *Register {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[system.RegData() *Register]")
	if regData == nil {
		regData = NewRegister()
	}
	return regData
}

func NewRegister() *Register {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[system.NewRegister() *Register]")
	if val, present := os.LookupEnv("MASTERURL"); present {
		masterUrl = val
	}
	if val, present := os.LookupEnv("MASTERPORT"); present {
		masterPort = val
	}
	return &Register{}
}

func (r *Register) GetBaseID() string {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[system.(r *Register) GetBaseID() string]")
	return r.baseID
}

func (r *Register) GetServiceID() uint32 {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[system.(r *Register) GetServiceID() uint32]")
	return r.serviceID
}

func (r *Register) GetKey() string {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[system.(r *Register) GetKey() string]")
	return r.key
}

func (r *Register) GetPort() uint32 {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[system.(r *Register) GetPort() uint32]")
	return r.data.Hostport
}

func (r *Register) GetServiceName() string {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[system.(r *Register) GetServiceName() string]")
	return r.data.ServiceName
}

func (r *Register) GetServiceType() string {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[system.(r *Register) GetServiceType() string]")
	return r.data.ServiceType
}

func (r *Register) Setup(baseData BaseData) {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[system.(r *Register) Setup(baseData BaseData)] {baseData=%v}", baseData)
	rawHostname, err := ioutil.ReadFile("/etc/hostname")
	if err != nil {
		logging.Fatal(err.Error())
	}

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		logging.Fatal(err.Error())
	}
	r.data = pb.RegisterRequest{
		ServiceName:      baseData.ServiceName,
		ServiceType:      baseData.ServiceType,
		Version:          &pb.Version{Major: uint32(baseData.Version.Major), Minor: uint32(baseData.Version.Minor), Hotfix: uint32(baseData.Version.Hotfix)},
		Hostname:         reg.ReplaceAllString(string(rawHostname), ""),
		Hostport:         freePort(),
		Multi:            baseData.Multi,
		SingleRun:        baseData.SingleRun,
		Webaccess:        baseData.WebAccess,
		NavItems:         baseData.NavItems,
		ConfigItems:      baseData.ConfigItems,
		RequiredServices: baseData.RequiredServices,
	}
}

func (r *Register) connection() (pb.RegistrationClient, *grpc.ClientConn, context.CancelFunc) {
	logging.Debug(debug.FUNCTIONCALLS|debug.GPRCCLIENT, "CALLED:[system.(r *Register) connection() pb.RegistrationClient]")
	url := masterUrl + ":" + masterPort
	connctx, conncancel := context.WithTimeout(context.Background(), time.Second)
	conn, err := grpc.DialContext(connctx, url, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		logging.Fatal(err.Error())
	}

	return pb.NewRegistrationClient(conn), conn, conncancel
}

func (r *Register) Connect() bool {
	logging.Debug(debug.FUNCTIONCALLS|debug.GPRCCLIENT, "CALLED:[system.(r *Register) Connect() bool]")
	client, conn, conncancel := r.connection()
	defer conncancel()
	defer conn.Close()
	header := GetFirstHeader()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, header)

	response, err := client.Register(ctx, &r.data, grpc.Header(&header))
	if err != nil {
		logging.Fatal(err.Error())
	}
	if response.GetSuccess() {
		r.baseID = response.GetBaseId()
		r.serviceID = response.GetServiceId()
		r.key = response.GetKey()
		return true
	}
	logging.Error(response.GetErrorMessage())
	return false
}

func (r *Register) Disconnect() {
	logging.Debug(debug.FUNCTIONCALLS|debug.GPRCCLIENT, "CALLED:[system.(r *Register) Disconnect()]")
	client, conn, conncancel := r.connection()
	defer conncancel()
	defer conn.Close()
	header := GetHeader()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, header)
	response, err := client.Disconnect(ctx, &pb.DisconnectRequest{
		BaseId: r.baseID,
	}, grpc.Header(&header))
	if err != nil {
		logging.Fatal(err.Error())
	}
	if !response.GetSuccess() {
		logging.Warningf("failed to disconnect service from master: %s", response.GetErrorMessage())
	}
}

func (r *Register) Activate() {
	logging.Debug(debug.FUNCTIONCALLS|debug.GPRCCLIENT, "CALLED:[system.(r *Register) Activate()]")
	client, conn, conncancel := r.connection()
	defer conncancel()
	defer conn.Close()
	header := GetHeader()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, header)
	response, err := client.Activate(ctx, &pb.ActivateRequest{
		BaseId: r.baseID,
	}, grpc.Header(&header))
	if err != nil {
		logging.Fatal(err.Error())
	}
	if !response.GetSuccess() {
		logging.Warningf("failed to disconnect service from master: %s", response.GetErrorMessage())
	}
}

func (r *Register) Deactivate() {
	logging.Debug(debug.FUNCTIONCALLS|debug.GPRCCLIENT, "CALLED:[system.(r *Register) Deactivate()]")
	client, conn, conncancel := r.connection()
	defer conncancel()
	defer conn.Close()
	header := GetHeader()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, header)
	response, err := client.Deactivate(ctx, &pb.DeactivateRequest{
		BaseId: r.baseID,
	}, grpc.Header(&header))
	if err != nil {
		logging.Fatal(err.Error())
	}
	if !response.GetSuccess() {
		logging.Warningf("failed to deactivate service with master: %s", response.GetErrorMessage())
	}
}

func freePort() (port uint32) {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[system.freePort() (port uint32) ]")
	bind := ""
	if val, ok := os.LookupEnv("BIND"); ok {
		bind = val
	}
	port = 50001
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", bind, port))
	for {
		if err == nil {
			break
		}
		port++
		ln, err = net.Listen("tcp", fmt.Sprintf("%s:%d", bind, port))
	}
	ln.Close()
	return
}

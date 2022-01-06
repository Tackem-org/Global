package system

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"regexp"
	"time"

	"github.com/Tackem-org/Global/logging"

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
	if regData == nil {
		regData = NewRegister()
	}
	return regData
}

func NewRegister() *Register {
	if val, present := os.LookupEnv("MASTERURL"); present {
		masterUrl = val
	}
	if val, present := os.LookupEnv("MASTERPORT"); present {
		masterPort = val
	}
	return &Register{}
}

func (r *Register) GetBaseID() string {
	return r.baseID
}

func (r *Register) GetServiceID() uint32 {
	return r.serviceID
}

func (r *Register) GetKey() string {
	return r.key
}

func (r *Register) GetPort() uint32 {
	return r.data.Hostport
}

func (r *Register) GetServiceName() string {
	return r.data.ServiceName
}

func (r *Register) GetServiceType() string {
	return r.data.ServiceType
}

func (r *Register) Setup(baseData BaseData) {

	rawHostname, err := ioutil.ReadFile("/etc/hostname")
	if err != nil {
		logging.Fatal(err)
	}

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		logging.Fatal(err)
	}
	r.data = pb.RegisterRequest{
		ServiceName: baseData.ServiceName,
		ServiceType: baseData.ServiceType,
		Version:     &pb.Version{Major: uint32(baseData.Version.Major), Minor: uint32(baseData.Version.Minor), Hotfix: uint32(baseData.Version.Hotfix), Suffix: baseData.Version.Suffix},
		Hostname:    reg.ReplaceAllString(string(rawHostname), ""),
		Hostport:    freePort(),
		Multi:       baseData.Multi,
		SingleRun:   baseData.SingleRun,
		Webaccess:   baseData.WebAccess,
		NavItems:    baseData.NavItems,
		ConfigItems: []*pb.ConfigItem{},
	}
}

func (r *Register) Connect() bool {

	url := masterUrl + ":" + masterPort
	connctx, conncancel := context.WithTimeout(context.Background(), time.Second)
	defer conncancel()
	conn, err := grpc.DialContext(connctx, url, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		logging.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewRegistrationClient(conn)

	header := GetFirstHeader()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, header)

	response, err := client.Register(ctx, &r.data, grpc.Header(&header))
	if err != nil {
		logging.Fatal(err)
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
	logging.Info("DISCONNECT CALLED")

	url := masterUrl + ":" + masterPort

	connctx, conncancel := context.WithTimeout(context.Background(), time.Second)
	defer conncancel()
	conn, err := grpc.DialContext(connctx, url, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		logging.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewRegistrationClient(conn)

	header := GetHeader()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, header)
	response, err := client.Disconnect(ctx, &pb.DisconnectRequest{
		BaseId: r.baseID,
	}, grpc.Header(&header))
	if err != nil {
		logging.Fatal(err)
	}
	if response.GetSuccess() {
		logging.Info("DISCONNECT FINISHED")
		return
	}
	logging.Fatal(errors.New("failed to disconnect system from master:" + response.GetErrorMessage()))
}

func freePort() (port uint32) {
	port = 50001
	ln, err := net.Listen("tcp", ":"+fmt.Sprint(port))
	for {
		if err == nil {
			break
		}
		port++
		ln, err = net.Listen("tcp", ":"+fmt.Sprint(port))
	}
	ln.Close()
	return
}

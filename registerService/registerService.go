package registerService

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
	"google.golang.org/grpc/metadata"
)

var (
	masterUrl  = "127.0.0.1" //"tackem_master"
	masterPort = "50001"
)

type BaseData struct {
	ServiceName string
	ServiceType string
	Multi       bool
	WebAccess   bool
	NavItems    []*pb.NavItem
}

type Register struct {
	baseID    string
	serviceID uint32

	data pb.RegisterRequest
}

func NewRegister() *Register {
	getMasterURL()
	return &Register{}
}

func (r *Register) GetBaseID() string {
	return r.baseID
}

func (r *Register) GetServiceID() uint32 {
	return r.serviceID
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
	hostname := reg.ReplaceAllString(string(rawHostname), "")
	port := FreePort()
	r.data = pb.RegisterRequest{
		ServiceName: baseData.ServiceName,
		ServiceType: baseData.ServiceType,
		Hostname:    hostname,
		Hostport:    port,
		Multi:       baseData.Multi,
		Webaccess:   baseData.WebAccess,
		NavItems:    baseData.NavItems,
	}
}

func (r *Register) Connect() bool {

	url := MakeMasterURL()
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logging.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewRegistrationClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, metadata.MD{})

	response, err := client.Register(ctx, &r.data)
	if err != nil {
		logging.Fatal(err)
	}
	if response.GetSuccess() {
		r.baseID = response.GetBaseId()
		r.serviceID = response.GetServiceId()
		return true
	}
	logging.Error(response.GetErrorMessage())
	return false
}

func (r *Register) Disconnect() {
	logging.Info("DISCONNECT CALLED")

	url := masterUrl + ":" + masterPort
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logging.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewRegistrationClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, metadata.MD{})

	response, err := client.Disconnect(ctx, &pb.IDRequest{BaseId: r.baseID}, grpc.Header(&metadata.MD{}))
	if err != nil {
		logging.Fatal(err)
	}
	if response.GetSuccess() {
		logging.Info("DISCONNECT FINISHED")
		return
	}
	logging.Fatal(errors.New("failed to disconnect system from master"))
}

func FreePort() (port uint32) {
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

func getMasterURL() {
	if val, present := os.LookupEnv("MASTERURL"); present {
		masterUrl = val
	}
	if val, present := os.LookupEnv("MASTERPORT"); present {
		masterPort = val
	}
}

func MakeMasterURL() string {
	return masterUrl + ":" + masterPort
}

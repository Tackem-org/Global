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
	Data       *Register
	masterUrl  = "127.0.0.1" //"tackem_master"
	masterPort = "50001"
)

type Register struct {
	baseID    string
	serviceID uint32

	data pb.RegisterRequest
}

func NewRegister() *Register {
	GetMasterURL()
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

func (r *Register) Setup(serviceName string, serviceType string, multi bool, webAccess bool, navItems []*pb.NavItem) {

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
		ServiceName: serviceName,
		ServiceType: serviceType,
		Hostname:    hostname,
		Hostport:    port,
		Multi:       multi,
		Webaccess:   webAccess,
		NavItems:    navItems,
	}
}

func (r *Register) Connect() bool {

	url := masterUrl + ":" + masterPort
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logging.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewRegistrationClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	md := metadata.New(map[string]string{})
	ctx = metadata.NewOutgoingContext(ctx, md)

	var header metadata.MD

	response, err := client.Register(ctx, &r.data, grpc.Header(&header))
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
	md := metadata.New(map[string]string{})
	ctx = metadata.NewOutgoingContext(ctx, md)

	var header metadata.MD

	response, err := client.Disconnect(ctx, &pb.IDRequest{BaseId: r.baseID}, grpc.Header(&header))
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

func GetMasterURL() {
	if val, present := os.LookupEnv("MASTERURL"); present {
		masterUrl = val
	}
	if val, present := os.LookupEnv("MASTERPORT"); present {
		masterPort = val
	}
}

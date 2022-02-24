package system

import (
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/structs"
	pb "github.com/Tackem-org/Proto/pb/registration"
	pbw "github.com/Tackem-org/Proto/pb/web"
	"google.golang.org/grpc"
)

type SetupData struct {
	BaseData    BaseData
	LogFile     string
	VerboseLog  bool
	DebugLevel  debug.Mask
	GRPCSystems func(server *grpc.Server)
	WebSystems  func()
	WebSockets  func()
	TaskGrabber func() []*pbw.TaskMessage
	MainSetup   func()
	MainSystem  func()
	Shutdown    func()
}

type BaseData struct {
	ServiceName      string
	ServiceType      string
	Version          structs.Version
	Multi            bool
	SingleRun        bool
	NavItems         []*pb.NavItem
	ConfigItems      []*pb.ConfigItem
	RequiredServices []*pb.RequiredService
}

type WebSocketRequest struct {
	Command string
	User    *structs.UserData
	Data    map[string]interface{}
}

type WebSocketReturn struct {
	StatusCode   uint32
	ErrorMessage string
	TellAll      bool
	Data         map[string]interface{}
}

type DependentService struct {
	BaseID    string
	Key       string
	IPAddress string
	Port      uint32
	SingleRun bool
}

type RequiredService struct {
	ServiceName string
	ServiceType string
	BaseID      string
	Key         string
	IPAddress   string
	Port        uint32
	SingleRun   bool
}

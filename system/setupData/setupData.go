package setupData

import (
	"embed"
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/Tackem-org/Global/structs"
	pb "github.com/Tackem-org/Proto/pb/registration"
	pbw "github.com/Tackem-org/Proto/pb/web"
	"google.golang.org/grpc"
)

var (
	mu        sync.RWMutex
	Data      *SetupData
	Active    bool = false
	BaseID    string
	ServiceID uint64
	Key       string
	Port      uint32 = 50001
)

type SetupData struct {
	mu               sync.RWMutex
	ServiceName      string
	ServiceType      string
	Version          structs.Version
	Multi            bool
	SingleRun        bool
	StartActive      bool
	NavItems         []*pb.NavItem
	ConfigItems      []*pb.ConfigItem
	RequiredServices []*pb.RequiredService
	Groups           []string
	Permissions      []string

	MasterConf  string
	LogFile     string
	VerboseLog  bool
	DebugLevel  debug.Mask
	GRPCSystems func(server *grpc.Server)

	StaticFS     *embed.FS
	AdminPaths   []*AdminPathItem
	Paths        []*PathItem
	Sockets      []*SocketItem
	TaskGrabber  func() []*pbw.TaskMessage
	MainSetup    func()
	MainSystem   func()
	MainShutdown func()
}

type PageFunc = func(in *structs.WebRequest) (*structs.WebReturn, error)
type SocketFunc = func(in *structs.SocketRequest) (*structs.SocketReturn, error)

type AdminPathItem struct {
	Path        string
	PostAllowed bool
	GetDisabled bool
	Call        PageFunc
}

type PathItem struct {
	Path        string
	Permission  string
	PostAllowed bool
	GetDisabled bool
	Call        PageFunc
}

type SocketItem struct {
	Command           string
	Permission        string
	AdminOnly         bool
	RequiredVariables []string
	Call              SocketFunc
}

func FreeTCPPort() (net.Listener, error) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.system.setupData.FreePort")
	mu.Lock()
	defer mu.Unlock()
	bind := ""
	if val, ok := os.LookupEnv("BIND"); ok {
		bind = val
	}
	for {
		ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", bind, Port))
		if err == nil {
			return ln, nil
		}
		Port++
	}
}

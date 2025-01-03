package setupData

import (
	"fmt"
	"io/fs"
	"net"
	"os"
	"sync"

	pb "github.com/Tackem-org/Global/pb/registration"
	pbw "github.com/Tackem-org/Global/pb/web"
	"github.com/Tackem-org/Global/structs"
	"google.golang.org/grpc"
)

type EmbedInterface interface {
	Open(name string) (fs.File, error)
	ReadFile(name string) ([]byte, error)
}

var (
	mu        sync.RWMutex
	Data      *SetupData
	Active    bool = false
	BaseID    string
	ServiceID uint64
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
	VerboseLog       bool
	GRPCSystems      func(server *grpc.Server)

	StaticFS            EmbedInterface
	AdminPaths          []*AdminPathItem
	Paths               []*PathItem
	Sockets             []*SocketItem
	Panels              []*PanelItem
	TaskGrabber         func() []*pbw.TaskMessage
	NotificationGrabber func() []*pbw.NotificationMessage
	MainSetup           func()
	MainSystem          func()
	MainShutdown        func()
}

type PageFunc = func(in *structs.WebRequest) (*structs.WebReturn, error)
type SocketFunc = func(in *structs.SocketRequest) (*structs.SocketReturn, error)
type PanelFunc = func(in *structs.PanelRequest) (*structs.PanelReturn, error)

type AdminPathItem struct {
	Path          string
	PostAllowed   bool
	GetDisabled   bool
	AllowedPanels []string
	Call          PageFunc
}

type PathItem struct {
	Path          string
	Permission    string
	PostAllowed   bool
	GetDisabled   bool
	AllowedPanels []string
	Call          PageFunc
}

type SocketItem struct {
	Command           string
	Permission        string
	AdminOnly         bool
	RequiredVariables []string
	Call              SocketFunc
}

type RequiredVariable struct {
	Name    string
	Options []string
}

type HorizontalAlignType int8
type VerticalAlignType int8

const (
	HCenter HorizontalAlignType = iota
	HLeft
	HRight
)

const (
	VCenter VerticalAlignType = iota
	VTop
	VBottom
)

type PanelLayout struct {
	HorizontalAlign HorizontalAlignType
	VerticalAlign   VerticalAlignType
	Width           uint32
	Height          uint32
	ScrollWidth     bool
	ScrollHeight    bool
	TitleBar        bool
	Minimise        bool
	Close           bool
}
type PanelItem struct {
	Name              string
	Label             string
	Description       string
	Layout            PanelLayout
	AdminOnly         bool
	Permission        string
	AllowedPaths      []string
	RequiredVariables []RequiredVariable

	HTMLCall   PanelFunc
	SocketCall SocketFunc
}

func FreeTCPPort() net.Listener {
	mu.Lock()
	defer mu.Unlock()
	bind := ""
	if val, ok := os.LookupEnv("BIND"); ok {
		bind = val
	}
	for {
		if ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", bind, Port)); err == nil {
			return ln
		}
		Port++
	}
}

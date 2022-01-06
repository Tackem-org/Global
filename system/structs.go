package system

import (
	pb "github.com/Tackem-org/Proto/pb/registration"
	"google.golang.org/grpc"
)

type SetupData struct {
	BaseData    BaseData
	LogFile     string
	VerboseLog  bool
	GPRCSystems func(server *grpc.Server)
	WebSystems  func()
	MainSystem  func()
	Shutdown    func()
}

type BaseData struct {
	ServiceName string
	ServiceType string
	Version     Versionstruct
	Multi       bool
	SingleRun   bool
	WebAccess   bool
	NavItems    []*pb.NavItem
}

type WebRequest struct {
	FullPath      string
	CleanPath     string
	UserID        uint64
	SessionToken  string
	Method        string
	QueryParams   map[string]interface{}
	Post          map[string]interface{}
	PathVariables map[string]interface{}
}

type WebReturn struct {
	FilePath       string
	PageString     string
	PageData       map[string]interface{}
	CustomPageName string
	CustomCss      []string
	CustomJs       []string
}

type Versionstruct struct {
	Major  uint8
	Minor  uint8
	Hotfix uint8
	Suffix string
}

package system

import (
	"embed"
	"sync"

	"github.com/Tackem-org/Global/helpers"
	"google.golang.org/grpc"
)

var (
	Data               SetupData
	Active             bool = false
	regData            *Register
	grpcServer         *grpc.Server
	WG                 *sync.WaitGroup
	MUp                helpers.Locker
	masterUrl          string = "127.0.0.1" //"tackem_master"
	masterPort         string = "50001"
	pagesData          map[string]func(in *WebRequest) (*WebReturn, error)
	adminPagesData     map[string]func(in *WebRequest) (*WebReturn, error)
	webSocketData      map[string]func(in *WebSocketRequest) (*WebSocketReturn, error)
	fileSystem         *embed.FS
	healthcheckHealthy bool
	healthcheckIssues  []string
	dependentServices  []DependentService
	requiredServices   []RequiredService
)

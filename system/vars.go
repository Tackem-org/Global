package system

import (
	"embed"
	"sync"

	"github.com/Tackem-org/Global/helpers"
	"google.golang.org/grpc"
)

var (
	Data               SetupData
	regData            *Register
	grpcServer         *grpc.Server
	WG                 *sync.WaitGroup
	MUp                helpers.Locker
	masterUrl          = "127.0.0.1" //"tackem_master"
	masterPort         = "50001"
	pagesData          map[string]func(in *WebRequest) (*WebReturn, error)
	adminPagesData     map[string]func(in *WebRequest) (*WebReturn, error)
	fileSystem         *embed.FS
	healthcheckHealthy bool
	healthcheckIssues  []string

	ShutdownCommand chan bool
)

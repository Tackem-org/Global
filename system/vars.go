package system

import (
	"embed"
	"sync"

	"google.golang.org/grpc"
)

var (
	Data           SetupData
	regData        *Register
	grpcServer     *grpc.Server
	WG             *sync.WaitGroup
	MasterUpLock   sync.Mutex
	masterUrl      = "127.0.0.1" //"tackem_master"
	masterPort     = "50001"
	pagesData      map[string]func(in *WebRequest) (*WebReturn, error)
	adminPagesData map[string]func(in *WebRequest) (*WebReturn, error)
	fileSystem     *embed.FS
)

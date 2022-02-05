package system

import (
	"embed"
	"sync"

	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/structs"
	"google.golang.org/grpc"

	pb "github.com/Tackem-org/Proto/pb/registration"
)

var (
	Data                SetupData
	Active              bool = false
	regData             *Register
	grpcServer          *grpc.Server
	WG                  *sync.WaitGroup
	MUp                 helpers.Locker
	masterUrl           string = "127.0.0.1" //"tackem_master"
	masterPort          string = "50001"
	pagesData           map[string]func(in *structs.WebRequest) (*structs.WebReturn, error)
	pagesProtoData      []*pb.WebLinkItem
	adminPagesData      map[string]func(in *structs.WebRequest) (*structs.WebReturn, error)
	adminPagesProtoData []*pb.AdminWebLinkItem
	webSocketData       map[string]func(in *WebSocketRequest) (*WebSocketReturn, error)
	adminWebSocketData  map[string]func(in *WebSocketRequest) (*WebSocketReturn, error)
	fileSystem          *embed.FS
	healthcheckHealthy  bool
	healthcheckIssues   []string
	dependentServices   []DependentService
	requiredServices    []RequiredService
)

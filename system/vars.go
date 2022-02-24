package system

import (
	"embed"
	"sync"

	"github.com/Tackem-org/Global/helpers"
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
	masterPort          string = "50000"
	pagesData           map[string]PageFunc
	pagesProtoData      []*pb.WebLinkItem
	adminPagesData      map[string]PageFunc
	adminPagesProtoData []*pb.AdminWebLinkItem
	webSocketData       map[string]SocketFunc
	webSocketProtoData  []*pb.WebSocketItem
	fileSystem          *embed.FS
	healthcheckHealthy  bool
	healthcheckIssues   []string
	dependentServices   []DependentService
	requiredServices    []RequiredService
)

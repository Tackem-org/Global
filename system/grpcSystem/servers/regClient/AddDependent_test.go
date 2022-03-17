package regClient_test

import (
	"testing"

	"github.com/Tackem-org/Global/system/grpcSystem/servers/regClient"
	"github.com/Tackem-org/Global/system/masterData"
	pb "github.com/Tackem-org/Proto/pb/regclient"
	"github.com/stretchr/testify/assert"
)

//TODO working from here you need to add in the code bellow to all other server connections and then test all of these.
//Should be as easy as test for access to fail first, then make it pass adding a dependent, then try and add the same
//again to get the 2nd failure to happen
//
//copyish this for the thers
func TestAddDependent(t *testing.T) {
	s := regClient.RegClientServer{}
	ctx1 := MakeTestHeader("", "", "")
	r1, err1 := s.AddDependent(ctx1, &pb.AddDependentRequest{})
	assert.NotNil(t, r1)
	assert.Nil(t, err1)
	assert.False(t, r1.Success)

	masterData.ConnectionInfo = masterData.ConnectionInfostruct{
		Key: "key1",
		IP:  "127.0.0.1",
	}
	ctx2 := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	r2, err2 := s.AddDependent(ctx2, &pb.AddDependentRequest{})
	assert.NotNil(t, r2)
	assert.Nil(t, err2)
	assert.True(t, r2.Success)

	ctx3 := MakeTestHeader("Test1", masterData.ConnectionInfo.Key, masterData.ConnectionInfo.IP)
	r3, err3 := s.AddDependent(ctx3, &pb.AddDependentRequest{})
	assert.NotNil(t, r3)
	assert.Nil(t, err3)
	assert.True(t, r3.Success)
}

// func (r *RegClientServer) AddDependent(ctx context.Context, in *pb.AddDependentRequest) (*pb.AddDependentResponse, error) {
// 	//HERE
// 	if _, err := helpers.GRPCAccessChecker(ctx, func(baseID string) helpers.ServiceKeyCheckInterface {
// 		return &masterData.ConnectionInfo
// 	}, "GRPC Add Dependent"); err != "" {
// 		return &pb.AddDependentResponse{Success: false, ErrorMessage: err}, nil
// 	}
// 	//TO HERE

// 	//2nd failure
// 	if s := dependentServices.GetByBaseID(in.BaseId); s != nil {
// 		return &pb.AddDependentResponse{
// 			Success: true,
// 		}, nil
// 	}

// 	dependentServices.Add(&dependentServices.DependentService{
// 		ServiceType: in.Type,
// 		ServiceName: in.Name,
// 		ServiceID:   in.Id,
// 		BaseID:      in.BaseId,
// 		Key:         in.Key,
// 		URL:         in.Url,
// 		Port:        in.Port,
// 		SingleRun:   in.SingleRun,
// 	})
// 	return &pb.AddDependentResponse{
// 		Success: true,
// 	}, nil
// }

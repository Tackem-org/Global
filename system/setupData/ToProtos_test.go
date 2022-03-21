package setupData_test

import (
	"testing"

	"github.com/Tackem-org/Global/pb/config"
	pb "github.com/Tackem-org/Global/pb/registration"
	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/Global/system/setupData"
	"github.com/stretchr/testify/assert"
)

func TestAllToProtos(t *testing.T) {
	setupData.Port = 50001
	in := &setupData.SetupData{
		ServiceName: "TestName",
		ServiceType: "TestType",
		Version:     structs.Version{Major: 1, Minor: 0, Hotfix: 0},
		Multi:       false,
		SingleRun:   false,
		StartActive: false,
		NavItems: []*pb.NavItem{{
			LinkType: pb.LinkType_Admin,
			Title:    "Test",
			Path:     "/test/",
		}},
		ConfigItems: []*pb.ConfigItem{
			{
				Key:          "test",
				DefaultValue: "test",
				Type:         config.ValueType_String,
				Label:        "Test",
				HelpText:     "Test Help",
				InputType:    pb.InputType_IText,
			},
		},
		RequiredServices: []*pb.RequiredService{},
		Groups:           []string{"TestGroup"},
		Permissions:      []string{"TestPermission"},
		AdminPaths: []*setupData.AdminPathItem{{
			Path: "testAdmin",
		}},
		Paths: []*setupData.PathItem{{
			Path: "test/",
		}},
		Sockets: []*setupData.SocketItem{{
			Command: "test.test",
		}},
	}

	expectedProto := &pb.RegisterRequest{
		ServiceName: "TestName",
		ServiceType: "TestType",
		Version:     &pb.Version{Major: 1, Minor: 0, Hotfix: 0},
		Port:        50001,
		Multi:       false,
		SingleRun:   false,
		StartActive: false,
		ConfigItems: []*pb.ConfigItem{
			{
				Key:          "test",
				DefaultValue: "test",
				Type:         config.ValueType_String,
				Label:        "Test",
				HelpText:     "Test Help",
				InputType:    pb.InputType_IText,
			},
		},
		RequiredServices: []*pb.RequiredService{},
		WebLinkItems: []*pb.WebLinkItem{{
			Path: "test/",
		}},
		AdminWebLinkItems: []*pb.AdminWebLinkItem{{
			Path: "testAdmin",
		}},
		NavItems: []*pb.NavItem{{
			LinkType: pb.LinkType_Admin,
			Title:    "Test",
			Path:     "/test/",
		}},
		WebSocketItems: []*pb.WebSocketItem{{
			Command: "test.test",
		}},
		Groups:      []string{"TestGroup"},
		Permissions: []string{"TestPermission"},
	}
	proto := in.RegisterProto()

	assert.Equal(t, expectedProto, proto)
}

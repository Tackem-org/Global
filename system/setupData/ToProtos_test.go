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
		Panels: []*setupData.PanelItem{
			{
				Name:        "test",
				Label:       "test",
				Description: "test pop up panel",
				Layout: setupData.PanelLayout{
					HorizontalAlign: setupData.HCenter,
					VerticalAlign:   setupData.VCenter,
					Width:           2,
					Height:          1,
					ScrollWidth:     false,
					ScrollHeight:    false,
					TitleBar:        true,
					Minimise:        true,
					Close:           true,
				},
				AdminOnly:         false,
				Permission:        "",
				RequiredVariables: []setupData.RequiredVariable{{Name: "userid", Options: []string{"[number]"}}},
			},
		},
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
		Panels: []*pb.PanelSetup{
			{
				Name:        "test",
				Label:       "test",
				Description: "test pop up panel",
				Layout: &pb.PanelLayout{
					Width:        2,
					Height:       1,
					ScrollWidth:  false,
					ScrollHeight: false,
					TitleBar:     true,
					Minimise:     true,
					Close:        true,
				},
				AdminOnly:         false,
				Permission:        "",
				RequiredVariables: []*pb.RequiredVariable{{Name: "userid", Options: []string{"[number]"}}},
			},
		},
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

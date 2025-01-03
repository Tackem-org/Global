package setupData_test

import (
	"fmt"
	"testing"

	"github.com/Tackem-org/Global/system/setupData"
	"github.com/stretchr/testify/assert"
)

func TestGetPath(t *testing.T) {
	path := "test/path"
	d := &setupData.SetupData{
		Paths: []*setupData.PathItem{
			{
				Path: path,
			},
		},
	}
	assert.NotNil(t, d.GetPath(path))
	assert.Nil(t, d.GetPath("SOME/Other/Path"))
}

func TestGetAdminPath(t *testing.T) {
	path := "test/path"
	d := &setupData.SetupData{
		AdminPaths: []*setupData.AdminPathItem{
			{
				Path: path,
			},
		},
	}
	assert.NotNil(t, d.GetAdminPath(path))
	assert.Nil(t, d.GetAdminPath("SOME/Other/Path"))
}

func TestGetSocket(t *testing.T) {
	command := "test.command"
	d := &setupData.SetupData{
		Sockets: []*setupData.SocketItem{
			{
				Command: command,
			},
		},
	}
	assert.NotNil(t, d.GetSocket(command))
	assert.Nil(t, d.GetSocket("SOME.Other.Path"))
}

func TestGePanel(t *testing.T) {
	name := "testpanel"
	d := &setupData.SetupData{
		Panels: []*setupData.PanelItem{
			{
				Name: name,
			},
		},
	}
	assert.NotNil(t, d.GetPanel(name))
	assert.Nil(t, d.GetPanel("someotherpanel"))
}

func TestName(t *testing.T) {
	sn1 := "testName"
	st1 := "testService"
	d1 := &setupData.SetupData{
		ServiceType: st1,
		ServiceName: sn1,
	}
	assert.Equal(t, fmt.Sprintf("%s %s", st1, sn1), d1.Name())

	sn2 := "testName2"
	st2 := "system"
	d2 := &setupData.SetupData{
		ServiceType: st2,
		ServiceName: sn2,
	}
	assert.Equal(t, sn2, d2.Name())
}

func TestFilename(t *testing.T) {
	sn1 := "testFilename"
	st1 := "testService"
	d1 := &setupData.SetupData{
		ServiceType: st1,
		ServiceName: sn1,
	}
	assert.Equal(t, fmt.Sprintf("%s-%s.json", st1, sn1), d1.Filename("json"))

	sn2 := "testFilename2"
	st2 := "system"
	d2 := &setupData.SetupData{
		ServiceType: st2,
		ServiceName: sn2,
	}
	assert.Equal(t, fmt.Sprintf("%s.json", sn2), d2.Filename("json"))
}

func TestURL(t *testing.T) {
	sn1 := "testName"
	st1 := "testService"
	d1 := &setupData.SetupData{
		ServiceType: st1,
		ServiceName: sn1,
	}
	assert.Equal(t, fmt.Sprintf("%s/%s", st1, sn1), d1.URL())

	sn2 := "testName2"
	st2 := "system"
	d2 := &setupData.SetupData{
		ServiceType: st2,
		ServiceName: sn2,
	}
	assert.Equal(t, sn2, d2.URL())
}

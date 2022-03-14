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

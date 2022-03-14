package structs_test

import (
	"testing"

	"github.com/Tackem-org/Global/structs"
	pb "github.com/Tackem-org/Proto/pb/remoteweb"
	"github.com/stretchr/testify/assert"
)

func TestUserDataHasPermission(t *testing.T) {
	adminUser := structs.UserData{
		ID:          0,
		Name:        "test",
		Icon:        "",
		IsAdmin:     true,
		Permissions: []string{},
	}
	allowedUser := structs.UserData{
		ID:          0,
		Name:        "test",
		Icon:        "",
		IsAdmin:     false,
		Permissions: []string{"test"},
	}
	uselessUser := structs.UserData{
		ID:          0,
		Name:        "test",
		Icon:        "",
		IsAdmin:     false,
		Permissions: []string{},
	}
	assert.True(t, adminUser.HasPermission("test"))
	assert.True(t, allowedUser.HasPermission("test"))
	assert.False(t, uselessUser.HasPermission("test"))
}

func TestGetUserData(t *testing.T) {
	new := pb.UserData{
		UserId:      5,
		Name:        "test",
		Icon:        "SPECIAL",
		IsAdmin:     false,
		Permissions: []string{"test"},
	}
	expected := structs.UserData{
		ID:          5,
		Name:        "test",
		Icon:        "SPECIAL",
		IsAdmin:     false,
		Permissions: []string{"test"},
	}
	data := structs.GetUserData(&new)
	assert.Equal(t, expected.ID, data.ID)
	assert.Equal(t, expected.Name, data.Name)
	assert.Equal(t, expected.Icon, data.Icon)
	assert.Equal(t, expected.IsAdmin, data.IsAdmin)
	assert.Equal(t, expected.Permissions, data.Permissions)
}

package structs

import (
	pb "github.com/Tackem-org/Global/pb/remoteweb"
)

type UserData struct {
	ID          uint64
	Name        string
	Icon        string
	IsAdmin     bool
	Permissions []string
}

func (u *UserData) HasPermission(name string) bool {
	if name == "" || u.IsAdmin {
		return true
	}
	for _, v := range u.Permissions {
		if v == name {
			return true
		}
	}
	return false
}

func GetUserData(in *pb.UserData) *UserData {
	return &UserData{
		ID:          in.UserId,
		Name:        in.Name,
		Icon:        in.Icon,
		IsAdmin:     in.IsAdmin,
		Permissions: in.Permissions,
	}
}

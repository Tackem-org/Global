package structs

import (
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	pb "github.com/Tackem-org/Proto/pb/remoteweb"
)

type UserData struct {
	ID          uint64
	Name        string
	Initial     string
	Icon        string
	IsAdmin     bool
	Permissions []string
}

func (u *UserData) HasPermission(name string) bool {
	if u.IsAdmin {
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
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[structs.GetUserData(r *http.Request) UserData]")
	return &UserData{
		ID:          in.UserId,
		Name:        in.Name,
		Initial:     in.Initial,
		Icon:        in.Icon,
		IsAdmin:     in.IsAdmin,
		Permissions: in.Permissions,
	}
}

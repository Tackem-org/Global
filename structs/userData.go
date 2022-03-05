package structs

import (
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	pb "github.com/Tackem-org/Proto/pb/remoteweb"
)

type UserData struct {
	ID          uint64
	Name        string
	Icon        string
	IsAdmin     bool
	Permissions []string
}

func (u *UserData) HasPermission(name string) bool {
	logging.Debugf(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.structs.UserData{%s}.HasPermission", u.Name)
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] name=%s", name)
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
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.structs.GetUserData")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] in=%v", in)
	return &UserData{
		ID:          in.UserId,
		Name:        in.Name,
		Icon:        in.Icon,
		IsAdmin:     in.IsAdmin,
		Permissions: in.Permissions,
	}
}

package structs

import (
	"fmt"
	"html/template"
)

type UserData struct {
	ID          uint64
	Name        string
	Initial     string
	Icon        template.URL
	IsAdmin     bool
	Permissions []string
}

func (u UserData) CheckID(num interface{}) bool {
	return fmt.Sprint(u.ID) == fmt.Sprint(num)
}
func (u UserData) HasPermission(name string) bool {
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

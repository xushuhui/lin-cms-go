package services

import (
	"lin-cms-go/internal/request"
	"lin-cms-go/pkg/core"
)

func GetUsers(req request.GetUsers) (data map[string]interface{}, err error) {
	return
}
func ChangeUserPassword(req request.ChangeUserPassword) (err error) {
	return
}
func DeleteUser(id int) (err error) {
	if id <= 0 {
		err = core.NewInvalidParamsError("id cannot be negative")
		return
	}
	return
}
func UpdateUser(req request.UpdateUser) (err error) {
	return
}

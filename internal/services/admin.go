package services

import (
	"golang.org/x/crypto/bcrypt"
	"lin-cms-go/global"
	"lin-cms-go/internal/model"
	"lin-cms-go/internal/request"
	"lin-cms-go/pkg/core"
)

func GetUsers(req request.GetUsers) (data map[string]interface{}, e error) {
	return
}
func ChangeUserPassword(req request.ChangeUserPassword) (e error) {
	userIdentityModel, err := model.GetLinUserIdentityOne("user_id=?", req)
	if err != nil {
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	userIdentityModel.Credential = string(hash)
	err = global.DBEngine.Save(&userIdentityModel).Error
	if err != nil {
		return
	}
	return
}
func DeleteUser(id int) (e error) {
	if id <= 0 {
		e = core.NewInvalidParamsError("id cannot be negative")
		return
	}
	return
}
func UpdateUser(req request.UpdateUser) (e error) {
	return
}

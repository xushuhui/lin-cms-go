package model

import (
	"lin-cms-go/global"
)

type Model struct {
	ID uint `gorm:"primarykey"`
}
type LinUser struct {
	Model
	Email    string `gorm:"column:email;type:varchar(100);default:''" json:"email"`      // 邮箱
	Username string `gorm:"column:username;type:varchar(24);default:''" json:"username"` // 手机号
	Nickname string `gorm:"column:nickname;type:varchar(24);default:''" json:"nickname"` // 手机号
	Avatar   string `gorm:"column:avatar;type:varchar(500);default:''" json:"avatar"`    // 密码

}
type LinUserIdentity struct {
	Model
	UserId       uint   `gorm:"column:user_id;" json:"user_id"`
	IdentityType string `gorm:"column:identity_type;" json:"identity_type"`
	Identifier   string `gorm:"column:identifier;" json:"identifier"`
	Credential   string `gorm:"column:credential;" json:"credential"`
}
type LinUserGroup struct {
	Model
	UserId  uint `gorm:"column:user_id;" json:"user_id"`
	GroupId int  `gorm:"column:group_id;" json:"group_id"`
}

func GetLinUserOne(where string, args ...interface{}) (model LinUser, err error) {
	err = global.DBEngine.First(&model, where, args).Error
	return
}
func GetLinUserByID(id uint) (model LinUser, err error) {
	err = global.DBEngine.First(&model, id).Error
	return
}
func GetLinUserIdentityOne(where string, args ...interface{}) (model LinUserIdentity, err error) {
	err = global.DBEngine.First(&model, where, args).Error
	return
}

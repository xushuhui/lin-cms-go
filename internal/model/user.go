package model

import (
	"lin-cms-go/global"
)

const GroupLevelRoot = 1

type Model struct {
	ID uint `gorm:"primarykey" json:"id"`
}
type User struct {
	ID           uint   `gorm:"primarykey" json:"id"`
	Email        string `gorm:"column:email;type:varchar(100);default:''" json:"email"`
	Username     string `gorm:"column:username;type:varchar(24);default:''" json:"username"`
	Nickname     string `gorm:"column:nickname;type:varchar(24);default:''" json:"nickname"`
	Avatar       string `gorm:"column:avatar;type:varchar(500);default:''" json:"avatar"`
	UserGroup    UserGroup
	UserIdentity UserIdentity
}
type UserIdentity struct {
	ID uint `gorm:"primarykey" json:"id"`

	UserID       uint   `gorm:"column:user_id;" json:"user_id"`
	IdentityType string `gorm:"column:identity_type;" json:"identity_type"`
	Identifier   string `gorm:"column:identifier;" json:"identifier"`
	Credential   string `gorm:"column:credential;" json:"credential"`
}
type UserGroup struct {
	ID      uint `gorm:"primarykey" json:"id"`
	UserID  uint `gorm:"column:user_id;" json:"user_id"`
	GroupID int  `gorm:"column:group_id;" json:"group_id"`
	Group   Group
}
type Group struct {
	ID    uint   `gorm:"primarykey" json:"id"`
	Name  string `gorm:"column:name;type:varchar(60);default:''" json:"name"`
	Info  string `gorm:"column:info" json:"info"`
	Level int    `gorm:"column:level" json:"level"`
}

func GetLinUserOne(where string, args ...interface{}) (model User, err error) {
	err = global.DBEngine.First(&model, where, args).Error
	return
}
func GetLinUserByID(id uint) (model User, err error) {
	err = global.DBEngine.First(&model, id).Error
	return
}
func GetLinUserIdentityOne(where string, args ...interface{}) (model UserIdentity, err error) {
	err = global.DBEngine.First(&model, where, args).Error
	return
}
func GetLinUserGroupByUid(uid uint) (groupId int, err error) {
	var userGroupModel UserGroup
	err = global.DBEngine.Model(&userGroupModel).Where("user_id=?", uid).Pluck("group_id", &groupId).Error
	return
}
func GetLinGroupById(id int) (model Group, err error) {
	err = global.DBEngine.First(&model, id).Error
	return
}

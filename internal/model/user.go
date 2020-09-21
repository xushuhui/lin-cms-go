package model

import (
	"lin-cms-go/global"
)

const GroupLevelRoot = 1

type Model struct {
	ID uint `gorm:"primarykey" json:"id"`
}
type User struct {
	ID           uint         `gorm:"primarykey" json:"id"`
	Email        string       `gorm:"column:email;type:varchar(100);default:''" json:"email"`
	Username     string       `gorm:"column:username;type:varchar(24);default:''" json:"username"`
	Nickname     string       `gorm:"column:nickname;type:varchar(24);default:''" json:"nickname"`
	Avatar       string       `gorm:"column:avatar;type:varchar(500);default:''" json:"avatar"`
	UserGroup    UserGroup    `json:"userGroup"`
	UserIdentity UserIdentity `json:"userIdentity"`
}
type UserIdentity struct {
	ID           uint   `gorm:"primarykey" json:"id"`
	UserID       uint   `gorm:"column:user_id;" json:"user_id"`
	IdentityType string `gorm:"column:identity_type;" `
	Identifier   string `gorm:"column:identifier;"`
	Credential   string `gorm:"column:credential;" json:"-"`
}
type UserGroup struct {
	ID      uint `gorm:"primarykey" json:"id"`
	UserID  uint `gorm:"column:user_id;" json:"user_id"`
	GroupID uint `gorm:"column:group_id;" json:"group_id"`
	Group   Group
}
type Group struct {
	ID              uint   `gorm:"primarykey" json:"id"`
	Name            string `gorm:"column:name;type:varchar(60);default:''" json:"name"`
	Info            string `gorm:"column:info" json:"info"`
	Level           int    `gorm:"column:level" json:"level"`
	GroupPermission GroupPermission
}
type GroupPermission struct {
	ID           uint `gorm:"primarykey" json:"id"`
	GroupID      uint `gorm:"column:group_id;" json:"group_id"`
	PermissionID uint `gorm:"column:permission_id;" json:"permission_id"`
	Permission   Permission
}
type Permission struct {
	ID     uint   `gorm:"primarykey" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	Mount  int    `gorm:"column:mount" json:"mount"`
	Module string `gorm:"column:module" json:"module"`
}
type Log struct {
	ID         uint   `gorm:"primarykey" json:"id"`
	UserID     uint   `gorm:"column:user_id;" json:"user_id"`
	Message    string `gorm:"column:message" json:"message"`
	Username   string `gorm:"column:username" json:"username"`
	StatusCode int    `gorm:"column:status_code" json:"status_code"`
	Method     string `gorm:"column:method" json:"method"`
	Path       string `gorm:"column:path" json:"path"`
	Permission string `gorm:"column:permission" json:"permission"`
}
type File struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Path      string `gorm:"column:path" json:"path"`
	Type      int    `gorm:"column:type" json:"type"`
	Name      string `gorm:"column:name" json:"name"`
	Extension string `gorm:"column:extension" json:"extension"`
	Size      int    `gorm:"column:size" json:"size"`
	Md5       string `gorm:"column:md5" json:"md5"`
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
func GetLinUserGroupByUid(uid uint) (userGroupModel UserGroup, err error) {
	err = global.DBEngine.Where("user_id=?", uid).Find(&userGroupModel).Error
	return
}
func GetLinGroupById(id int) (model Group, err error) {
	err = global.DBEngine.First(&model, id).Error
	return
}

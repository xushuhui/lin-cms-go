// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameLinUser = "lin_user"

// LinUser mapped from table <lin_user>
type LinUser struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Username  string         `gorm:"column:username;not null;comment:用户名，唯一" json:"username"` // 用户名，唯一
	Nickname  string         `gorm:"column:nickname;comment:用户昵称" json:"nickname"`            // 用户昵称
	Avatar    string         `gorm:"column:avatar;comment:头像url" json:"avatar"`               // 头像url
	Phone     string         `gorm:"column:phone;comment:邮箱" json:"phone"`                    // 邮箱
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName LinUser's table name
func (*LinUser) TableName() string {
	return TableNameLinUser
}

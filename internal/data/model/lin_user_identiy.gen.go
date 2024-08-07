// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameLinUserIdentiy = "lin_user_identiy"

// LinUserIdentiy mapped from table <lin_user_identiy>
type LinUserIdentiy struct {
	ID           int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID       int64          `gorm:"column:user_id;not null;comment:用户id" json:"user_id"` // 用户id
	IdentityType string         `gorm:"column:identity_type;not null" json:"identity_type"`
	Identifier   string         `gorm:"column:identifier" json:"identifier"`
	Credential   string         `gorm:"column:credential" json:"credential"`
	CreatedAt    time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName LinUserIdentiy's table name
func (*LinUserIdentiy) TableName() string {
	return TableNameLinUserIdentiy
}

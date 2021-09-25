package model

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

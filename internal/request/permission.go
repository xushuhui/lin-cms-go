package request

type DispatchPermissions struct {
	GroupId       int   `json:"group_id" binding:"required"`
	PermissionIds []int `json:"permission_ids" binding:"required"`
}
type DispatchPermission struct {
	GroupId      int `json:"group_id" binding:"required"`
	PermissionId int `json:"permission_id" binding:"required"`
}

type RemovePermissions struct {
	GroupId       int   `json:"group_id" binding:"required"`
	PermissionIds []int `json:"permission_ids" binding:"required"`
}

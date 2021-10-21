package request

type DispatchPermissions struct {
	GroupId       int   `json:"group_id" validate:"required"`
	PermissionIds []int `json:"permission_ids" validate:"required"`
}
type DispatchPermission struct {
	GroupId      int `json:"group_id" validate:"required"`
	PermissionId int `json:"permission_id" validate:"required"`
}

type RemovePermissions struct {
	GroupId       int   `json:"group_id" validate:"required"`
	PermissionIds []int `json:"permission_ids" validate:"required"`
}

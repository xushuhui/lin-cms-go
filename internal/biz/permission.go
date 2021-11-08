package biz

import (
	"context"
	"lin-cms-go/internal/data"
	"lin-cms-go/internal/request"
)

func GetAllPermissions() (data map[string]interface{}, err error) {
	return
}

func DispatchPermission(req request.DispatchPermission) (err error) {
	return
}

func DispatchPermissions(ctx context.Context, groupId int, permissionIds []int) (err error) {
	err = data.BatchCreateGroupPermission(ctx, groupId, permissionIds)
	return
}
func RemovePermissions(req request.RemovePermissions) (err error) {
	return
}

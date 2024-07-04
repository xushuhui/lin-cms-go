package biz

import (
	"context"

	"lin-cms-go/api"
)

type PermissionRepo interface{}

func GetAllPermissions(ctx context.Context) (res interface{}, err error) {
	// list, err := data.ListAllPermissions(ctx)
	// if err != nil {
	// 	return
	// }
	// m := make(map[string][]model.LinPermission)
	// for _, v := range list {
	// 	m[v.Module] = append(m[v.Module], *v)
	// }
	// res = m
	return
}

func DispatchPermission(req api.DispatchPermissionRequest) (err error) {
	return
}

func DispatchPermissions(ctx context.Context, groupId int, permissionIds []int) (err error) {
	// err = data.BatchCreateGroupPermission(ctx, groupId, permissionIds)

	return
}

func RemovePermissions(ctx context.Context, groupId int, permissionIds []int) (err error) {
	// _, err = data.GetGroupPermissionByGroupId(ctx, groupId)
	// if err != nil {
	// 	return
	// }
	// err = data.DeleteGroupPermission(ctx, groupId, permissionIds)
	return
}

func CreateIfNoPermissions(ctx context.Context, p Permission) (err error) {
	// _, err = data.GetPermission(ctx, p.Name, p.Module)
	// if model.IsNotFound(err) {
	// 	err = data.CreatePermission(context.Background(), p.Name, p.Module)
	// }

	return
}

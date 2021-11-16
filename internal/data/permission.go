package data

import (
	"context"
	"lin-cms-go/internal/data/model"
	"lin-cms-go/internal/data/model/linpermission"
)

func GetLinPermissionById(ctx context.Context, id int) (permission *model.LinPermission, err error) {
	permission, err = GetDB().LinPermission.Query().Where(linpermission.ID(id)).First(ctx)
	return
}

func GetPermission(ctx context.Context, name, module string) (p *model.LinPermission, err error) {
	p, err = GetDB().LinPermission.Query().Where(linpermission.Module(module), linpermission.Name(name)).First(ctx)
	return
}

func CreatePermission(ctx context.Context, name, module string) (err error) {
	_, err = GetDB().LinPermission.Create().SetModule(module).SetName(name).SetMount(1).Save(ctx)
	return
}

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
func GetUserPermission(ctx context.Context, userId int, name, module string) {

}

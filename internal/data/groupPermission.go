package data

import (
	"context"
	"lin-cms-go/internal/data/model"
)

func BatchCreateGroupPermission(ctx context.Context, groupId int, permissionId []int) (err error) {
	bulk := make([]*model.LinGroupPermissionCreate, len(permissionId))
	for i, pid := range permissionId {
		bulk[i] = GetDB().LinGroupPermission.Create().SetGroupID(groupId).SetPermissionID(pid)
	}
	_, err = GetDB().LinGroupPermission.CreateBulk(bulk...).Save(ctx)
	return
}

package data

import (
	"context"
	"lin-cms-go/internal/data/model"
	"lin-cms-go/internal/data/model/lingrouppermission"
)

func BatchCreateGroupPermission(ctx context.Context, groupId int, permissionId []int) (err error) {
	bulk := make([]*model.LinGroupPermissionCreate, 0)
	for _, pid := range permissionId {
		bulk = append(bulk, GetDB().LinGroupPermission.Create().SetGroupID(groupId).SetPermissionID(pid))
	}
	_, err = GetDB().LinGroupPermission.CreateBulk(bulk...).Save(ctx)
	return
}

func GetGroupPermissionByGroupId(ctx context.Context, groupId int) (groupPermission []*model.LinGroupPermission, err error) {
	groupPermission, err = GetDB().LinGroupPermission.Query().Where(lingrouppermission.GroupID(groupId)).All(ctx)
	return
}

func DeleteGroupPermission(ctx context.Context, groupId int) (err error) {
	groupPermission, err := GetGroupPermissionByGroupId(ctx, groupId)
	if err != nil {
		return err
	}
	if len(groupPermission) == 0 {
		return nil
	}

	_, err = GetDB().LinGroupPermission.Delete().Where(lingrouppermission.GroupIDEQ(groupId)).Exec(ctx)
	return
}

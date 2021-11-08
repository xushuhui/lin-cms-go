package data

import (
	"context"
	"lin-cms-go/internal/data/model/lingroup"
)

func BatchCreateGroupPermission(ctx context.Context, groupId int, permissionId []int) (err error) {

	_, err = GetDB().LinGroup.Update().Where(lingroup.ID(groupId)).AddLinPermissionIDs(permissionId...).Save(ctx)
	return
}

//func GetGroupPermissionByGroupId(ctx context.Context, groupId int) (groupPermission []*model.LinGroupPermission, err error) {
//	groupPermission, err = GetDB().LinGroupPermission.Query().Where(lingrouppermission.GroupID(groupId)).All(ctx)
//	return
//}

//func DeleteGroupPermission(ctx context.Context, groupId int) (err error) {
//	groupPermission, err := GetGroupPermissionByGroupId(ctx, groupId)
//	if err != nil {
//		return err
//	}
//	if len(groupPermission) == 0 {
//		return nil
//	}
//
//	_, err = GetDB().LinGroupPermission.Delete().Where(lingrouppermission.GroupIDEQ(groupId)).Exec(ctx)
//	return
//}

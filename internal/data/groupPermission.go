package data

import (
	"context"
)

// func ListAllPermissions(ctx context.Context) (list []*model.LinPermission, err error) {
// 	list, err = GetDB().LinPermission.Query().All(ctx)
// 	return
// }
// func BatchCreateGroupPermission(ctx context.Context, groupId int, permissionId []int) (err error) {

// 	_, err = GetDB().LinGroup.Update().Where(lingroup.ID(groupId)).AddLinPermissionIDs(permissionId...).Save(ctx)
// 	return
// }

// func GetGroupPermissionByGroupId(ctx context.Context, groupId int) (groupPermission *model.LinGroup, err error) {
// 	groupPermission, err = GetDB().LinGroup.Query().Where(lingroup.ID(groupId)).WithLinPermission().First(ctx)
// 	return
// }

func DeleteGroupPermission(ctx context.Context, groupId int, permissionId []int) (err error) {
	//TODO 删除前判断数据是否存在放biz,不然error不好处理，职责不清晰，有可能别人调用的时候又查一次，从命名上看这个方法就是做删除，不用里面加其他操作
	//group, err := GetGroupPermissionByGroupId(ctx, groupId)
	//if err != nil {
	//	return err
	//}

	// _, err = GetDB().LinGroup.Update().Where(lingroup.ID(groupId)).RemoveLinPermissionIDs(permissionId...).Save(ctx)
	//_, err = GetDB().LinGroupPermission.Delete().Where(lingrouppermission.GroupIDEQ(groupId)).Exec(ctx)
	return
}

// func ClearLinPermission(ctx context.Context, groupId int) (err error) {
// 	_, err = GetDB().LinGroup.Update().Where(lingroup.ID(groupId)).ClearLinPermission().Save(ctx)
// 	return
// }

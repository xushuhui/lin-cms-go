package data

import (
	"lin-cms-go/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type groupRepo struct {
	data *Data
	log  *log.Helper
}

func NewGroupRepo(data *Data, logger log.Logger) biz.GroupRepo {
	return &groupRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// func GetAllGroup(ctx context.Context) (groups []*model.LinGroup, err error) {
// 	groups, err = GetDB().LinGroup.Query().All(ctx)
// 	return
// }

// func GetLinGroupById(ctx context.Context, groupId int) (group *model.LinGroup, err error) {
// 	group, err = GetDB().LinGroup.Query().Where(lingroup.ID(groupId)).First(ctx)
// 	return
// }

// func GetRootLinGroup(ctx context.Context, groupId int) (group *model.LinGroup, err error) {
// 	group, err = GetDB().LinGroup.Query().Where(lingroup.ID(groupId)).Where(lingroup.And(lingroup.LevelGT(enum.ROOT))).First(ctx)
// 	return
// }

// func GetLinGroupByName(ctx context.Context, name string) (group *model.LinGroup, err error) {
// 	group, _ = GetDB().LinGroup.Query().Where(lingroup.Name(name)).First(ctx)
// 	return
// }

// func CreateGroup(ctx context.Context, name string, info string, level int8) (res *model.LinGroup, err error) {
// 	res, err = GetDB().LinGroup.Create().SetName(name).SetInfo(info).SetLevel(level).Save(ctx)
// 	if err != nil {
// 		return
// 	}
// 	return
// }

// func UpdateGroup(ctx context.Context, id int, name string, info string) (err error) {
// 	_, err = GetDB().LinGroup.UpdateOneID(id).SetName(name).SetInfo(info).Save(ctx)
// 	return
// }

// // update delete_time
// func DeleteGroup(ctx context.Context, id int) (err error) {
// 	_, err = GetDB().LinGroup.Update().SetDeleteTime(time.Now()).ClearLinPermission().ClearLinUser().Where(lingroup.ID(id)).Save(ctx)

// 	return err
// }

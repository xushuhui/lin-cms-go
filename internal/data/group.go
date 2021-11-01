package data

import (
	"context"
	"lin-cms-go/internal/data/model"
	"lin-cms-go/internal/data/model/lingroup"
	"lin-cms-go/pkg/enum"
)

func GetAllGroup(ctx context.Context) (groups []*model.LinGroup, err error) {
	groups, err = GetDB().LinGroup.Query().All(ctx)
	return
}

func GetLinGroupById(ctx context.Context, groupId int) (group *model.LinGroup, err error) {
	group, err = GetDB().LinGroup.Query().Where(lingroup.ID(groupId)).Where(lingroup.And(lingroup.LevelGT(enum.ROOT))).First(ctx)
	return
}

func GetLinGroupByName(ctx context.Context, name string) (group *model.LinGroup, err error) {
	group, _ = GetDB().LinGroup.Query().Where(lingroup.Name(name)).First(ctx)
	return
}

func CreateGroup(ctx context.Context, name string, info string) (res *model.LinGroup, err error) {
	res, err = GetDB().LinGroup.Create().SetName(name).SetInfo(info).SetLevel(3).Save(ctx)
	if err != nil {
		return
	}
	return
}

func UpdateGroup(ctx context.Context, id int, name string, info string) (err error) {
	_, err = GetDB().LinGroup.UpdateOneID(id).SetName(name).SetInfo(info).Save(ctx)
	return
}

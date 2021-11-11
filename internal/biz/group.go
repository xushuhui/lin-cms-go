package biz

import (
	"context"
	"github.com/xushuhui/goal/core"
	"lin-cms-go/internal/data"
	"lin-cms-go/internal/data/model"
	"lin-cms-go/internal/request"
	"lin-cms-go/pkg/enum"
	"lin-cms-go/pkg/errcode"
)

type Group struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Info  string `json:"info"`
	Level int8   `json:"level"`
}

func GetGroups(ctx context.Context) (res interface{}, err error) {
	var linGroupModel []*model.LinGroup
	linGroupModel, err = data.GetAllGroup(ctx)
	res = simplifyGroup(linGroupModel)
	return
}

func GetGroup(ctx context.Context, id int) (res interface{}, err error) {
	var linGroupModel *model.LinGroup
	linGroupModel, err = data.GetLinGroupById(ctx, id)
	if model.IsNotFound(err) {
		err = core.NewErrorCode(errcode.GroupNotFound)
		return
	}
	res = Group{Id: linGroupModel.ID, Name: linGroupModel.Name, Info: linGroupModel.Info, Level: linGroupModel.Level}
	return
}

func simplifyGroup(groupModel []*model.LinGroup) []Group {
	res := make([]Group, 0, len(groupModel))
	for _, v := range groupModel {
		res = append(res, Group{Id: v.ID, Name: v.Name, Info: v.Info, Level: v.Level})
	}
	return res
}

func CreateGroup(ctx context.Context, name string, info string, permissionIds []int) (err error) {
	group, _ := data.GetLinGroupByName(ctx, name)
	if group != nil {
		err = core.NewErrorCode(errcode.GroupFound)
		return
	}
	for _, v := range permissionIds {
		_, err = data.GetLinPermissionById(ctx, v)
		if model.IsNotFound(err) {
			err = core.NewErrorCode(errcode.PermissionNotFound)
			return
		}
	}
	groupModel, err := data.CreateGroup(ctx, name, info, enum.USER)
	if err != nil {
		return
	}

	err = data.BatchCreateGroupPermission(ctx, groupModel.ID, permissionIds)
	if err != nil {
		return
	}
	return
}

func UpdateGroup(ctx context.Context, id int, req request.UpdateGroup) (err error) {
	_, err = data.GetLinGroupById(ctx, id)
	if model.IsNotFound(err) {
		err = core.NewErrorCode(errcode.GroupNotFound)
		return
	}
	err = data.UpdateGroup(ctx, id, req.Name, req.Info)
	return
}

func DeleteGroup(ctx context.Context, id int) (err error) {
	var linGroup *model.LinGroup
	linGroup, err = data.GetLinGroupById(ctx, id)
	if model.IsNotFound(err) {
		err = core.NewErrorCode(errcode.GroupNotFound)
		return
	}
	if err != nil {
		return
	}
	if linGroup.Level == enum.ROOT {
		err = core.NewErrorCode(errcode.RootGroupNotAllowDelete)
		return
	}

	if linGroup.Level == enum.GUEST {
		err = core.NewErrorCode(errcode.GuestGroupNotAllowDelete)
		return
	}

	err = data.DeleteGroup(ctx, id)

	return
}

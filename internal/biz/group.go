package biz

import (
	"context"

	"lin-cms-go/api"
)

type (
	GroupRepo interface {
		CreateGroup(ctx context.Context, g *Group) error
		GetGroup(ctx context.Context, id int) (*Group, error)
		UpdateGroup(ctx context.Context, g *Group) error
		DeleteGroup(ctx context.Context, id int) error
		ListGroup(ctx context.Context, page, size int) ([]*Group, error)
	}
	Group struct {
		Id    int64   
		Name  string 
		Info  string 
		Level int8  
	}
)


func GetGroups(ctx context.Context) (res interface{}, err error) {
	// var linGroupModel []*model.LinGroup
	// linGroupModel, err = data.GetAllGroup(ctx)
	// res = simplifyGroup(linGroupModel)
	return
}

func GetGroup(ctx context.Context, id int) (res interface{}, err error) {
	// var linGroupModel *model.LinGroup
	// linGroupModel, err = data.GetLinGroupById(ctx, id)
	// if model.IsNotFound(err) {
	// 	err = core.NotFoundError(errcode.GroupNotFound)
	// 	return
	// }

	// res = Group{Id: linGroupModel.ID, Name: linGroupModel.Name, Info: linGroupModel.Info, Level: linGroupModel.Level}
	return
}

// func simplifyGroup(groupModel []*model.LinGroup) []Group {
// 	res := make([]Group, 0, len(groupModel))
// 	for _, v := range groupModel {
// 		res = append(res, Group{Id: v.ID, Name: v.Name, Info: v.Info, Level: v.Level})
// 	}
// 	return res
// }

func CreateGroup(ctx context.Context, name string, info string, permissionIds []int) (err error) {
	// group, _ := data.GetLinGroupByName(ctx, name)
	// if group != nil {
	// 	err = core.ParamsError(errcode.GroupFound)
	// 	return
	// }
	// for _, v := range permissionIds {
	// 	_, err = data.GetLinPermissionById(ctx, v)
	// 	if model.IsNotFound(err) {
	// 		err = core.NotFoundError(errcode.PermissionNotFound)
	// 		return
	// 	}
	// 	if err != nil {
	// 		err = errors.Wrap(err, "GetLinPermissionById err")
	// 		return
	// 	}
	// }
	// groupModel, err := data.CreateGroup(ctx, name, info, enum.USER)
	// if err != nil {
	// 	err = errors.Wrap(err, "CreateGroup err")
	// 	return
	// }

	// err = data.BatchCreateGroupPermission(ctx, groupModel.ID, permissionIds)
	// if err != nil {
	// 	err = errors.Wrap(err, "BatchCreateGroupPermission err")
	// 	return
	// }
	return
}

func UpdateGroup(ctx context.Context, id int, req *api.UpdateGroupRequest) (err error) {
	// _, err = data.GetLinGroupById(ctx, id)
	// if model.IsNotFound(err) {
	// 	err = core.NotFoundError(errcode.GroupNotFound)
	// 	return
	// }
	// if err != nil {
	// 	err = errors.Wrap(err, "GetLinGroupById")
	// 	return
	// }
	// err = data.UpdateGroup(ctx, id, req.Name, req.Info)
	// if err != nil {
	// 	err = errors.Wrap(err, "UpdateGroup")
	// 	return
	// }
	return
}

func DeleteGroup(ctx context.Context, id int) (err error) {
	// var linGroup *model.LinGroup
	// linGroup, err = data.GetLinGroupById(ctx, id)
	// if model.IsNotFound(err) {
	// 	err = core.NotFoundError(errcode.GroupNotFound)
	// 	return
	// }
	// if err != nil {
	// 	err = errors.Wrap(err, "GetLinGroupById ")
	// 	return
	// }
	// if linGroup.Level == enum.ROOT {
	// 	err = core.ParamsError(errcode.RootGroupNotAllowDelete)
	// 	return
	// }
	//
	// if linGroup.Level == enum.GUEST {
	// 	err = core.ParamsError(errcode.GuestGroupNotAllowDelete)
	// 	return
	// }
	//
	// err = data.DeleteGroup(ctx, id)
	// if err != nil {
	// 	err = errors.Wrap(err, "DeleteGroup ")
	// 	return
	// }
	return
}

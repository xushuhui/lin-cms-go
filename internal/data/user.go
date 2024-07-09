package data

import (
	"context"
	"fmt"

	"lin-cms-go/internal/biz"
	"lin-cms-go/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

type linUserRepo struct {
	data *Data
	log  *log.Helper
}

func NewLinUserRepo(data *Data, logger log.Logger) biz.LinUserRepo {
	return &linUserRepo{data: data, log: log.NewHelper(logger)}
}

func (r *linUserRepo) GetUserByUsername(ctx context.Context, username string)(*biz.LinUser, error) {
	var u model.LinUser
	err := r.data.db.First(&u, "username = ?", username).Error
	if  err != nil {
		return nil, err
	}
	return &biz.LinUser{
		ID:       u.ID,
		Username: u.Username,
		Password: u.Password,
		Phone:    u.Phone,
	}, nil
}

func (r *linUserRepo) ListUser(ctx context.Context, page, size int32, groupId int64) ([]*biz.LinUser, int64, error) {
	panic("not implemented") // TODO: Implement
}

func (r *linUserRepo) GetUser(ctx context.Context, userId int) (*biz.LinUser, error) {
	panic("not implemented") // TODO: Implement
}

func (r *linUserRepo) CreateUser(ctx context.Context, user *biz.LinUser) error {
	err := r.data.db.Create(&model.LinUser{
		Username: user.Username,
		Password: user.Password,
		Phone:    user.Phone,
		Nickname: user.Username,
	}).Error
	if err != nil {
		return fmt.Errorf("CreateUser failed %w ", err)
	}
	return nil
}

func (r *linUserRepo) ChangeUserPassword(ctx context.Context, userId int, password string) error {
	panic("not implemented") // TODO: Implement
}

func (r *linUserRepo) UpdateUser(ctx context.Context, user *biz.LinUser) error {
	panic("not implemented") // TODO: Implement
}

// func GetLinUserIdentityByIdentifier(ctx context.Context, identifier string) (model *model.LinUserIdentiy, err error) {
// 	model, err = GetDB().LinUserIdentiy.Query().Where(linuseridentiy.Identifier(identifier)).First(ctx)
// 	return
// }

// func GetLinUserIdentityByUserId(ctx context.Context, userId int) (model *model.LinUserIdentiy, err error) {
// 	model, err = GetDB().LinUserIdentiy.Query().Where(linuseridentiy.UserID(userId)).First(ctx)
// 	return
// }

// func CreateLinUser(ctx context.Context, username, password, email string, groupId int) (err error) {
// 	userModel, err := GetDB().LinUser.Create().
// 		SetUsername(username).
// 		SetNickname(username).
// 		SetEmail(email).
// 		AddLinGroupIDs(groupId).
// 		Save(ctx)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = GetDB().LinUserIdentiy.Create().
// 		SetUserID(userModel.ID).
// 		SetIdentifier(username).
// 		SetCredential(password).
// 		SetIdentityType("USERNAME_PASSWORD").
// 		Save(ctx)

// 	return
// }

// func GetLinUserById(ctx context.Context, uid int) (model *model.LinUser, err error) {
// 	model, err = GetDB().LinUser.Query().Where(linuser.ID(uid)).First(ctx)

// 	return
// }

// func GetLinUserWithGroupById(ctx context.Context, uid int) (users *model.LinUser, err error) {
// 	users, err = GetDB().LinUser.Query().Where(linuser.ID(uid)).WithLinGroup(func(query *model.LinGroupQuery) {
// 		query.WithLinPermission()
// 	}).First(ctx)

// 	return
// }

// func UpdateLinUser(ctx context.Context, uid int, avatar, nickname, email string) (err error) {
// 	_, err = GetDB().LinUser.Update().Where(linuser.ID(uid)).SetEmail(email).SetAvatar(avatar).SetNickname(nickname).
// 		Save(ctx)

// 	return
// }

// func UpdateLinUserIdentityPassword(ctx context.Context, username, password string) (err error) {
// 	_, err = GetDB().LinUserIdentiy.Update().Where(linuseridentiy.Identifier(username)).
// 		SetCredential(password).Save(ctx)

// 	return
// }

// func SoftDeleteUser(ctx context.Context, userId int) (err error) {
// 	_, err = GetDB().LinUser.Update().Where(linuser.ID(userId)).
// 		SetDeleteTime(time.Now()).Save(ctx)

// 	return
// }

// func (p *Paging) ListUser(ctx context.Context) (users []*model.LinUser, err error) {
// 	users, err = GetDB().LinUser.Query().WithLinGroup().Limit(p.Size).Offset(p.Offset).All(ctx)
// 	return
// }

// func (p *Paging) ListUserByGroupId(ctx context.Context, groupId int) (users []*model.LinUser, err error) {
// 	groups, err := GetDB().LinGroup.Query().Where(lingroup.ID(groupId)).WithLinUser(func(query *model.LinUserQuery) {
// 		query.Limit(p.Size).Offset(p.Offset).WithLinGroup()
// 	}).All(ctx)
// 	if err != nil {
// 		return
// 	}
// 	for _, g := range groups {
// 		users = append(users, g.Edges.LinUser...)
// 	}

// 	return
// }

// func AddLinUserGroupIDs(ctx context.Context, userId int, groupIds []int) (err error) {
// 	_, err = GetDB().LinUser.Update().Where(linuser.ID(userId)).ClearLinGroup().AddLinGroupIDs(groupIds...).Save(ctx)
// 	return
// }

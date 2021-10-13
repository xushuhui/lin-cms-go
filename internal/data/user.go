package data

import (
	"context"
	"lin-cms-go/internal/biz"
	"lin-cms-go/internal/data/ent"
	"lin-cms-go/internal/data/ent/linuser"
	"lin-cms-go/internal/data/ent/linuseridentiy"
)

var _ biz.LinUserRepo = (*linUserRepo)(nil)

type linUserRepo struct {
	data *Data
}

func NewLinUserRepo(data *Data) biz.LinUserRepo {
	return &linUserRepo{
		data: data,
	}
}

func (r *linUserRepo) GetLinUserIdentityByIdentifier(ctx context.Context, identifier string) (*ent.LinUserIdentiy, error) {

	po, err := r.data.db.LinUserIdentiy.Query().Where(linuseridentiy.Identifier(identifier)).First(ctx)
	if err != nil {
		return nil, err
	}
	return po, nil
}
func (r *linUserRepo) CreateLinUser(ctx context.Context, username, password, email string, groupId int) error {
	identiyModel, err := r.data.db.LinUserIdentiy.Create().
		SetIdentifier(username).
		SetCredential(password).
		Save(ctx)
	_, err = r.data.db.LinUser.Create().
		SetUsername(username).
		SetEmail(email).
		AddLinGroupIDs(groupId).
		AddLinUserIdentiy(identiyModel).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}
func (r *linUserRepo) GetLinUserById(ctx context.Context, uid int) (*ent.LinUser, error) {
	model, err := r.data.db.LinUser.Query().Where(linuser.ID(uid)).First(ctx)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (r *linUserRepo) UpdateLinUser(ctx context.Context, uid int, avatar, nickname, email string) error {
	_, err := r.data.db.LinUser.Update().Where(linuser.ID(uid)).SetEmail(email).SetAvatar(avatar).SetNickname(nickname).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}
func (r *linUserRepo) UpdateLinUserIdentityPassword(ctx context.Context, username, password string) error {
	_, err := r.data.db.LinUserIdentiy.Update().Where(linuseridentiy.Identifier(username)).
		SetCredential(password).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

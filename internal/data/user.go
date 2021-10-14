package data

import (
	"context"
	"lin-cms-go/internal/data/ent"
	"lin-cms-go/internal/data/ent/linuser"
	"lin-cms-go/internal/data/ent/linuseridentiy"
)





func  GetLinUserIdentityByIdentifier(ctx context.Context, identifier string) (*ent.LinUserIdentiy, error) {

	po, err := GetDBClient().LinUserIdentiy.Query().Where(linuseridentiy.Identifier(identifier)).First(ctx)
	if err != nil {
		return nil, err
	}
	return po, nil
}
func  CreateLinUser(ctx context.Context, username, password, email string, groupId int) error {
	identiyModel, err := GetDBClient().LinUserIdentiy.Create().
		SetIdentifier(username).
		SetCredential(password).
		Save(ctx)
	_, err = GetDBClient().LinUser.Create().
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
func  GetLinUserById(ctx context.Context, uid int) (*ent.LinUser, error) {
	model, err := GetDBClient().LinUser.Query().Where(linuser.ID(uid)).First(ctx)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func  UpdateLinUser(ctx context.Context, uid int, avatar, nickname, email string) error {
	_, err := GetDBClient().LinUser.Update().Where(linuser.ID(uid)).SetEmail(email).SetAvatar(avatar).SetNickname(nickname).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}
func  UpdateLinUserIdentityPassword(ctx context.Context, username, password string) error {
	_, err := GetDBClient().LinUserIdentiy.Update().Where(linuseridentiy.Identifier(username)).
		SetCredential(password).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

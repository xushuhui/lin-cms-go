package data

import (
	"context"
	"lin-cms-go/internal/data/ent"
	"lin-cms-go/internal/data/ent/linuser"
	"lin-cms-go/internal/data/ent/linuseridentiy"
)

func GetLinUserIdentityByIdentifier(ctx context.Context, identifier string) (po *ent.LinUserIdentiy, err error) {

	po, err = GetDB().LinUserIdentiy.Query().Where(linuseridentiy.Identifier(identifier)).First(ctx)

	return
}
func CreateLinUser(ctx context.Context, username, password, email string, groupId int) (err error) {

	userModel, err := GetDB().LinUser.Create().
		SetUsername(username).
		SetNickname(username).
		SetEmail(email).
		AddLinGroupIDs(groupId).
		Save(ctx)
	if err != nil {
		return err
	}
	_, err = GetDB().LinUserIdentiy.Create().
		SetUserID(userModel.ID).
		SetIdentifier(username).
		SetCredential(password).
		SetIdentityType("USERNAME_PASSWORD").
		Save(ctx)

	return
}
func GetLinUserById(ctx context.Context, uid int) (model *ent.LinUser, err error) {
	model, err = GetDB().LinUser.Query().Where(linuser.ID(uid)).First(ctx)

	return
}

func UpdateLinUser(ctx context.Context, uid int, avatar, nickname, email string) (err error) {
	_, err = GetDB().LinUser.Update().Where(linuser.ID(uid)).SetEmail(email).SetAvatar(avatar).SetNickname(nickname).
		Save(ctx)

	return
}
func UpdateLinUserIdentityPassword(ctx context.Context, username, password string) (err error) {
	_, err = GetDB().LinUserIdentiy.Update().Where(linuseridentiy.Identifier(username)).
		SetCredential(password).Save(ctx)

	return
}

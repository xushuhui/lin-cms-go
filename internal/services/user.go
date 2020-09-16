package services

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"lin-cms-go/global"
	"lin-cms-go/internal/model"
	"lin-cms-go/internal/request"
	"lin-cms-go/pkg/core"
	"lin-cms-go/pkg/errcode"
	"lin-cms-go/pkg/lib"
	"lin-cms-go/pkg/utils"
)

func Login(req request.Login) (data map[string]interface{}, err error) {

	// 正确密码验证
	userIdentityModel, err := model.GetLinUserIdentityOne("identifier=?", req.Username)
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(userIdentityModel.Credential), []byte(req.Password))
	if err != nil {
		err = core.NewError(errcode.ErrorPassWord)
		return
	}
	token, err := lib.GenerateToken(userIdentityModel.UserID)
	if err != nil {
		return
	}
	data = make(map[string]interface{})
	data["access_token"] = token
	return
}
func Register(req request.Register) (err error) {
	userModel, err := model.GetLinUserOne("username=?", req.Username)
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	if err != nil {
		return
	}

	if userModel.ID > 0 {
		err = core.NewError(errcode.UserExist)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	userModel = model.User{
		Username:  req.Username,
		Email:     req.Email,
		UserGroup: model.UserGroup{GroupID: req.GroupId},
		UserIdentity: model.UserIdentity{
			Identifier: req.Username,
			Credential: string(hash),
		},
	}
	global.DBEngine.Create(&userModel)
	//TODO:关联插入
	//err = global.DBEngine.Transaction(func(tx *gorm.DB) (err error) {
	//	// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
	//	err = tx.Create(&userModel).Error
	//	if err != nil {
	//		return
	//	}
	//	uid := userModel.ID
	//	list := make([]model.UserGroup, 0)
	//	//for _, v := range req.GroupIds {
	//	//
	//	//	var item model.UserGroup
	//	//	item.UserId = uid
	//	//	item.GroupId = v
	//	//	list = append(list, item)
	//	//}
	//	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	//	if err != nil {
	//		return
	//	}
	//	userIdentityModel := model.UserIdentity{
	//		UserId:     uid,
	//		Identifier: req.Username,
	//		Credential: string(hash),
	//	}
	//	err = tx.Create(&list).Error
	//	if err != nil {
	//		return
	//	}
	//	err = tx.Create(&userIdentityModel).Error
	//	if err != nil {
	//		return
	//	}
	//	// 返回 nil 提交事务
	//	return nil
	//})
	//if err != nil {
	//	err = core.NewErrorMessage(errcode.DBError, err.Error())
	//	return
	//}
	return
}
func UpdateMe(req request.UpdateMe, uid uint) (err error) {
	userModel, err := model.GetLinUserByID(uid)
	if err != nil {
		return
	}
	userModel.Avatar = req.Avatar
	userModel.Nickname = req.Nickname
	userModel.Email = req.Email
	err = global.DBEngine.Save(&userModel).Error

	if err != nil {
		return
	}
	return
}
func ChangeMyPassword(req request.ChangeMyPassword, uid uint) (err error) {
	userIdentityModel, err := model.GetLinUserIdentityOne("user_id=?", uid)
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(userIdentityModel.Credential), []byte(req.OldPassword))
	if err != nil {
		err = core.NewError(errcode.ErrorPassWord)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	userIdentityModel.Credential = string(hash)
	err = global.DBEngine.Save(&userIdentityModel).Error
	if err != nil {

		return
	}
	return
}
func GetMyPermissions(uid uint) (data map[string]interface{}, err error) {
	userModel, err := model.GetLinUserByID(uid)
	if err != nil {
		return
	}
	data = utils.Struct2MapJson(userModel)
	groupIds, err := model.GetLinUserGroupByUid(uid)
	if err != nil {
		return
	}
	var isAdmin bool

	groupModel, err := model.GetLinGroupById(groupIds)
	if err != nil {
		return
	}
	if groupModel.Level == model.GroupLevelRoot {
		isAdmin = true
	}
	data["is_admin"] = isAdmin
	return
}
func GetMyInfomation(uid uint) (data map[string]interface{}, err error) {
	return
}

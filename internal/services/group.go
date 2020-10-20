package services

import (
	"lin-cms-go/global"
	"lin-cms-go/internal/model"
	"lin-cms-go/internal/request"
)

func GetGroups() (data map[string]interface{}, e error) {
	return
}
func GetGroup(id int) (data map[string]interface{}, e error) {
	//groupModel,e := model.GetLinGroupById(id)
	//if e != nil {
	//	return
	//}
	return
}
func CreateGroup(req request.CreateGroup) (e error) {
	groupModel := model.Group{
		Name: req.Name,
		Info: req.Info,
	}
	global.DBEngine.Create(&groupModel)

	return
}
func UpdateGroup(req request.UpdateGroup) (e error) {
	return
}
func DeleteGroup(id int) (e error) {
	return
}

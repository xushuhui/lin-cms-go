package biz

import (
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
func CreateGroup(name, info string) (e error) {
	//groupModel := model.Group{
	//	Name: name,
	//	Info:info,
	//}

	return
}
func UpdateGroup(req request.UpdateGroup) (e error) {
	return
}
func DeleteGroup(id int) (e error) {
	return
}

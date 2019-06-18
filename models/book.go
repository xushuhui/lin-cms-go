package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(User))
}

//func (u *User) TableName() string {
//	return "lin_book"
//}

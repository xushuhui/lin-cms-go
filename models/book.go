package models

import "github.com/astaxie/beego/orm"

type Book struct {
	Id        int    `json:"id"`
	Title     string `orm:"size(128)"json:"title"`
	Author    string `orm:"size(128)"json:"author"`
	Summary   string `orm:"size(128)"json:"summary"`
	Image     string `orm:"size(128)"json:"image"`
	CreatedAt string `orm:"size(128)"json:"-"`
	UpdatedAt string `orm:"size(128)"json:"-"`
	DeletedAt string `orm:"size(128)"json:"-"`
}

func init() {
	orm.RegisterModel(new(Book))
}

func (b *Book) TableName() string {
	return "books"
}

func GetBookById(id int) (v *Book, err error) {
	o := orm.NewOrm()
	v = &Book{Id: id}
	if err = o.QueryTable(new(Book)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}
func GetBooks() (v []Book, err error) {
	o := orm.NewOrm()
	var book []Book
	if _, err := o.QueryTable(new(Book)).Filter("deleted_at__isnull", true).All(&book); err == nil {
		return book, nil
	}
	return nil, err
}

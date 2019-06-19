package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Book struct {
	Id        int    `json:"id"`
	Title     string `orm:"size(128)"json:"title"`
	Author    string `orm:"size(128)"json:"author"`
	Summary   string `orm:"size(128)"json:"summary"`
	Image     string `orm:"size(128)"json:"image"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
	DeletedAt string `json:"-"`
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
func DeleteBook(id int) (err error) {
	o := orm.NewOrm()
	var num int64
	if num, err = o.Delete(&Book{Id: id}); err == nil {
		fmt.Println("Number of records deleted in database:", num)
	}
	return
}
func SoftDeleteBook(id int) (err error) {
	o := orm.NewOrm()
	now := time.Now().Format("2006-01-02 15:04:05")
	var num int64
	if num, err = o.Update(&Book{Id: id, DeletedAt: now}, "deleted_at"); err == nil {
		fmt.Println("Number of records deleted in database:", num)
	}
	return
}

//UpdateUserById
func UpdateBookById(m *Book) (err error) {
	o := orm.NewOrm()
	m.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	var num int64
	if num, err = o.Update(m, "title", "author", "summary", "image", "updated_at"); err == nil {
		fmt.Println("Number of records updated in database:", num)
	}
	return
}

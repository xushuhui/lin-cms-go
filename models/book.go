/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Book struct {
	Id        int    `json:"id"`
	Title     string `orm:"size(128)" json:"title"`
	Author    string `orm:"size(128)" json:"author"`
	Summary   string `orm:"size(128)" json:"summary"`
	Image     string `orm:"size(128)" json:"image"`
	CreatedAt int64  `json:"-"`
	UpdatedAt int64  `json:"-"`
	DeletedAt int64  `json:"-"`
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
	if err = o.QueryTable(new(Book)).Filter("Id", id).Filter("deleted_at", 0).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}
func GetBooks() (v []Book, err error) {
	o := orm.NewOrm()
	var book []Book
	if _, err := o.QueryTable(new(Book)).Filter("deleted_at", 0).All(&book); err == nil {
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
	now := time.Now().Unix()
	var num int64
	if num, err = o.Update(&Book{Id: id, DeletedAt: now}, "deleted_at"); err == nil {
		fmt.Println("Number of records deleted in database:", num)
	}
	return
}

//UpdateUserById
func UpdateBookById(m *Book) (err error) {
	o := orm.NewOrm()
	m.UpdatedAt = time.Now().Unix()
	var num int64
	if num, err = o.Update(m, "title", "author", "summary", "image", "updated_at"); err == nil {
		fmt.Println("Number of records updated in database:", num)
	}
	return
}
func AddBook(m *Book) (id int64, err error) {
	o := orm.NewOrm()
	now := time.Now().Unix()
	m.CreatedAt = now
	m.UpdatedAt = now
	id, err = o.Insert(m)
	return
}

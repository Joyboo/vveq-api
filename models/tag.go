package models

import (
	"github.com/astaxie/beego/orm"
)

type Tag struct {
	Id      int64
	Name    string
	Status  int
	Insuid  int64
	Instime int64
}

func (this *Tag) TableName() string {
	return "tag"
}

func NewTag() *Tag {
	return &Tag{}
}

func (this *Tag) Add() (int64, error) {
	return 0, nil
}

// 获取全部标签
func (this *Tag) GetAll() ([]*Tag, error) {
	tag := []*Tag{}
	_, err := orm.NewOrm().QueryTable(this.TableName()).Filter("status", 1).All(&tag)
	return tag, err
}

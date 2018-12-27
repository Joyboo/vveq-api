package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Tag struct {
	Id      int64
	Name    string
	Status  int
	Insuid  int64
	Instime int64
}

func NewTag() *Tag {
	return &Tag{}
}

func (this *Tag) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Tag)).Filter("status", 1)
}

func (this *Tag) Add() (int64, error) {
	this.Status = 1
	this.Instime = time.Now().Unix()
	return orm.NewOrm().Insert(this)
}

// 获取全部标签
func (this *Tag) GetAll() ([]*Tag, error) {
	tag := []*Tag{}
	_, err := this.Query().All(&tag)
	return tag, err
}

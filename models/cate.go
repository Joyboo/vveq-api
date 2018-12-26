package models

import "github.com/astaxie/beego/orm"

type Cate struct {
	Id      int64
	Pid     int
	Name    string
	Sort    int
	Route   string
	Status  int
	Insuid  int64
	Instime int64
	Upduid  int64
	Updtime int64
}

func (this *Cate) TableName() string {
	return "cate"
}

func NewCate() *Cate {
	return &Cate{}
}

func (this *Cate) Add() (int64, error) {
	return 0, nil
}

// 获取全部分类
func (this *Cate) GetAll() ([]*Cate, error) {
	cate := []*Cate{}
	_, err := orm.NewOrm().QueryTable(this.TableName()).Filter("status", 1).OrderBy("-sort").All(&cate)
	return cate, err
}

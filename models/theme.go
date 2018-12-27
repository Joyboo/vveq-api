package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Theme struct {
	Id            int64
	Cid           int
	Title         string
	Content       string
	Uid           int64
	Tagid         int
	Sort          int
	Click         int64
	Like          int64
	Status        int
	Instime       int64
	Lastreplytime int64
}

var (
	pageSize = 20
)

func NewTheme() *Theme {
	return &Theme{}
}

func (this *Theme) Add() (int64, error) {
	this.Status = 1
	this.Instime = time.Now().Unix()
	return orm.NewOrm().Insert(this)
}

func (this *Theme) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Theme)).Filter("status", 1)
}

func (this *Theme) Get(page int) ([]*Theme, error) {
	theme := []*Theme{}
	_, err := this.Query().OrderBy("-sort").Limit(pageSize, page).All(&theme)
	return theme, err
}

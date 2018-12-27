package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Theme struct {
	Id      int64
	Cid     int
	Title   string
	Content string
	Uid     int64
	Tagid   int
	Sort    int
	Click   int64
	Like    int64
	Status  int
	Instime time.Time `orm:"column(instime);auto_now_add;type(datetime)"`
}

var (
	pageSize = 20
)

func NewTheme() *Theme {
	return &Theme{}
}

func (this *Theme) Add() (int64, error) {
	this.Status = 1
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

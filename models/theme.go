package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Theme struct {
	Id      int64     `json:"id"`
	Cid     int       `json:"cid"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Uid     int64     `json:"uid"`
	Tagid   int       `json:"tagid"`
	Sort    int       `json:"sort"`
	Click   int64     `json:"click"`
	Like    int64     `json:"like"`
	Status  int       `json:"status"`
	Instime time.Time `orm:"column(instime);auto_now_add;type(datetime)" json:"instime"`
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

func (this *Theme) GetThemeById(id int64) (Theme, error) {
	var theme Theme
	err := this.Query().Filter("id", id).One(&theme)
	return theme, err
}

func (this *Theme) Gets(page int) ([]*Theme, error) {
	theme := []*Theme{}
	_, err := this.Query().OrderBy("-sort", "-instime").Limit(pageSize, page).All(&theme)
	return theme, err
}

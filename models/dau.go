package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Dau struct {
	Id      int64     `json:"id"`
	Uid     int64     `json:"uid"`
	Ip      int64     `json:"ip"`
	Instime time.Time `orm:"column(instime);auto_now_add;type(datetime)" json:"instime"`
}

func NewDau() *Dau {
	return &Dau{}
}

func (this *Dau) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Dau))
}

func (this *Dau) Add() (int64, error) {
	// TODO 走redis队列
	return orm.NewOrm().Insert(this)
}

// 记录
func (this *Dau) Log(uid int64, ip string) {
	go func() {
		this.Id = uid
		this.Ip = IpString2Int(ip)
		this.Add()
	}()
}

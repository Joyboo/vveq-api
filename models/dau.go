package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Dau struct {
	Id      int64
	Uid     int64
	Ip      int64
	Instime int64
}

func NewDau() *Dau {
	return &Dau{}
}

func (this *Dau) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Dau))
}

func (this *Dau) Add() (int64, error) {
	this.Instime = time.Now().Unix()
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

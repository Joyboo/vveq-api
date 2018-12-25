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

func (this *Dau) TableName() string {
	return "dau"
}

func newDau() *Dau {
	return &Dau{}
}

func (this *Dau) Add() (int64, error) {
	this.Instime = time.Now().Unix()
	// TODO 走redis队列
	return orm.NewOrm().Insert(this)
}

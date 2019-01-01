package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Like struct {
	Id      int64     `json:"id"`
	Uid     int64     `json:"uid"`
	Pid     int64     `json:"pid"`
	Type    int       `json:"type"`
	Status  int       `json:"status"`
	Instime time.Time `orm:"column(instime);auto_now_add;type(datetime)" json:"instime"`
}

func NewLike() *Like {
	return &Like{}
}

func (this *Like) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Like)).Filter("status", 1)
}

// 点赞事件
func (this *Like) Add() bool {
	// 重复写入时视为取消点赞
	sql := "INSERT INTO `like` SET type=?,uid = ?,pid=?,status=1,instime=? ON DUPLICATE KEY UPDATE status=IF(status=0,1,0)"
	res, _ := orm.NewOrm().Raw(sql, this.Type, this.Uid, this.Pid, time.Now().Format(TimeFormart)).Exec()
	num, _ := res.RowsAffected()
	if num > 0 {
		return true
	}
	return false
}

package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Comment struct {
	Id      int64     `json:"id"`
	Tid     int64     `json:"tid"`
	Uid     int64     `json:"uid"`
	Content string    `json:"content"`
	Like    int       `json:"like"`
	Instime time.Time `orm:"column(instime);auto_now_add;type(datetime)" json:"instime"`
}

type CommentNumTheme struct {
	Tid        int64
	CommentNum int
}

func NewComment() *Comment {
	return &Comment{}
}

func (this *Comment) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Comment)).Filter("status", 1)
}

// 获取主题的回复数
func (this *Comment) GetCommentNumByThemes(themes []*Theme) {
	if len(themes) > 0 {
		// 获取所有id
		var tids []int64
		for _, v := range themes {
			tids = append(tids, v.Id)
		}
		//this.Query().Filter("id__in", tids).GroupBy("tid").Count()
		var countnum []CommentNumTheme
		sql := "SELECT COUNT(1) AS common_num,tid FROM comment WHERE tid IN(?) GROUP BY tid"
		num, err := orm.NewOrm().Raw(sql, tids).QueryRows(&countnum)
		fmt.Printf("num=%d , err=%v\n", num, err)
		fmt.Println(countnum)
	}
}

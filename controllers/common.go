package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"vveq-api/models"
)

type CommonController struct {
	BaseController
}

type IndexParams struct {
	Cate  []*models.Cate
	Tag   []*models.Tag
	Theme []*models.Theme
}

// @Title Common
// @Param	body		body 	models.Common	true		"body for Common content"
// @Success 200 {int}
// @Failure 403 body is empty
// @router / [get]
func (this *CommonController) Get() {
	// 获取标签
	tag, err := models.NewTag().GetAll()
	if err != nil {
		beego.Error("get tag err: ", err)
	}
	fmt.Println(tag)
	// 获取分类
	cate, err := models.NewCate().GetAll()
	if err != nil {
		beego.Error("get tag err: ", err)
	}
	fmt.Println(cate)
	// 获取主题
	theme, err := models.NewTheme().Get(0)
	if err != nil {
		beego.Error("get tag err: ", err)
	}
	fmt.Println(theme)
	this.Data["json"] = IndexParams{
		Cate:  cate,
		Tag:   tag,
		Theme: theme,
	}
	this.ServeJSON()
}

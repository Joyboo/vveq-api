package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"vveq-api/models"
)

// Operations about Users
type ThemeController struct {
	BaseController
}

// @Title 获取主题(主页)
// @Description get Tag
// @Param	page	formData 	int	true		"当前第几页"
// @Success 200 {int} model.Tag
// @router / [get]
func (this *ThemeController) Get() {
	page, _ := this.GetInt("page", 0)
	themes, err := models.NewTheme().Get(page)
	if err != nil {
		beego.Error("Get err ", err)
		this.Data["json"] = ErrResponse{Status: 0}
		this.ServeJSON()
		return
	}

	// 获取回复数
	models.NewComment().GetCommentNumByThemes(themes)

	this.Data["json"] = Response{
		Status: 1,
		Data:   themes,
	}
	this.ServeJSON()
}

// @Title CreateTheme
// @Description create Theme
// @Param	body		body 	models.Theme	true		"body for Theme content"
// @Success 200 {int} model.Theme.Id
// @Failure 403 body is empty
// @router / [post]
func (this *ThemeController) Post() {
	var postParams models.Theme
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &postParams)
	if err != nil {
		beego.Error("theme post parse err: ", err)
		this.Data["json"] = ErrResponse{Status: 0}
		this.ServeJSON()
		return
	}
	id, err := postParams.Add()
	if err != nil || id <= 0 {
		beego.Error("theme post add err: ", err)
		this.Data["json"] = ErrResponse{Status: 0}
	} else {
		this.Data["json"] = Response{1, id}
	}
	this.ServeJSON()
}

package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"vveq-api/models"
)

// Operations about Tag
type TagController struct {
	BaseController
}

// @Title CreateTheme
// @Description create Theme
// @Param	body		body 	models.Tag	true		"body for Tag content"
// @Success 200 {int} model.Tag.Id
// @Failure 403 body is empty
// @router / [post]
func (this *TagController) Post() {
	var postParams models.Tag
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &postParams)
	if err != nil {
		this.Data["json"] = ErrResponse{Status: 0}
		this.ServeJSON()
		return
	}
	id, err := postParams.Add()
	if err != nil || id <= 0 {
		beego.Error("tag post id=", id, ", err:", err)
		this.Data["json"] = ErrResponse{Status: 0}
	} else {
		this.Data["json"] = Response{Status: 1}
	}
	this.ServeJSON()
}

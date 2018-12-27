package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"vveq-api/models"
)

// Operations about Cate
type CateController struct {
	BaseController
}

// @Title CreateTheme
// @Description create Theme
// @Param	body		body 	models.Cate	true		"body for Cate content"
// @Success 200 {int} model.Cate.Id
// @router / [post]
func (this *CateController) Post() {
	var postParams models.Cate
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &postParams)
	if err != nil {
		this.Data["json"] = map[string]int{"status": 0}
		this.ServeJSON()
		return
	}
	id, err := postParams.Add()
	if err != nil || id <= 0 {
		beego.Error("cate post id=", id, ", err:", err)
		this.Data["json"] = map[string]int{"status": 0}
	} else {
		this.Data["json"] = map[string]int{"status": 1}
	}
	this.ServeJSON()
}

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

// @Title 获取标签
// @Description get Tag
// @Param	type	formData 	int	true		"数据格式类型：1-普通格式，2-键值对格式 id:name"
// @Success 200 {int} model.Tag
// @router / [get]
func (this *TagController) Get() {
	types, err := this.GetInt("type", 1)
	if err != nil {
		beego.Error("Tag get err1: ", err)
		this.Data["json"] = ErrResponse{Status: 0}
		this.ServeJSON()
		return
	}

	tags, err := models.NewTag().GetAll()
	if err != nil {
		beego.Error("Tag get err2: ", err)
		this.Data["json"] = ErrResponse{Status: 0}
		this.ServeJSON()
		return
	}

	if types == 1 {
		this.Data["json"] = Response{
			Status: 1,
			Data:   tags,
		}
	} else if types == 2 {
		var data = make(map[int64]string)
		for _, v := range tags {
			data[v.Id] = v.Name
		}
		this.Data["json"] = Response{
			Status: 1,
			Data:   data,
		}
	}
	this.ServeJSON()
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

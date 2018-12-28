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

// @Title 获取分类
// @Description get Cate
// @Param	type	formData 	int	true		"数据格式类型：1-普通格式，2-键值对格式 id:name"
// @Success 200 {int} model.Cate
// @router / [get]
func (this *CateController) Get() {
	types, err := this.GetInt("type", 1)
	if err != nil {
		beego.Error("Cate get err1: ", err)
		this.Data["json"] = ErrResponse{Status: 0}
		this.ServeJSON()
		return
	}

	cates, err := models.NewCate().GetAll()
	if err != nil {
		beego.Error("Cate get err2: ", err)
		this.Data["json"] = ErrResponse{Status: 0}
		this.ServeJSON()
		return
	}

	if types == 1 {
		this.Data["json"] = Response{
			Status: 1,
			Data:   cates,
		}
	} else if types == 2 {
		var data = make(map[int64]string)
		for _, v := range cates {
			data[v.Id] = v.Name
		}
		this.Data["json"] = Response{
			Status: 1,
			Data:   data,
		}
	}
	this.ServeJSON()
}

// @Title 创建一个分类
// @Description create Cate
// @Param	body		body 	models.Cate	true		"body for Cate content"
// @Success 200 {int} model.Cate.Id
// @router / [post]
func (this *CateController) Post() {
	var postParams models.Cate
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &postParams)
	if err != nil {
		beego.Error("Cate post err: ", err)
		this.Data["json"] = ErrResponse{Status: 0}
		this.ServeJSON()
		return
	}
	id, err := postParams.Add()
	if err != nil || id <= 0 {
		beego.Error("cate post id=", id, ", err:", err)
		this.Data["json"] = ErrResponse{Status: 0}
	} else {
		this.Data["json"] = Response{Status: 1}
	}
	this.ServeJSON()
}

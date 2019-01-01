package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"vveq-api/models"
)

// Operations about Like
type LikeController struct {
	BaseController
}

// @Title CreateTheme
// @Description create Theme
// @Param	uid		body 	models.Like.Uid     true	 	"用户id"
// @Param	type	body 	models.Like.type	true		"对象类型 0-主题 1-评论"
// @Param	pid 	body 	models.Like.Pid 	true		"对象id"
// @Success 200 {int} model.Like.Id
// @Failure 403 body is empty
// @router / [post]
func (this *LikeController) Post() {
	var postParams models.Like
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &postParams)
	if err != nil {
		beego.Error("param err: ", err)
		this.Data["json"] = ErrResponse{Status: 0, Msg: "param error"}
		this.ServeJSON()
		return
	}

	bool := postParams.Add()
	if bool {
		this.Data["json"] = Response{Status: 1}
	} else {
		beego.Error("like post err: postParams=", postParams)
		this.Data["json"] = ErrResponse{Status: 0}
	}
	this.ServeJSON()
}

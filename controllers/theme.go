package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"vveq-api/models"
)

// Operations about Users
type ThemeController struct {
	BaseController
}

// @Title 根据id获取主题
// @Description get Tag
// @Param	id	formData 	int	true		"主题id"
// @Success 200 {int} model.Theme
// @router /:id [get]
func (this *ThemeController) Get() {
	id, err := this.GetInt64(":id", 0)
	fmt.Printf("param id=%d \n", id)
	if err != nil || id <= 0 {
		beego.Error("theme get param1: id=", id, ", err=", err)
		this.Data["json"] = Response{Status: 0}
		this.ServeJSON()
		return
	}
	theme, err := models.NewTheme().GetThemeById(id)
	if err != nil {
		beego.Error("theme get param2: id=", id)
		this.Data["json"] = Response{Status: 0}
		this.ServeJSON()
		return
	}
	this.Data["json"] = Response{
		Status: 1,
		Data:   theme.Content,
	}

	/*comments, err := models.NewComment().GetCommentsByTid(id)
	if err != nil {
		beego.Error("theme get param3: id=", id)
		this.Data["json"] = Response{Status: 0}
		this.ServeJSON()
		return
	}

	// 获取评论用户信息
	var uids []int64
	uids = append(uids, theme.Uid)
	for _, v := range comments {
		uids = append(uids, v.Uid)
	}
	users := models.NewUser().UsersByUserId(uids)

	var ud []map[int64]interface{}
	for _, v := range users {
		one := make(map[int64]interface{})
		one[v.Id] = map[string]interface{}{
			"id":       v.Id,
			"username": v.Username,
			"nickname": v.Nickname,
			"avatar":   v.Avatar,
		}
		ud = append(ud, one)
	}

	var data ResponseDataType
	data["id"] = theme.Id
	data["cid"] = theme.Cid
	data["title"] = theme.Title
	data["content"] = theme.Content
	data["uid"] = theme.Uid
	data["tagid"] = theme.Tagid
	data["click"] = theme.Click
	data["like"] = theme.Like
	data["instime"] = theme.Instime.Format(TimeFormart)
	data["comments"] = comments
	data["users"] = ud
	this.Data["json"] = Response{
		Status: 1,
		Data:   data,
	}*/
	this.ServeJSON()
}

// @Title 获取主题(主页)
// @Description get Tag
// @Param	page	formData 	int	true		"当前第几页"
// @Success 200 {int} model.Theme
// @router /getIndexTheme/:page [get]
func (this *ThemeController) GetIndexTheme() {
	page, _ := this.GetInt(":page", 0)
	themes, err := models.NewTheme().Gets(page)
	if err != nil {
		beego.Error("Get err ", err)
		this.Data["json"] = ErrResponse{Status: 0}
		this.ServeJSON()
		return
	}

	// 获取回复数
	models.NewComment().GetCommentNumByThemes(themes)

	// 只返回指定字段
	var one ResponseDataType
	var result []ResponseDataType
	var uids []int64
	for _, v := range themes {
		uids = append(uids, v.Id)
	}
	avatar := models.NewUser().UsersByUserId(uids)

	for _, v := range themes {
		one = ResponseDataType{
			"id":         v.Id,
			"cid":        v.Cid,
			"title":      v.Title,
			"uid":        v.Uid,
			"avatar":     avatar[v.Uid].Avatar,
			"username":   avatar[v.Uid].Username,
			"nickname":   avatar[v.Uid].Nickname,
			"tagid":      v.Tagid,
			"click":      v.Click,
			"like":       v.Like,
			"instime":    v.Instime.Format(TimeFormart),
			"commentNum": 0,
		}
		result = append(result, one)
	}
	this.Data["json"] = Response{
		Status: 1,
		Data:   result,
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

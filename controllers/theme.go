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

// @Title 根据id获取主题
// @Description get Tag
// @Param	id	formData 	int	true		"主题id"
// @Success 200 {int} model.Theme
// @router /:id [get]
func (this *ThemeController) Get() {
	id, err := this.GetInt64(":id", 0)
	if err != nil || id <= 0 {
		beego.Error("theme get param1: id=", id, ", err=", err)
		this.Data["json"] = ErrResponse{Status: 0, Msg: "param err"}
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
		Data:   theme,
	}
	this.ServeJSON()
}

// @Title 获取主题列表(主页)
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

// @Title 获取主题(含评论，用户信息)
// @Description get Tag
// @Param	id	formData 	int	true		"models.Theme.Id"
// @Success 200 {int} model.Theme
// @router /themeAndComment/:id [get]
/*func (this *ThemeController) ThemeAndComment() {
	id, err := this.GetInt64(":id", 0)
	if err != nil || id <= 0 {
		beego.Error("ThemeAndComment param err : id=", id, ", err=", err)
		this.Data["json"] = ErrResponse{Status: 0, Msg: "param err"}
		this.ServeJSON()
		return
	}
	// 获取theme
	theme, _ := models.NewTheme().GetThemeById(id)
	// 获取评论
	models.NewComment().GetCommentNumByThemes()
}*/

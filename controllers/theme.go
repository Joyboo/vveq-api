package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about Users
type ThemeController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} model.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *ThemeController) Post() {
	/*var postParams models.Theme
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &postParams)
	if err != nil {
		u.Data["json"] = map[string]int{"status": 0}
		u.ServeJSON()
		return
	}
	// 验证码校验
	if verifyResult := postParams.Verify.Compare(); !verifyResult {
		u.Data["json"] = map[string]int{"status": -1}
		u.ServeJSON()
		return
	}

	uid, err := postParams.From.Add()
	if err != nil || uid <= 0 {
		beego.Error("uid->", uid, ", 注册用户失败: ", err)
		u.Data["json"] = map[string]int{"status": 0}
	} else {
		u.Data["json"] = map[string]interface{}{
			"status": 1,
			"data": map[string]interface{}{
				"id":       uid,
				"username": postParams.From.Username,
				"nickname": postParams.From.Nickname,
				"email":    postParams.From.Email,
				"tel":      postParams.From.Tel,
				"avatar":   postParams.From.Avatar,
				"instime":  postParams.From.Instime,
			},
		}
	}*/
	u.ServeJSON()
}

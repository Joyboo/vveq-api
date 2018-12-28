package controllers

import (
	"encoding/json"
	"vveq-api/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	BaseController
}

type postPar struct {
	From   models.User
	Verify models.ConfigVerifyBody
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} model.User.Id
// @Failure 403 body is empty
// @router / [post]
func (this *UserController) Post() {
	var postParams postPar
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &postParams)
	if err != nil {
		this.Data["json"] = ErrResponse{Status: 0}
		this.ServeJSON()
		return
	}
	// 验证码校验
	if verifyResult := postParams.Verify.Compare(); !verifyResult {
		this.Data["json"] = ErrResponse{Status: -1}
		this.ServeJSON()
		return
	}

	uid, err := postParams.From.Add()
	if err != nil || uid <= 0 {
		beego.Error("uid->", uid, ", 注册用户失败: ", err, ", params: ", postParams.From)
		this.Data["json"] = ErrResponse{Status: 0}
	} else {
		this.Data["json"] = Response{
			Status: 1,
			Data: map[string]interface{}{
				"id":       uid,
				"username": postParams.From.Username,
				"nickname": postParams.From.Nickname,
				"email":    postParams.From.Email,
				"tel":      postParams.From.Tel,
				"avatar":   postParams.From.Avatar,
				"instime":  postParams.From.Instime,
			},
		}
	}
	this.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
/*func (u *UserController) GetAll() {
	users := user.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}*/

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (this *UserController) Get() {
	// todo 权限鉴定
	uid, _ := this.GetInt64(":uid")
	if uid > 0 {
		user, err := models.NewUser().GetUserById(uid)
		if err != nil {
			beego.Error("User Get err: ", err)
			this.Data["json"] = ErrResponse{Status: 0}
		} else {
			this.Data["json"] = Response{1, user}
		}
	}
	this.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	user.User	true		"body for user content"
// @Success 200 {object} user.User
// @Failure 403 :uid is not int
// @router /:uid [put]
/*func (u *UserController) Put() {
	uid, err := u.GetInt(":uid")
	if uid != 0 && err != nil {
		var userData user.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &userData)
		uu, err := user.UpdateUser(uid, &userData)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}*/

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
/*func (u *UserController) Delete() {
	uid, err := u.GetInt(":uid")
	if err != nil {
		u.Data["json"] = "delete error!"
		u.ServeJSON()
	}
	user.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}*/

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (this *UserController) Login() {
	var user models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	user, err := user.Login()
	if err != nil {
		beego.Error("登录出错了-->", err)
		this.Data["json"] = ErrResponse{Status: 0}
	} else {

		models.NewDau().Log(user.Id, this.Ctx.Request.RemoteAddr)

		this.Data["json"] = Response{
			Status: 1,
			Data: map[string]interface{}{
				"id":       user.Id,
				"username": user.Username,
				"nickname": user.Nickname,
				"email":    user.Email,
				"tel":      user.Tel,
				"avatar":   user.Avatar,
				"instime":  user.Instime,
			},
		}
	}
	this.ServeJSON()
}

// @Title UsernameIsExists
// @Description 用户名是否存在
// @Success 200 {string} logout success
// @router /usernameIsExists/:username [get]
func (this *UserController) UsernameIsExists() {
	username := this.GetString(":username")
	if username != "" {
		exist := models.NewUser().UsernameExist(username)
		if !exist {
			this.Data["json"] = Response{Status: 1}
		} else {
			this.Data["json"] = ErrResponse{Status: 0}
		}
	}
	this.ServeJSON()
}

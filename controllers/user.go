package controllers

import (
	"encoding/json"
	"vveq-api/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
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
func (u *UserController) Post() {
	var postParams postPar
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

	uid, err := postParams.From.AddUser()
	if err != nil || uid <= 0 {
		beego.Error("uid->", uid, ", 注册用户失败: ", err)
		u.Data["json"] = map[string]int{"status": 0}
	} else {
		postParams.From.Id = uid
		u.Data["json"] = map[string]interface{}{
			"status": 1,
			"data":   postParams.From,
		}
	}
	u.ServeJSON()
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
func (u *UserController) Get() {
	// todo 权限鉴定
	uid, _ := u.GetInt64(":uid")
	if uid > 0 {
		user, err := models.NewUser().GetUserById(uid)
		if err != nil {
			u.Data["json"] = map[string]interface{}{
				"status": 0,
				"data":   err.Error(),
			}
		} else {
			u.Data["json"] = map[string]interface{}{
				"status": 1,
				"data":   user,
			}
		}
	}
	u.ServeJSON()
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
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	user, err := models.NewUser().Login(username, password)
	if err != nil {
		u.Data["json"] = map[string]int{"status": 0}
	} else {
		u.Data["json"] = map[string]interface{}{
			"status": 1,
			"data":   user,
		}
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
/*func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}*/

// @Title UsernameIsExists
// @Description 用户名是否存在
// @Success 200 {string} logout success
// @router /usernameIsExists/:username [get]
func (u *UserController) UsernameIsExists() {
	username := u.GetString(":username")
	if username != "" {
		num, err := models.NewUser().GetUserByName(username)
		if err == nil && num <= 0 {
			u.Data["json"] = map[string]int{"status": 1}
		} else {
			u.Data["json"] = map[string]int{"status": 0}
		}
	}
	u.ServeJSON()
}

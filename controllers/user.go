package controllers

import (
	"encoding/json"
	"fmt"
	"vveq-api/models/user"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

type returnType map[string]interface{}

// @Title CreateUser
// @Description create users
// @Param	body		body 	user.User	true		"body for user content"
// @Success 200 {int} user.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user user.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	fmt.Println(user)
	//uid := user.AddUser(user)
	uid := 123
	u.Data["json"] = map[string]int{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} user.User
// @router / [get]
func (u *UserController) GetAll() {
	users := user.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} user.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	// todo 权限鉴定
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := user.GetUser(uid)
		if err != nil {
			u.Data["json"] = returnType{
				"status":  0,
				"message": err.Error(),
			}
		} else {
			u.Data["json"] = returnType{
				"status":  1,
				"message": "success",
				"data":    user,
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
func (u *UserController) Put() {
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
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid, err := u.GetInt(":uid")
	if err != nil {
		u.Data["json"] = "delete error!"
		u.ServeJSON()
	}
	user.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	if user.Login(username, password) {
		u.Data["json"] = "login success"
	} else {
		u.Data["json"] = "user not exist"
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}

// @Title UsernameIsExists
// @Description 用户名是否存在
// @Success 200 {string} logout success
// @router /usernameIsExists/:username [get]
func (u *UserController) UsernameIsExists() {
	username := u.GetString(":username")
	if username != "" {
		_, err := user.GetUser(username)
		if err != nil {
			u.Data["json"] = returnType{"status": 0}
		} else {
			u.Data["json"] = returnType{"status": 1}
		}
	}
	u.ServeJSON()
}

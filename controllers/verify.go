package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about verify
type VerifyController struct {
	beego.Controller
}

// @Title Get
// @Description 获取验证码
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (o *VerifyController) Get() {

}

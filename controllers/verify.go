package controllers

import (
	"github.com/astaxie/beego"
	"github.com/mojocn/base64Captcha"
	"log"
	"vveq-api/models/verify"
)

// Operations about verify
type VerifyController struct {
	beego.Controller
}

// @Title Get
// @Description 获取验证码
// @Param	postParameters verify.ConfigJsonBody		"verify.ConfigJsonBody"
// @Success 200 {object} verify.ConfigJsonBody
// @Failure 403 param is empty
// @router / [post]
func (o *VerifyController) Post() {
	//接收客户端发送来的请求参数
	var postParameters verify.ConfigJsonBody
	if err := o.ParseForm(&postParameters); err != nil {
		log.Println(err)
		o.Data["code"] = 0
		o.Data["message"] = "param error"
		o.ServeJSON()
		return
	}

	//创建base64图像验证码
	var config interface{}
	switch postParameters.CaptchaType {
	case "audio":
		config = postParameters.ConfigAudio
	case "character":
		config = postParameters.ConfigCharacter
	default:
		config = postParameters.ConfigDigit
	}
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	captchaId, digitCap := base64Captcha.GenerateCaptcha(postParameters.Id, config)
	base64Png := base64Captcha.CaptchaWriteToBase64Encoding(digitCap)

	//你也可以是用默认参数 生成图像验证码
	//base64Png := captcha.GenerateCaptchaPngBase64StringDefault(captchaId)

	// 响应
	//w.Header().Set("Content-Type", "application/json; charset=utf-8")
	o.Data["code"] = 1
	o.Data["data"] = base64Png
	o.Data["captchaId"] = captchaId
	o.Data["msg"] = "success"
	o.ServeJSON()
}

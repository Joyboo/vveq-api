package controllers

import (
	"encoding/json"
	"github.com/mojocn/base64Captcha"
	"log"
	"vveq-api/models"
)

// Operations about verify
type VerifyController struct {
	BaseController
}

// @Title GetCaptcha
// @Description 获取验证码
// @Param	postParameters models.ConfigVerifyBody		"models.ConfigVerifyBody"
// @Success 200 {object} models.ConfigVerifyBody
// @Failure 403 param is empty
// @router /getCaptcha [post]
func (this *VerifyController) GetCaptcha() {
	//接收客户端发送来的请求参数
	var postParameters models.ConfigVerifyBody
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &postParameters); err != nil {
		log.Println(err)
		this.Data["json"] = map[string]int{"status": 0}
		this.ServeJSON()
		return
	}

	//创建base64图像验证码
	config := postParameters.GetConfig()
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	captchaId, digitCap := base64Captcha.GenerateCaptcha(postParameters.Id, config)
	base64Png := base64Captcha.CaptchaWriteToBase64Encoding(digitCap)

	//你也可以是用默认参数 生成图像验证码
	//base64Png := captcha.GenerateCaptchaPngBase64StringDefault(captchaId)

	// 响应
	//w.Header().Set("Content-Type", "application/json; charset=utf-8")
	this.Data["json"] = map[string]interface{}{
		"status":    1,
		"data":      base64Png,
		"captchaId": captchaId,
	}
	this.ServeJSON()
}

// @Title VerifyCaptcha
// @Description 验证码校验
// @Param	postParameters models.ConfigVerifyBody		"models.ConfigVerifyBody"
// @Success 200 {object} models.ConfigVerifyBody
// @Failure 403 param is empty
// @router /verifyCaptcha [post]
func (this *VerifyController) VerifyCaptcha() {
	//接收客户端发送来的请求参数
	var postParameters models.ConfigVerifyBody
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &postParameters); err != nil {
		log.Println(err)
		this.Data["json"] = map[string]int{"status": 0}
		this.ServeJSON()
		return
	}
	//比较图像验证码
	if verifyResult := postParameters.Compare(); verifyResult {
		this.Data["json"] = map[string]int{"status": 1}
	} else {
		this.Data["json"] = map[string]int{"status": 0}
	}
	this.ServeJSON()
}

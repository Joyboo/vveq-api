package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/mojocn/base64Captcha"
	"log"
	"vveq-api/models/verify"
)

// Operations about verify
type VerifyController struct {
	beego.Controller
}

// @Title GetCaptcha
// @Description 获取验证码
// @Param	postParameters verify.ConfigJsonBody		"verify.ConfigJsonBody"
// @Success 200 {object} verify.ConfigJsonBody
// @Failure 403 param is empty
// @router /getCaptcha [post]
func (o *VerifyController) GetCaptcha() {
	//接收客户端发送来的请求参数
	var postParameters verify.ConfigJsonBody
	if err := json.Unmarshal(o.Ctx.Input.RequestBody, &postParameters); err != nil {
		log.Println(err)
		o.Data["json"] = map[string]interface{}{
			"status":  0,
			"message": "param error",
		}
		o.ServeJSON()
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
	o.Data["json"] = map[string]interface{}{
		"status":    1,
		"data":      base64Png,
		"captchaId": captchaId,
		"message":   "success",
	}
	o.ServeJSON()
}

// @Title VerifyCaptcha
// @Description 验证码校验
// @Param	postParameters verify.ConfigJsonBody		"verify.ConfigJsonBody"
// @Success 200 {object} verify.ConfigJsonBody
// @Failure 403 param is empty
// @router /verifyCaptcha [post]
func (o *VerifyController) VerifyCaptcha() {
	//接收客户端发送来的请求参数
	var postParameters verify.ConfigJsonBody
	if err := json.Unmarshal(o.Ctx.Input.RequestBody, &postParameters); err != nil {
		log.Println(err)
		o.Data["json"] = map[string]interface{}{
			"status":  0,
			"message": "参数有误",
		}
		o.ServeJSON()
	}
	//比较图像验证码
	verifyResult := base64Captcha.VerifyCaptcha(postParameters.Id, postParameters.VerifyValue)

	if verifyResult {
		o.Data["json"] = map[string]interface{}{
			"status":  1,
			"message": "验证通过",
		}
	} else {
		o.Data["json"] = map[string]interface{}{
			"status":  0,
			"message": "验证失败",
		}
	}
	o.ServeJSON()
}

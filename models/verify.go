package models

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
)

type ConfigVerifyBody struct {
	Id          string
	CaptchaType string
	VerifyValue string
}

func NewConfigVerifyBody() *ConfigVerifyBody {
	return &ConfigVerifyBody{}
}

func (this *ConfigVerifyBody) GetConfig() (config interface{}) {
	switch this.CaptchaType {
	case "audio":
		config = this.ConfigAudio()
	case "character":
		config = this.ConfigCharacter()
	default:
		config = this.ConfigDigit()
	}
	return
}

func (this *ConfigVerifyBody) ConfigAudio() base64Captcha.ConfigAudio {
	return base64Captcha.ConfigAudio{
		CaptchaLen: 6,
		Language:   "zh",
	}
}

func (this *ConfigVerifyBody) ConfigCharacter() base64Captcha.ConfigCharacter {
	return base64Captcha.ConfigCharacter{
		Height:             60,
		Width:              240,
		Mode:               2,
		ComplexOfNoiseText: 0,
		ComplexOfNoiseDot:  0,
		IsUseSimpleFont:    true,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         6,
	}
}

func (this *ConfigVerifyBody) ConfigDigit() base64Captcha.ConfigDigit {
	return base64Captcha.ConfigDigit{
		Height:     80,
		Width:      240,
		CaptchaLen: 5,
		MaxSkew:    0.7,
		DotCount:   80,
	}
}

// 验证码校验
func (this *ConfigVerifyBody) Compare() bool {
	return base64Captcha.VerifyCaptcha(this.Id, this.VerifyValue)
}

func (this *ConfigVerifyBody) demoCodeCaptchaCreate() {
	//config struct for digits
	//数字验证码配置
	var configD = base64Captcha.ConfigDigit{
		Height:     80,
		Width:      240,
		MaxSkew:    0.7,
		DotCount:   80,
		CaptchaLen: 5,
	}
	//config struct for audio
	//声音验证码配置
	var configA = base64Captcha.ConfigAudio{
		CaptchaLen: 6,
		Language:   "zh",
	}
	//config struct for Character
	//字符,公式,验证码配置
	var configC = base64Captcha.ConfigCharacter{
		Height: 60,
		Width:  240,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeNumber,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         6,
	}
	//创建声音验证码
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idKeyA, capA := base64Captcha.GenerateCaptcha("", configA)
	//以base64编码
	base64stringA := base64Captcha.CaptchaWriteToBase64Encoding(capA)
	//创建字符公式验证码.
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idKeyC, capC := base64Captcha.GenerateCaptcha("", configC)
	//以base64编码
	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)
	//创建数字验证码.
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idKeyD, capD := base64Captcha.GenerateCaptcha("", configD)
	//以base64编码
	base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)

	fmt.Println(idKeyA, base64stringA, "\n")
	fmt.Println(idKeyC, base64stringC, "\n")
	fmt.Println(idKeyD, base64stringD, "\n")
}

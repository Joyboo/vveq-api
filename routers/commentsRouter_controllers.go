package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["vveq-api/controllers:UserController"] = append(beego.GlobalControllerRouter["vveq-api/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["vveq-api/controllers:UserController"] = append(beego.GlobalControllerRouter["vveq-api/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["vveq-api/controllers:UserController"] = append(beego.GlobalControllerRouter["vveq-api/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["vveq-api/controllers:UserController"] = append(beego.GlobalControllerRouter["vveq-api/controllers:UserController"],
		beego.ControllerComments{
			Method:           "UsernameIsExists",
			Router:           `/usernameIsExists/:username`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["vveq-api/controllers:VerifyController"] = append(beego.GlobalControllerRouter["vveq-api/controllers:VerifyController"],
		beego.ControllerComments{
			Method:           "GetCaptcha",
			Router:           `/getCaptcha`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["vveq-api/controllers:VerifyController"] = append(beego.GlobalControllerRouter["vveq-api/controllers:VerifyController"],
		beego.ControllerComments{
			Method:           "VerifyCaptcha",
			Router:           `/verifyCaptcha`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

}

package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["vveq-api/controllers:CateController"] = append(beego.GlobalControllerRouter["vveq-api/controllers:CateController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["vveq-api/controllers:CateController"] = append(beego.GlobalControllerRouter["vveq-api/controllers:CateController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["vveq-api/controllers:LikeController"] = append(beego.GlobalControllerRouter["vveq-api/controllers:LikeController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["vveq-api/controllers:TagController"] = append(beego.GlobalControllerRouter["vveq-api/controllers:TagController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["vveq-api/controllers:TagController"] = append(beego.GlobalControllerRouter["vveq-api/controllers:TagController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["vveq-api/controllers:ThemeController"] = append(beego.GlobalControllerRouter["vveq-api/controllers:ThemeController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["vveq-api/controllers:ThemeController"] = append(beego.GlobalControllerRouter["vveq-api/controllers:ThemeController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["vveq-api/controllers:ThemeController"] = append(beego.GlobalControllerRouter["vveq-api/controllers:ThemeController"],
		beego.ControllerComments{
			Method:           "GetIndexTheme",
			Router:           `/getIndexTheme/:page`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["vveq-api/controllers:UserController"] = append(beego.GlobalControllerRouter["vveq-api/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["vveq-api/controllers:UserController"] = append(beego.GlobalControllerRouter["vveq-api/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["vveq-api/controllers:UserController"] = append(beego.GlobalControllerRouter["vveq-api/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["vveq-api/controllers:UserController"] = append(beego.GlobalControllerRouter["vveq-api/controllers:UserController"],
		beego.ControllerComments{
			Method:           "UsernameIsExists",
			Router:           `/usernameIsExists/:username`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["vveq-api/controllers:VerifyController"] = append(beego.GlobalControllerRouter["vveq-api/controllers:VerifyController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["vveq-api/controllers:VerifyController"] = append(beego.GlobalControllerRouter["vveq-api/controllers:VerifyController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}

// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"vveq-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	view := beego.NewNamespace("/api",
		beego.NSNamespace("/user", beego.NSInclude(&controllers.UserController{})),
		beego.NSNamespace("/verify", beego.NSInclude(&controllers.VerifyController{})),
		beego.NSNamespace("/theme", beego.NSInclude(&controllers.ThemeController{})),
		beego.NSNamespace("/common", beego.NSInclude(&controllers.CommonController{})),
		beego.NSNamespace("/cate", beego.NSInclude(&controllers.CateController{})),
		beego.NSNamespace("/tag", beego.NSInclude(&controllers.TagController{})),
	)

	admin := beego.NewNamespace("/admin",
		beego.NSNamespace("/cate", beego.NSInclude(&controllers.CateController{})),
		beego.NSNamespace("/tag", beego.NSInclude(&controllers.TagController{})),
	)

	beego.AddNamespace(view, admin)
}

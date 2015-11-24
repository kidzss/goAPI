// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"goAPI/controllers"

	"github.com/astaxie/beego"
	"os"
)

func init() {
	// ns := beego.NewNamespace("/v1",
	// 	beego.NSNamespace("/object",
	// 		beego.NSInclude(
	// 			&controllers.ObjectController{},
	// 		),
	// 	),
	// 	beego.NSNamespace("/user",
	// 		beego.NSInclude(
	// 			&controllers.UserController{},
	// 		),
	// 	),
	// )
	beego.Router("/login", &controllers.LoginContriller{})
	beego.Router("/reg", &controllers.RegisterContriller{})
	beego.Router("/upload", &controllers.UploadContriller{})

	os.Mkdir("images", os.ModePerm)
	//set image static url
	beego.SetStaticPath("/images", "images")
	// beego.AddNamespace(ns)
}

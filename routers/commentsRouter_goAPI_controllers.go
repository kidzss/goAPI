package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["goAPI/controllers:ObjectController"] = append(beego.GlobalControllerRouter["goAPI/controllers:ObjectController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["goAPI/controllers:ObjectController"] = append(beego.GlobalControllerRouter["goAPI/controllers:ObjectController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["goAPI/controllers:ObjectController"] = append(beego.GlobalControllerRouter["goAPI/controllers:ObjectController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["goAPI/controllers:ObjectController"] = append(beego.GlobalControllerRouter["goAPI/controllers:ObjectController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["goAPI/controllers:ObjectController"] = append(beego.GlobalControllerRouter["goAPI/controllers:ObjectController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["goAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["goAPI/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["goAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["goAPI/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["goAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["goAPI/controllers:UserController"],
		beego.ControllerComments{
			"Get",
			`/:uid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["goAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["goAPI/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:uid`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["goAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["goAPI/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:uid`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["goAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["goAPI/controllers:UserController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["goAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["goAPI/controllers:UserController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

}

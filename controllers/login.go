package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"goAPI/database"
)

type LoginContriller struct {
	beego.Controller
}

func (this *LoginContriller) Get() {
	fmt.Println("request:", this.Input().Get("name"))
	this.Data["json"] = "hello ray get "
	this.ServeJson()
}

func (this *LoginContriller) Post() {
	fmt.Println("request:", this.Input().Get("name"))
	name := this.Input().Get("name")
	password := this.Input().Get("password")
	if name == "ray" && password == "123" {
		this.Data["json"] = "login success"
		userinfo, err := database.GetUserInfo(1)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(userinfo)
	} else {
		this.Data["json"] = "login fail"
	}

	this.ServeJson()
}

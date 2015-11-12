package controllers

import (
	// "encoding/json"

	"github.com/astaxie/beego"
	"goAPI/models"
)

type RegisterContriller struct {
	beego.Controller
}

func (this *RegisterContriller) Post() {
	name := this.Input().Get("name")
	password := this.Input().Get("password")
	account := new(models.Accounts)
	account.Account = name
	account.Password = password

	_, status := models.InsertAccount(account)
	var m Message
	m.Status = status
	m.Msg = "register"
	this.Data["json"] = OToJson(m)
	this.ServeJson()
}

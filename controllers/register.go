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
	authKey := this.Input().Get("authKey")

	account := new(models.Accounts)
	baseData := new(models.BaseData)
	account.Account = name
	account.Password = password
	if len(authKey) == 0 {
		baseData.Status = 1
		baseData.Msg = "register"
	} else {

		_, status := models.InsertAccount(account)

		baseData.Status = status
		baseData.Msg = "register"
		baseData.Result = account
	}
	this.Data["json"] = OToJson(baseData)
	this.ServeJson()
}

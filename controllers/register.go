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
	account.Account = name
	account.Password = password
	var mapdata MapData
	if len(authKey) == 0 {
		mapdata.Accounts = nil
		this.Data["json"] = map[string]interface{}{"errorCode": 1, "msg": "fail", "result": mapdata}
	} else {

		_, status := models.InsertAccount(account)
		if status == 0 {
			mapdata.Accounts = account
			this.Data["json"] = map[string]interface{}{"errorCode": status, "msg": "register success", "result": mapdata}
		} else {
			this.Data["json"] = map[string]interface{}{"errorCode": 1, "msg": "this account is registed", "result": mapdata}
		}

	}

	this.ServeJson()
}

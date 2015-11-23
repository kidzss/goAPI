package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"goAPI/models"
)

type LoginContriller struct {
	beego.Controller
}

func (this *LoginContriller) Post() {
	name := this.Input().Get("name")
	password := this.Input().Get("password")
	authKey := this.Input().Get("authKey")

	fmt.Println("name:", name)
	fmt.Println("password:", password)
	fmt.Println("authKey:", authKey)

	account, err := models.QueryAccount(name, password)
	var mapdata MapData
	if err == nil && len(authKey) != 0 {
		mapdata.Accounts = *account
		this.Data["json"] = map[string]interface{}{"errorCode": 0, "msg": "success", "result": mapdata}

	} else {
		fmt.Println(err)
		this.Data["json"] = map[string]interface{}{"errorCode": 1, "msg": "fail" + err.Error(), "result": mapdata}
	}

	this.ServeJson()
}

func OToJson(o interface{}) string {
	body, err := json.Marshal(o)
	if err != nil {
		panic(err.Error())
	}
	return string(body)
}

type MapData struct {
	Accounts interface{}
}

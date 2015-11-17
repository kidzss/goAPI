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
	var baseData models.BaseData
	var mapdata MapData
	if err == nil && len(authKey) != 0 {
		baseData.Status = 0
		baseData.Msg = "login success"
		mapdata.Accounts = *account

	} else {
		fmt.Println(err)
		baseData.Status = 1
		baseData.Msg = "login fail"
	}
	baseData.Result = mapdata
	this.Data["json"] = OToJson(baseData)
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

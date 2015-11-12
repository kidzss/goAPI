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

	account, err := models.QueryAccount(name, password)
	var m Message
	if err == nil {
		m.Status = 0
		m.Msg = "login success"
		m.Acc = *account

	} else {
		fmt.Println(err)
		m.Status = 1
		m.Msg = "login fail"
	}
	this.Data["json"] = OToJson(m)
	this.ServeJson()
}

type Message struct {
	Msg    string
	Status int
	Acc    models.Accounts
}

func OToJson(o interface{}) string {
	body, err := json.Marshal(o)
	if err != nil {
		panic(err.Error())
	}
	return string(body)
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson" // for json get
	// "goAPI/database"
	_ "goAPI/docs"
	// "goAPI/models"
	_ "goAPI/routers"
)

type mapData struct {
	Name string
}

type BaseData struct {
	Msg    string
	Status float32
	Result mapData
}

func main() {
	// defer database.CloseDatabase()
	// if beego.RunMode == "dev" {
	// 	beego.DirectoryIndex = true
	// 	beego.StaticDir["/swagger"] = "swagger"
	// }
	// u1 := models.User{Name: "hahah", Age: 24, Sex: 1, Tel: "15971470520"}
	// account1 := models.AccountInfo{MemberId: "15071379972", Account: "15071379972", Password: "111111"}
	// models.Insert(&account1)

	// testJson()
	beego.Run()
}
func testJson() {
	var m1 mapData
	var b1 BaseData
	m1.Name = "ray"
	b1.Msg = "success"
	b1.Status = 0
	b1.Result = m1

	body, err := json.Marshal(b1)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("body:", string(body))

	js, err := simplejson.NewJson(body)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(js)
}

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

type MyData struct {
	Name  string
	Other float32
	Msgs  []Message
}
type Message struct {
	Title string
	Count float32
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
	var detail MyData

	detail.Name = "1"

	detail.Other = 2
	detail.Msgs = make([]Message, 3)

	for i := 0; i < len(detail.Msgs); i++ {
		detail.Msgs[i] = Message{"ray", 33}
	}
	// fmt.Println(detail.Msgs)
	body, err := json.Marshal(detail)
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

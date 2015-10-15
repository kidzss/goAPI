package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson" // for json get
)

type MyData struct {
	Name  string
	Other float32
	Msg   Message
}
type Message struct {
	Title string
	Count float32
}

func main() {
	var detail MyData

	detail.Name = "1"

	detail.Other = 2
	detail.Msg = Message{"ray", 33}
	fmt.Println(detail.Msg)
	body, err := json.Marshal(detail)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(body))

	js, err := simplejson.NewJson(body)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(js)
}

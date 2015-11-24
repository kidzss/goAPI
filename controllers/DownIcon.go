package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"net/url"
	"os"
)

type DownContriller struct {
	beego.Controller
}

func (this *DownContriller) Get() {
	filePath, err := url.QueryUnescape(this.Ctx.Request.RequestURI[1:])
	fmt.Println(this.Ctx.Request.RequestURI)
	if err != nil {
		return
	}
	f, err1 := os.Open(filePath)
	if err1 != nil {
		return
	}
	defer f.Close()
	_, err2 := io.Copy(this.Ctx.ResponseWriter, f)
	if err2 != nil {
		return
	}
}

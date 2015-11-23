package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	// "time"
)

type UploadContriller struct {
	beego.Controller
}

func (this *UploadContriller) Post() {
	// 获取上传文件 _ diao biao h
	f, h, err := this.GetFile("img")

	if err == nil {
		// 关闭文件
		defer f.Close()
	}
	name := this.GetString("fileName", "")
	if len(name) == 0 {
		defer f.Close()
		this.Data["json"] = map[string]interface{}{"success": 1, "message": "input file name"}
		this.ServeJson()
		return

	}
	if err != nil {
		// 获取错误则输出错误信息
		this.Data["json"] = map[string]interface{}{"success": 1, "message": err.Error()}
		this.ServeJson()
		return
	}
	// 获取当前年月
	//datePath := time.Now().Format("2006/01")
	// 设置保存目录
	dirPath := "./images" //+ datePath
	// 设置保存文件名
	FileName := name + h.Filename
	// fmt.Println("time.Now():", time.Now())
	// 将文件保存到服务器中
	err = this.SaveToFile("img", fmt.Sprintf("%s/%s", dirPath, FileName))
	if err != nil {
		// 出错则输出错误信息
		this.Data["json"] = map[string]interface{}{"success": 1, "message": err.Error()}
		this.ServeJson()
		return
	}
	this.Data["json"] = map[string]interface{}{"success": 0, "message": "upload success", "imgUrl": "http://114.215.194.193:8088" + fmt.Sprintf("%s/%s", "/images", FileName)}
	this.ServeJson()
}

package controllers

import (
	"fmt"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego"
	"os"
	"path"
	// "time"
	"path/filepath"
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
	name := this.GetString("userName", "")
	if len(name) == 0 {
		defer f.Close()
		this.Data["json"] = map[string]interface{}{"errorCode": 1, "msg": "input file name", "result": nil}
		this.ServeJson()
		return

	}
	if err != nil {
		// 获取错误则输出错误信息
		this.Data["json"] = map[string]interface{}{"errorCode": 1, "msg": err.Error(), "result": nil}
		this.ServeJson()
		return
	}
	// 设置保存目录
	dirPath := "./images/" + name + "/icon/"
	// 设置保存文件名
	FileName := dirPath + h.Filename
	// 将文件保存到服务器中
	// 检查文件 and delete
	getFilelist(dirPath)
	if !com.IsExist(FileName) {
		os.MkdirAll(path.Dir(FileName), os.ModePerm)
	}

	err = this.SaveToFile("img", FileName)
	if err != nil {
		// 出错则输出错误信息
		this.Data["json"] = map[string]interface{}{"errorCode": 1, "msg": err.Error(), "result": nil}
		this.ServeJson()
		return
	}
	var url imageUrl
	url.ImageUrl = "http://114.215.194.193:8088" + "/images/" + name + "/icon/" + h.Filename
	this.Data["json"] = map[string]interface{}{"errorCode": 0, "msg": "upload success", "result": url}
	this.ServeJson()
}

type imageUrl struct {
	ImageUrl string
}

func getFilelist(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		// println(path)
		os.Remove(path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

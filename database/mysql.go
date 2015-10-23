package database

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"goAPI/models"
)

func init() {

	orm.RegisterModel(new(models.User))                                           //注册表studentinfo 如果没有会自动创建
	orm.RegisterDriver("mysql", orm.DR_MySQL)                                     //注册mysql驱动
	orm.RegisterDataBase("default", "mysql", "root:w41615465@/user?charset=utf8") //设置conn中的数据库为默认使用数据库
	orm.RunSyncdb("default", false, false)                                        //后一个使用true会带上很多打印信息，数据库操作和建表操作的；第二个为true代表强制创建表
	orm.Debug = true                                                              //true 打印数据库操作日志信息
}

func Insert(user *models.User) error {
	dbObj := orm.NewOrm() //实例化数据库操作对象
	dbObj.Using("user")
	_, err := dbObj.Insert(user)
	return err
}

// func Query() interface{} {
// 	dbObj := orm.NewOrm() //实例化数据库操作对象
// 	dbObj.Using("user")
// }

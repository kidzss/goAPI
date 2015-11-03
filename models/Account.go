package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Accounts struct {
	Id       int
	MemberId string `orm:"size(20)"`
	Account  string `orm:"size(20)"`
	Password string `orm:"size(20)"`
}

func init() {

	orm.RegisterModel(new(Accounts))                                              //注册表studentinfo 如果没有会自动创建
	orm.RegisterDriver("mysql", orm.DR_MySQL)                                     //注册mysql驱动
	orm.RegisterDataBase("default", "mysql", "root:w41615465@/test?charset=utf8") //设置conn中的数据库为默认使用数据库
	orm.RunSyncdb("default", false, false)                                        //后一个使用true会带上很多打印信息，数据库操作和建表操作的；第二个为true代表强制创建表
	orm.Debug = true                                                              //true 打印数据库操作日志信息
}
func InsertAccount(account *Accounts) error {
	dbObj := orm.NewOrm() //实例化数据库操作对象
	// dbObj.Using("account")
	_, err := dbObj.Insert(account)
	return err
}

func QueryAccount(account, password string) (*Accounts, error) {
	dbObj := orm.NewOrm() //实例化数据库操作对象
	// dbObj.Using("account")
	CurrentAccount := new(Accounts) //记录读取，需要指定主键
	CurrentAccount.Password = password
	CurrentAccount.Account = account
	err := dbObj.QueryTable("accounts").Filter("account", account).Filter("password", password).One(CurrentAccount)
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		fmt.Printf("Returned Multi Rows Not One")
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		fmt.Printf("Not row found")
	}
	return CurrentAccount, err
}

// func DeleteAccount(account *AccountInfo) error {
// 	o := orm.NewOrm()
// 	o.Using()
// }

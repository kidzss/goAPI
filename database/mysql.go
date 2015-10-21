package database

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"goAPI/models"
)

var db *sql.DB
var err1 error

func init() {
	fmt.Println("init database")
	db, err1 = sql.Open("mysql", "root:w41615465@/user?charset=utf8")
	checkErr(err1)

}

func InsertIntoMysql(name string, sex int, tel string, age int) {

	stmt, err := db.Prepare("INSERT user_info SET name=?,sex=?,tel=?,age=?")
	checkErr(err)
	res, err := stmt.Exec(name, sex, tel, age)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

}

// func DeleteFromMysql(int id) {
// 	db, err := sql.Open("mysql", "root:w41615465@/user?charset=utf8")
// 	checkErr(err)
// 	//删除数据
// 	stmt, err = db.Prepare("delete from user_info where id=?")
// 	checkErr(err)

// 	res, err = stmt.Exec(id)
// 	checkErr(err)

// 	affect, err = res.RowsAffected()
// 	checkErr(err)

// 	fmt.Println(affect)

// 	db.Close()
// }

// func SelectDataFromSql(id int) string {
// 	o := orm.NewOrm()

// 	return ""
// }

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func CloseDatabase() {
	db.Close()
}

func GetUserInfo(id int) (*models.Userinfo, error) {
	o := orm.NewOrm()
	info := new(models.Userinfo)
	qs := o.QueryTable("user_info")
	err := qs.Filter("id", id).One(info)
	return info, err
}

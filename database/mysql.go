package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	fmt.Println("init database")

}

func InsertIntoMysql(name string, sex int, tel string, age int) {
	db, err := sql.Open("mysql", "root:w41615465@/user?charset=utf8")
	checkErr(err)
	stmt, err := db.Prepare("INSERT user_info SET name=?,sex=?,tel=?,age=?")
	checkErr(err)
	res, err := stmt.Exec(name, sex, tel, age)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	db.Close()
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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

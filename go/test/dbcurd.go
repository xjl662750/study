package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//"time"
)

func main() {
	//create()
	//update(1066)
	read()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func create() {
	db, err := sql.Open("mysql", "xjl:662750@/test?charset=utf8")
	checkErr(err)
	//插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)
	/*for i := 0; i < 100; i++ {
	  	 res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	  checkErr(err)
	  fmt.Println("shiming",i,"res",res)
	  }*/
	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	db.Close()

}
func update(id int) {
	db, err := sql.Open("mysql", "xjl:662750@/test?charset=utf8")
	checkErr(err)
	//更新数据
	stmt, err := db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err := stmt.Exec("www", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

}
func read() {
	db, err := sql.Open("mysql", "xjl:662750@/test?charset=utf8")
	checkErr(err)
	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	db.Close()
}
func delete(id int) {
	db, err := sql.Open("mysql", "xjl:662750@/test?charset=utf8")
	checkErr(err)
	//删除数据
	stmt, err := db.Prepare("delete from userinfo where uid=?")
	checkErr(err)
	//for id := 7; id < 1050; id++ {
	res, err := stmt.Exec(id)
	checkErr(err)
	//}
	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
	db.Close()
}

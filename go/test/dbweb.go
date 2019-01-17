package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//"time"
	"log"
	"net/http"
	"strings"
)

func main() {
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
	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

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

	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)
	//for id := 7; id < 1050; id++ {
	//res, err = stmt.Exec(id)
	checkErr(err)
	//}
	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
	http.HandleFunc("/", sayhelloName)        //设置访问的路由
	err1 := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err1 != nil {
		log.Fatal("ListenAndServe: ", err1)
	}
	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Wrold!") //这个写入到w的是输出到客户端的
	fmt.Fprintf(w, "sadadad")
}

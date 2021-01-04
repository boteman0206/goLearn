package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id       int
	name     string
	password string
}

func main() {
	// 数据库的地址
	dsn := "root:1234@tcp(127.0.0.1:3306)/login"

	db, e := sql.Open("mysql", dsn)
	if e != nil {
		fmt.Println("数据库连接失败！", e)
	}

	db.SetMaxOpenConns(10) // 设置最大的连接数10
	//ping := db.Ping()

	fmt.Println("db:", db)

	sql1 := "select * from user where id = ?"

	row := db.QueryRow(sql1, 1)

	/**
	方式一：
	var user User
	row.Scan(&user.id, &user.name, &user.password)

	fmt.Println("i: ", user)

	*/

	/**
	方式二：
	*/
	var id, name, password string
	row.Scan(&id, &name, &password)
	fmt.Println("map : ", id, " ", name, " ", password)

	rows, i := db.Query("select * from user ")
	if i != nil {
		fmt.Println("iiiiiii", i)
	}
	for rows.Next() {
		rows.Scan(&id, &name, &password)
		fmt.Println("查询多行： ", id, " ", name, " ", password)
	}

}

package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User1 struct {
	Id       int    `sql:"id"`
	Name     string `sql:"name"`
	Password string `sql:"password"`
}

func main() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/login"

	db12, e1 := sqlx.Connect("mysql", dsn)
	if e1 != nil {
		fmt.Println("数据库连接失败！", e1)
	}
	fmt.Println("db12 : ", db12)

	/**
	查询单个sql
	*/
	sqlr1 := "select id, name, password from user where id =?"
	var us User1
	db12.Get(&us, sqlr1, 4)

	fmt.Println("us12 : ", us)

	/**
	查询list sql
	*/
	sqlr2 := "select id, name, password from user"
	users := make([]User1, 0, 10)
	db12.Select(&users, sqlr2)
	fmt.Println("users list : ", users)

	/**
	sql的注入
	*/

}

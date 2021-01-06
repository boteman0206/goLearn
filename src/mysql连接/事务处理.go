package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 数据库的地址
	dsn := "root:1234@tcp(127.0.0.1:3306)/login"

	db, e := sql.Open("mysql", dsn)
	if e != nil {
		fmt.Println("数据库连接失败！", e)
	}

	/**
	todo 开启事务
	*/
	tx, i := db.Begin()
	if i != nil {
		fmt.Println("开始事务失败")
		return
	}
	sql1 := "update user set name = 'k112' where id = ?"
	sql2 := "update user set name = name +1 where id = ?" // 错误的sql

	result1, i21 := tx.Exec(sql1, 4)

	if i21 != nil {
		fmt.Println("result1 :", result1)
		tx.Rollback()

	}
	result2, i22 := tx.Exec(sql2, 4)
	if i22 != nil {
		fmt.Println("result2 :", result2)
		tx.Rollback()

	}
	commit := tx.Commit()

	fmt.Println("commit : ", commit)

}

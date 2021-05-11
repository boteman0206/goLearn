package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Member struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {
	//database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/demo")
	if err != nil {
		fmt.Println("conn error !")
		return
	}
	Db = database
	fmt.Println("successful conn db :", Db)

	//defer database.Close() // todo  注意这行代码要写在上面err判断的下面

}

func main() {

	// todo 1： insert操作数据库
	//exec, err := Db.Exec("insert into member(username, sex, email)values (?, ?, ?)", "student03", "女", "111@qq.com")
	//if err != nil {
	//	fmt.Println("conn error :", err.Error())
	//	return
	//}
	//fmt.Println("exec 是： ", exec)
	//id, err1 := exec.LastInsertId()
	//affected, err2 := exec.RowsAffected()
	//fmt.Println(id, " 影响的行数：" , affected, err1, err2) // 3  影响的行数： 1 <nil> <nil>

	// todo 2 select 多行操作数据库
	var member []Member
	err1 := Db.Select(&member, "select * from member where sex = ? ", "女")
	if err1 != nil {
		fmt.Println("select error : ", err1.Error())
		return
	}
	for i, mem := range member {
		fmt.Println("i : ", i, " member ", mem)
	}

	fmt.Println("-------------查询封装map数据------------------")

	selectdata := fmt.Sprintf("select * from member where user_id = %d ", 4)
	query, err1 := Db.Query(selectdata)
	fmt.Println("map error : ", err1)
	fmt.Println(query.Columns())
	map1 := make(map[string]interface{})
	for query.Next() {
		err1 := query.Scan(map1)
		fmt.Scan("err1 : ", err1)
	}
	fmt.Println("map1 data :", map1)

	//todo 查询一行数据

	sqlStr := "SELECT * FROM member WHERE user_id = ?"

	var member1 Member
	if err := Db.Get(&member1, sqlStr, 4); err != nil {
		fmt.Printf("get data failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d, name:%s, age:%d\n", member1.UserId, member1.Username, member1.Sex, member1.Email)
	fmt.Println()
	fmt.Println("测试tag标签不加有啥用。。。。")
	row := Db.QueryRow("select * from member where user_id = ?", 4)
	// todo  必须要这样进行scan
	fmt.Println(row.Scan(&member1.UserId, &member1.Username, &member1.Sex, &member1.Email))
	if err1 != nil {
		fmt.Println("查询一行失败： ", err1.Error())
	}
	fmt.Println("member1 : ", member1)

	// todo 查询count函数
	var num int
	str1 := "select count(*) num  from member"
	queryRow := Db.QueryRow(str1)
	err1 = queryRow.Scan(&num)
	fmt.Println(err1, num)

}

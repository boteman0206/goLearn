package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

type UserData struct {
	Id        int64
	Name      string
	CreatedAt time.Time `json:"created_at"`
}

var Engine1 *xorm.Engine

func init() {
	var err error
	Engine1, err = xorm.NewEngine("mysql", "root:123456@(localhost:3306)/demo?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println("err2 : ", err.Error())
		return
	}
	Engine1.ShowSQL(true)
}

func main() {
	//err := Engine1.Sync2(&UserData{})
	//if err!= nil {
	//	fmt.Println("sync data error！")
	//	return
	//}

	data := UserData{
		Name:      "user2",
		Id:        2,
		CreatedAt: time.Now(), // 自动创建时间
	}
	insert, err := Engine1.Insert(&data)
	fmt.Println(" err: ", err, "insert : ", insert)

}

/**

todo
	location, err = time.LoadLocation("Asia/Shanghai")
	engine.TZLocation = location
	root:123456@(localhost:3306)/demo?charset=utf8&parseTime=true&loc=Local
	加上&loc=Local是可以更新正确的时间
*/

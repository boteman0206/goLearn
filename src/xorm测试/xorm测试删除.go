package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:pw123456@(localhost:3306)/demo?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println("err2 : ", err.Error())
		return
	}
	engine.ShowSQL(true)
}

type TEnumExample struct {
	Color string `xorm:"ENUM('blue','green','red')"`
	Id    int    `xorm:"not null pk autoincr INT(11)"`
	Name  string `xorm:"VARCHAR(10)"`
}

func main() {

	fmt.Println(engine)

	example := TEnumExample{}

	b, e := engine.Id(1).Get(&example)
	fmt.Println(b, e)

	example1 := TEnumExample{}
	i, e := engine.In("id", 1, 2, 3).Delete(&example1) // 直接就删除了 DELETE FROM `t_enum_example` WHERE `id` IN (?,?,?) []interface {}{1, 2, 3}
	fmt.Println("i: ", i, e)
}

package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

var Engine *xorm.Engine

func init() {
	var err error
	Engine, err = xorm.NewEngine("mysql", "root:123456@(localhost:3306)/demo?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println("err2 : ", err.Error())
		return
	}
	Engine.ShowSQL(true)
}

type UserFlag struct {
	Id             int       `xorm:"not null pk autoincr comment('主键') INT(11)"`
	UserId         string    `xorm:"not null comment('用户ID--企微添加人的id') VARCHAR(36)"`
	ExternalUserid string    `xorm:"default NULL comment('外部客户的企微id') VARCHAR(50)"`
	FlagId         int       `xorm:"default NULL comment('标签ID') INT(11)"`
	TagId          string    `xorm:"default NULL comment('外部客户的企微id') VARCHAR(50)"`
	CreateDate     time.Time `xorm:"creat_date created comment('创建时间') DATETIME"`      // todo 指定这个created会自动创建不会覆盖
	CreateTime     time.Time `xorm:"create_time default current_timestamp() DATETIME"` // 不传的时候为啥无效还不清楚
}

func main() {
	//fmt.Println(Engine)
	//err := Engine.Sync2(new(UserFlag))
	//fmt.Println("同步数据库到mysql失败 ： ", err.Error())
	flag := UserFlag{
		Id:             0,
		UserId:         "i11213",
		ExternalUserid: "t1o1",
		FlagId:         3,
		TagId:          "sdd",
		CreateDate:     time.Now().Add(3 * time.Hour),
		CreateTime:     time.Now().Add(3 * time.Hour), //
	}

	//add := time.Now().Add(3 * time.Hour).Format("2006-01-02 15:04:05")
	//fmt.Println(add)
	one, err := Engine.InsertOne(&flag)
	fmt.Println("插入数据： ", one, err)

	//one, err := Engine.Cols("id, user_id, create_time, tag_id").InsertOne(&flag)
	//fmt.Println("cols 插入： ", one , err)

}

//todo 指定了created的xorm标签，修改了时间不起作用
// todo 指定默认值，可以修改时间

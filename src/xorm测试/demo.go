package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

var engine *xorm.Engine

type User struct {
	Id   int64
	Name string `xorm:"varchar(25) notnull unique 'usr_name' comment('姓名')"`
}

func getUser(c echo.Context) error {
	fmt.Println("create user ....")
	param := c.QueryParam("id")
	fmt.Println("获取的参数： ", param)
	query, err09 := engine.Table("user_info").Query("select * from member")
	fmt.Println("err09 : ", err09)
	for i, v := range query {
		fmt.Println("i : ", i, "v: ", v)
	}
	fmt.Println("-------------------------")
	var user Member
	user.UserId = 11
	engine.Get(&user)
	fmt.Println("user : ", user)

	user1 := new(Member)
	user1.Sex = "女"
	get, err09 := engine.Get(user1)
	fmt.Println(err09, " has ", get)

	/**
	todo 查询多条数据使用Find方法，Find方法的第一个参数为slice的指针或Map指针，
	 	即为查询后返回的结果，第二个参数可选，为查询的条件struct的指针。
	*/
	everyone := make([]Member, 0)
	err12 := engine.Find(&everyone)

	pEveryOne := make([]*Member, 0)
	err121 := engine.Find(&pEveryOne)
	fmt.Println(err12, " --- ", err121)

	fmt.Println("===========================")
	// todo 使用map作为查询
	/**
	todo 使用 map[int]user 查询时候必须要加上  xorm:"not null pk autoincr comment('自增主键') INT(11)"
	*/
	Persons := make(map[int64]Person, 0)
	err09 = engine.Find(&Persons)
	fmt.Println("map查询的error ： ", err09)
	users := make([]Person, 0)
	err := engine.Table("person").Where("name like ?", "%a%").Find(&users)
	fmt.Println(err)

	// todo ID查询
	/**
	todo 使用ID查询时候必须要加上  xorm:"not null pk autoincr comment('自增主键') INT(11)"
	*/
	var per Person
	i, err09 := engine.ID(1).Get(&per)
	fmt.Println("delect i : ", i, " error : ", err09)
	return c.JSON(http.StatusCreated, Persons)

}

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:123456@(localhost:3306)/demo?charset=utf8")
	if err != nil {
		fmt.Println("err2 : ", err.Error())
		return
	}
	engine.ShowSQL(true)
	//todo 同步数据库
	//err = engine.Sync2(new(UserInfo))
	if err != nil {
		fmt.Println("同步数据库异常： ", err.Error())
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	fmt.Println("echo run ...")
	// Routes
	e.POST("/users/get", getUser)
	//e.GET("/users/:id", getUser)
	//e.PUT("/users/:id", updateUser)
	//e.DELETE("/users/:id", deleteUser)
	// Start server
	e.Logger.Fatal(e.Start(":9003"))

}

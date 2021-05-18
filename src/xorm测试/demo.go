package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/cast"
	"net/http"
	"vo/src/xorm测试/models"
	"xorm.io/builder"
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
	var user models.Member
	user.UserId = 11
	engine.Get(&user)
	fmt.Println("user : ", user)

	user1 := new(models.Member)
	user1.Sex = "女"
	get, err09 := engine.Get(user1)
	fmt.Println(err09, " has ", get)

	/**
	todo 查询多条数据使用Find方法，Find方法的第一个参数为slice的指针或Map指针，
	 	即为查询后返回的结果，第二个参数可选，为查询的条件struct的指针。
	*/
	everyone := make([]models.Member, 0)
	err12 := engine.Find(&everyone)

	pEveryOne := make([]*models.Member, 0)
	err121 := engine.Find(&pEveryOne)
	fmt.Println(err12, " --- ", err121)

	fmt.Println("===========================")
	// todo 使用map作为查询
	/**
	todo 使用 map[int]user 查询时候必须要加上  xorm:"not null pk autoincr comment('自增主键') INT(11)"
	*/
	Persons := make(map[int64]models.Person, 0)
	err09 = engine.Find(&Persons)
	fmt.Println("map查询的error ： ", err09)
	users := make([]models.Person, 0)
	err := engine.Table("person").Where("name like ?", "%a%").Find(&users)
	fmt.Println(err)

	// todo ID主键查询查询  务必注意点：使用主键ID和find配合无效
	/**
	todo 使用ID查询时候必须要加上  xorm:"not null pk autoincr comment('自增主键') INT(11)"
	*/
	var per models.Person
	i, err09 := engine.ID(1).Get(&per)
	fmt.Println("delect i : ", i, " error : ", err09)
	var mem1 models.Member
	b, err09 := engine.Id(3).Get(&mem1)
	fmt.Println(b, " id 查询 ： ", err09)
	//var mem2 models.Member
	//b1, err091 := engine.ID(3).Get(&mem2)  // todo 使用主键必须配合GET使用
	//fmt.Println(b1, " id 查询 ： ", err091)

	/**
	todo 务必注意点：使用主键ID和find配合无效
		var mem2 []models.Member
		err09 = engine.ID(3).Find(&mem2)
		fmt.Println("使用主键ID和find配合无效： ", err09)
	*/

	return c.JSON(http.StatusCreated, Persons)

}

func deleteUser(c echo.Context) error {
	param := c.QueryParam("id")
	fmt.Println("param : ", param)

	/**
	TODO delete删除用户
	*/
	var per models.Person
	engine.ID(1).Delete(&per)

	results, err := engine.QueryInterface(builder.Select("*").From("member"))
	fmt.Println(err, " result : ", results)
	for i, res := range results {
		fmt.Printf("%T !\n", res)
		fmt.Println("i : ", i, " res : ", res["email"])
	}

	//todo or条件查询
	var mems = make([]models.Member, 0)
	err = engine.Table("member").Or("user_id = ?", 11).Or("user_id= ?", 3).Find(&mems)
	fmt.Println("mems or 查询: ", err, mems)

	fmt.Println("err_new : ", mems)
	return c.JSON(http.StatusOK, mems)
}

func getUser1(c echo.Context) error {
	mems := make([]models.Member, 0)

	err := engine.Desc("user_id").Find(&mems)
	fmt.Println("error : ", err)

	mems = make([]models.Member, 0)
	engine.In("user_Id", 1, 2, 3).Find(&mems)

	// 查询特定的字段
	var user models.Member
	engine.Cols("username", "sex").Get(&user)
	fmt.Println("只有部分字段显示的： ", user)

	// todo 更新部分字段
	user.Email = "12300@bb.com"
	user.Username = "更改之后"
	user.Sex = "其他"
	engine.Cols("username", "email").Where("user_id = ?", 4).Update(&user)

	// limit字段
	mems = make([]models.Member, 0)
	engine.Desc("user_id").Limit(2).Find(&mems)
	fmt.Println("limit mems : ", mems)

	// todo 直接执行sql语句解析到结构体种
	mems = make([]models.Member, 0)
	type groupData struct {
		Member models.Member `xorm:"extends"`
		// todo 如果表里面有相同的字段则无法展示，可以在字段前面加上字段去区分
		Person models.Person `xorm:"extends"`
	}
	var joindata = make([]groupData, 0)
	sql := engine.SQL("SELECT * FROM member a join person.go b on a.user_id=b.id").Find(&joindata)
	fmt.Println("sql error : ", sql)

	// todo findAndCount函数使用
	count, err := engine.Where(builder.Eq{"user_id": 12}).FindAndCount(&mems)
	fmt.Println("findAndCount 函数的使用： ", count, " err : ", err)
	return c.JSON(http.StatusOK, joindata)
}

func testManyChoose(c echo.Context) error {
	where := engine.Table("member").Where("1=1")
	if 1 == 1 {
		where.And("user_id = ? ", 3)
	}
	if 2 == 2 {
		where.And("sex = ?", "男")
	}
	if 3 == 3 {
		where.And("username like ?", "%student%")
	}
	var mem []models.Member
	err := where.Find(&mem)
	fmt.Println("获取的用户信息： ", mem)

	//todo 使用链式的builder调用
	filter := builder.Select("*")

	if 1 == 1 {
		filter.Or(builder.Eq{"user_id": 3}).Or(builder.Eq{"user_id": 4})
	}
	if 2 == 2 {
		filter.And(builder.Eq{"sex": "男"})
	}
	if 3 == 3 {
		filter.And(builder.Like{"username", "student"})
	}
	sql, i, err := filter.From("member").ToSQL()
	fmt.Println("filter : ", sql, i, err)
	var mem1 []models.Member
	err = engine.SQL(sql, i...).Find(&mem1)
	fmt.Println("filter 构造过滤： ", mem1, err)
	/**
	// todo 这样是链式的调用复杂的链式调用
	*/
	mems := make([]models.Member, 0)
	//err_new := engine.Table("member").And(builder.Or(builder.Eq{"user_id": 3})).Find(&mems)
	// 查询user_id=3或者4, 并且性别为女的用户
	engine.Table("member").
		Or(builder.Eq{"user_id": 3}, builder.Eq{"user_id": 4}).
		And(builder.Eq{"sex": "女"}).
		Find(&mems)

	toSQL, i2, err := builder.ToSQL(builder.And(builder.Eq{"a": 1}, builder.Like{"b", "c"}, builder.Neq{"d": 2}))
	fmt.Println(toSQL, i2, err)

	return err
}

func filterSql(c echo.Context) error {
	//方式三使用builder构造复杂查询
	queryString := c.QueryString
	name := c.QueryParam("name")
	id := c.QueryParam("id")
	fmt.Println("id : ", id, "name : ", name, "queryString : ", queryString())
	toInt := cast.ToInt(id)
	from := builder.Select("*").From("member")
	if toInt > 0 {
		from.And(builder.Eq{"user_id": id})
	}
	if len(name) > 0 {
		from.And(builder.Like{"username", name})
	}
	sql, i, err := from.ToSQL()

	fmt.Println("拼接的完整的sql和参数： ", sql, i, err)

	var mem []models.Member
	err = engine.SQL(sql, i...).Find(&mem)
	fmt.Println("error : ", err)
	return c.JSON(http.StatusOK, mem)
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
	e.GET("/users/get1", getUser1)
	//e.GET("/users/:id", getUser)
	//e.PUT("/users/:id", updateUser)
	e.GET("/users/id", deleteUser)
	e.GET("/users/testManyChoose", testManyChoose)
	e.GET("/users/filter", filterSql)
	// Start server
	e.Logger.Fatal(e.Start(":9003"))

}

package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"strconv"
)

type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	users = map[int]*user{}
	seq   = 1
)

var users1 map[string]string

//----------
// Handlers
//----------

func createUser(c echo.Context) error {
	fmt.Println("create user ....")
	u := &user{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}

	fmt.Println("初始化之前 ： ", users)
	users = make(map[int]*user)
	fmt.Println("初始化之后： ", users)
	users[u.ID] = u
	seq++
	for i, v := range users {
		fmt.Println("i : ", i, "v : ", v)
	}
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

func updateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}

//todo 测试请求头的参数

func getHeader(c echo.Context) error {

	// 只能是请求上下文set数据才能获取
	app_id := c.Get("app_id")

	s := c.Request().Header.Get("app_id")
	fmt.Println("s :  ", s)
	// 设置
	c.Set("app_id", "mytestAppid")
	app_id1 := c.Get("app_id")
	//app_id_string := app_id.(string)
	fmt.Println("获取的请求头的参数appId： ", app_id, app_id1)

	return c.JSON(http.StatusOK, "ok")
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())

	e.Use(middleware.Recover())

	fmt.Println("echo run ...")
	// Routes
	e.POST("/users", createUser)
	e.GET("/getHeader", getHeader)
	//e.GET("/users/:id", getUser)
	//e.PUT("/users/:id", updateUser)
	//e.DELETE("/users/:id", deleteUser)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))

}

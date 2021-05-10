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

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	fmt.Println("echo run ...")
	// Routes
	e.POST("/users", createUser)
	//e.GET("/users/:id", getUser)
	//e.PUT("/users/:id", updateUser)
	//e.DELETE("/users/:id", deleteUser)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

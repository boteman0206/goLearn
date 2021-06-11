package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func demoFunc(c echo.Context) error {

	go test1()
	fmt.Println(" test  go run  ")
	return c.JSON(http.StatusOK, "没有报错信息")
}

func test1() {
	fmt.Println("阻塞的chann 。。。。test1 ")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获阻塞的异常、、、、、", err)
		}
	}()
	panic("ppppp=-")

	ints := make(chan int)
	ints <- 19
	ints <- 8

}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())

	//e.Use(middleware.Recover())

	fmt.Println("echo run ...")
	// Routes
	e.GET("/users", demoFunc)

	e.Logger.Fatal(e.Start(":1323"))

}

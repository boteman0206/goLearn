package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func main() {

	e := echo.New()

	// todo 前置的和路由级别的还是会走， group和router的不会走了
	e.GET("/first", get)
	// 注册路由之前的中间件
	e.Pre(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			fmt.Println("前置的路由级别设置。。。。")
			return handlerFunc(context)
		}
	})

	// 路由级别的中间件
	e.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			fmt.Println("全局的中间件，所有的路由都会走到这里")
			return handlerFunc(c)
		}

	})

	group := e.Group("/my", func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			fmt.Println("group 级别的中间件")
			return handlerFunc(context)
		}
	})

	group.GET("/get", get, func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			fmt.Println("路由级别的中间件。。。。")
			return handlerFunc(context)
		}
	})

	e.Start(":9090")

}

func get(c echo.Context) error {
	return c.JSON(http.StatusOK, "get 方法。。。")
}

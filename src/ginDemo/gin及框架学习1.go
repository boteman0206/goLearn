package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "go语言中文文档www.topgoer.com",
		})
	})
	r.Run(":9990") // listen and serve on 0.0.0.0:8080
}

/**
r.Run() 的源码：

	如果没有获取到addr地址的话：先获取env环境的PORT变量，没有获取到的话，设置默认变量为8080
	否则直接返回设置的addr


	然后走http的流程


*/

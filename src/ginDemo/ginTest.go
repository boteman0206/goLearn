package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/event"
	"github.com/sirupsen/logrus"
	kit "github.com/tricobbler/rp-kit"
	"io"
	"os"
)

//添加事件触发的操作

func init() {
	// Register event listener
	event.On("evt1", event.ListenerFunc(func(e event.Event) error {
		fmt.Printf("handle event: %s , %s, %s\n", e.Name(), "second", kit.JsonEncode(e.Data()))
		return nil
	}), event.Normal)

	// Register multiple listeners
	event.On("evt1", event.ListenerFunc(func(e event.Event) error {
		fmt.Printf("handle event: %s, %s, %s\n", e.Name(), " first", kit.JsonEncode(e.Data()))
		return nil
	}), event.High)

}

// 添加日志打印功能
func init() {

	writer2 := os.Stdout
	writer3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		logrus.Fatalf("create file log.txt failed: %v", err)
	}

	logrus.SetOutput(io.MultiWriter(writer2, writer3))
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("info msg")

}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {

		logrus.Info("ping runing....")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 创建记录日志的文件
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	//
	//// 如果需要将日志同时写入文件和控制台，请使用以下代码
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)

	router.Run(":9010") // listen and serve on 0.0.0.0:8080
}

type Person struct {
	Age  int    `json:"age" form:"age" binding:"required,gt=10"`
	Name string `json:"name" form:"name" binding:"NotNullAndAdmin"`
}

func getting(c *gin.Context) {
	logrus.Info("geting runing ...")

	// ... ...

	// Trigger event
	// Note: The second listener has a higher priority, so it will be executed first.
	// get操作的时候触发
	event.MustFire("evt1", event.M{"arg0": "val0", "arg1": "val1"})

	c.JSON(200, "getting")
}

func posting(c *gin.Context) {

	vo := Person{}
	err := c.Bind(&vo)
	fmt.Println(" err: ", err, " vo: ", vo)
	if err != nil {
		return
	}
	logrus.Info("posting runing .......")
	c.JSON(200, "postting")
}

func putting(c *gin.Context) {

	logrus.Info("puting runing .........")
	c.JSON(200, "putting")
}

func deleting(c *gin.Context) {

	logrus.Info("deleting run ......")
	c.JSON(200, "deleteing")
}

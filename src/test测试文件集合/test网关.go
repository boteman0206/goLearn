package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	zkOt "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	zkHttp "github.com/openzipkin/zipkin-go/reporter/http"
	"log"
)

func main() {

	engine := gin.Default()

	reporter := zkHttp.NewReporter("http://localhost:9411/api/v2/spans")
	defer reporter.Close()
	endpoint, err := zipkin.NewEndpoint("main3", "localhost:9010")
	if err != nil {
		log.Fatalf("unable to create local endpoint: %+v\n", err)
	}
	nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		log.Fatalf("unable to create tracer: %+v\n", err)
	}
	zkTracer := zkOt.Wrap(nativeTracer)
	opentracing.SetGlobalTracer(zkTracer)

	// 第三步: 添加一个 middleWare, 为每一个请求添加span
	engine.Use(func(c *gin.Context) {
		span := zkTracer.StartSpan(c.FullPath())
		defer span.Finish()
		c.Next()
	})

	engine.GET("/api/test/test01", func(context *gin.Context) {
		url := context.Request.URL.Path
		context.JSON(200, "i am  "+url)
	})

	engine.GET("/lua/test", func(context *gin.Context) {
		url := context.Request.URL.Path
		header := context.Request.Header

		fmt.Println(header)

		context.JSON(200, "i am  "+url)
	})

	engine.GET("/bbc/test", func(context *gin.Context) {
		url := context.Request.URL.Path
		header := context.Request.Header

		fmt.Println(header)

		context.JSON(200, "i am  "+url)
	})

	engine.Run(":9010") //
}




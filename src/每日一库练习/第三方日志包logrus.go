package main

/**
logrus支持如下特性：
	完全兼容Go标准库日志模块。logrus拥有六种日志级别：debug、info、warn、error、fatal和panic，这是Go标准库日志模块的API的超集。如果你的项目使用标准库日志模块，完全可以用最低的代价迁移到logrus上。
	可扩展的Hook机制。允许使用者通过hook方式，将日志分发到任意地方，如本地文件系统、标准输出、logstash、elasticsearch或者mq等，或者通过hook定义日志内容和格式等。
	可选的日志输出格式。**logrus内置了两种日志格式，JSONFormatter和TextFormatter。**如果这两个格式不满足需求，可以自己动手实现接口Formatter，来定义自己的日志格式。
	Field机制。logrus鼓励通过Field机制进行精细化、结构化的日志记录，而不是通过冗长的消息来记录日志。
	logrus是一个可插拔的、结构化的日志框架。
	logrus不提供的功能：

没有提供行号和文件名的支持
	输出到本地文件系统没有提供日志分割功能
	没有提供输出到ELK等日志处理中心的功能
	这些功能都可以通过自定义hook来实现 。
*/

import (
	"bytes"
	logrus "github.com/sirupsen/logrus"
	"io"
	"os"
)

func initLogRus() {
	// todo 设置日志格式为json格式
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// todo 打印级别设置
	/**
	2.3 设置日志打印级别
	logrus 提供 6 档日志级别，分别是：

	PanicLevel
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	*/
	logrus.SetLevel(logrus.InfoLevel)
}

//2.5 自定义日志输出路径
func init() {
	//设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)
	//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	file, _ := os.OpenFile("1.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	logrus.SetOutput(file)
	//log.SetOutput(os.Stderr)

	//设置最低loglevel
	logrus.SetLevel(logrus.InfoLevel)
}

var writer1 = &bytes.Buffer{}

//todo 又能保存日志，又可以输出到console
func initAllLog() {

	writer2 := os.Stdout
	writer3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		logrus.Fatalf("create file log.txt failed: %v", err)
	}

	logrus.SetOutput(io.MultiWriter(writer1, writer2, writer3))
	logrus.SetLevel(logrus.InfoLevel)
	logrus.Info("info msg")
}

func main() {
	initLogRus() // 设置日志格式为json格式
	logrus.Info("我是一条日志")
	logrus.WithFields(logrus.Fields{"key": "value"}).Info("我要打印了")

	initAllLog() //
	logrus.WithFields(logrus.Fields{
		"age":  12,
		"name": "xiaoming",
		"sex":  1,
	}).Info("这是web1")

	logrus.WithFields(logrus.Fields{
		"age":  13,
		"name": "xiaohong",
		"sex":  0,
	}).Error("这是web2")

	logrus.WithFields(logrus.Fields{
		"age":  14,
		"name": "xiaofang",
		"sex":  1,
	}).Fatal("这是web3")

	//logrus.Info("writer: ", writer1.String())
}

//https://mojotv.cn/2018/12/27/golang-logrus-tutorial#NC4zIExvZ3J1cy1Ib29rIOaXpeW/l+WIhumalA==

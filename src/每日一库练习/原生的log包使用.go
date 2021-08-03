package main

import (
	"io"
	"log"
	"os"
)

/**
上面说到log打印的时候默认是自带时间的，那如果除了时间以外，我们还想要别的信息呢，当然log也是支持的。

SetFlags(flag int)方法提供了设置打印默认信息的能力，下面的字段是log中自带的支持的打印类型：
Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
Ltime                         // the time in the local time zone: 01:23:23
Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
Llongfile                     // full file name and line number: /a/b/c/d.go:23
Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
LstdFlags     = Ldate | Ltime // initial values for the standard logger
这是log包定义的一些抬头信息，有日期、时间、毫秒时间、绝对路径和行号、文件名和行号等，在上面都有注释说明，这里需要注意的是：如果设置了Lmicroseconds，那么Ltime就不生效了；设置了Lshortfile， Llongfile也不会生效，大家自己可以测试一下。

LUTC比较特殊，如果我们配置了时间标签，那么如果设置了LUTC的话，就会把输出的日期时间转为0时区的日期时间显示。

最后一个LstdFlags表示标准的日志抬头信息，也就是默认的，包含日期和具体时间。
*/
func initLog() {
	log.SetFlags(log.Ldate | log.Lshortfile)
}

func main() {
	initLog()
	log.Print("我就是一条日志")
	log.Printf("%s,", "谁说我是日志了，我是错误")

	Info.Println("我就是一条日志啊")
	Warning.Printf("我真的是一条日志哟%s\n", "别骗我")
	Error.Println("好了，我要报错了")

}

//======================================================

//通过上面的学习，你其实知道了，日志的实现是通过New()函数构造了Logger对象来处理的。
// 那我们只用构造不同的Logger对象来处理不同类型的日记即可。下面是一个简单的实现：
var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	infoFile, err := os.OpenFile("./info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	warnFile, err := os.OpenFile("./warn.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	errFile, err := os.OpenFile("./errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if infoFile == nil || warnFile == nil || err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}

	//Info = log.New(os.Stdout,"Info:",log.Ldate | log.Ltime | log.Lshortfile)
	//Warning = log.New(os.Stdout,"Warning:",log.Ldate | log.Ltime | log.Lshortfile)
	//Error = log.New(io.MultiWriter(os.Stderr,errFile),"Error:",log.Ldate | log.Ltime | log.Lshortfile)

	Info = log.New(io.MultiWriter(os.Stderr, infoFile), "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(io.MultiWriter(os.Stderr, warnFile), "Warning:", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, errFile), "Error:", log.Ldate|log.Ltime|log.Lshortfile)

}

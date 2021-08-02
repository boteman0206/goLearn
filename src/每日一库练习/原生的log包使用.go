package main

import "github.com/lunny/log"

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
func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
}

func main() {

	log.Print("我就是一条日志")
	log.Printf("%s,", "谁说我是日志了，我是错误")

}

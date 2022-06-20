package main

import (
	"bytes"
	"fmt"
)

func main() {

	//Read
	fmt.Println("=======Read============")
	bufRead := bytes.NewBufferString("hello")
	fmt.Println(bufRead.String())

	var sRead = make([]byte, 3)   // 定义读出的[]byte为3，表示一次可读出3个byte
	bufRead.Read(sRead)           // 读出
	fmt.Println(bufRead.String()) // 打印结果为lo,因为前三个被读出了
	fmt.Println(string(sRead))    // 打印结果为hel,读取的是hello的前三个字母

	bufRead.Read(sRead)           // 接着读，但是bufRead只剩下lo，所以只有lo被读出了
	fmt.Println(bufRead.String()) // 打印结果为空
	fmt.Println(string(sRead))    // 打印结果lol，前两位的lo表示的本次的读出，因为bufRead只有两位，后面的l还是上次的读出结果

	//ReadByte
	fmt.Println("=======ReadByte============")

	buf := bytes.NewBufferString("hello")
	fmt.Println(buf.String()) // buf.String()方法是吧buf里的内容转成string，>以便于打印
	b, _ := buf.ReadByte()    // 读取第一个byte，赋值给b
	fmt.Println(buf.String()) // 打印 ello，缓冲器头部第一个h被拿掉
	fmt.Println(string(b))    // 打印 h

	//ReadRune
	fmt.Println("=======ReadRune============")
	//
	buf1 := bytes.NewBufferString("好hello")
	fmt.Println(buf1.String()) // buf.String()方法是吧buf里的内容转成string，>以便于打印
	readRune, i, err := buf1.ReadRune()
	if err != nil {
		return
	} // 读取第一个rune，赋值给b
	fmt.Println(buf1.String())    // 打印 hello
	fmt.Println(string(readRune)) // 打印中文字： 好，缓冲器头部第一个“好”被拿掉
	fmt.Println(i)                // 打印3，“好”作为utf8储存占3个byte
	read2, n, err := buf1.ReadRune()
	if err != nil {
		return
	}
	// 再读取第一个rune，赋值给b
	fmt.Println(buf.String())  // 打印 ello
	fmt.Println(string(read2)) // 打印h，缓冲器头部第一个h被拿掉
	fmt.Println(n)             // 打印 1，“h”作为utf8储存占1个byte

	//ReadBytes
	fmt.Println("=======ReadBytes============")

	var d byte = 'e' //分隔符为e
	buf2 := bytes.NewBufferString("hello")
	fmt.Println(buf2.String()) // buf.String()方法是吧buf里的内容转成string，以便于打印
	b1, _ := buf2.ReadBytes(d) // 读到分隔符，并返回给b
	fmt.Println(buf2.String()) // 打印 llo，缓冲器被取走一些数据
	fmt.Println(string(b1))    // 打印 he，找到e了，将缓冲器从头开始，到e的内容都返回给b

	//ReadString
	fmt.Println("=====ReadString===============")
	var d1 byte = 'e' //分隔符为e
	buf3 := bytes.NewBufferString("hello")
	fmt.Println(buf3.String())   // buf.String()方法是吧buf里的内容转成string，以便于打印
	b3, _ := buf3.ReadString(d1) // 读到分隔符，并返回给b
	fmt.Println(buf3.String())   // 打印 llo，缓冲器被取走一些数据
	fmt.Println(b3)              // 打印 he，找到e了，将缓冲器从头开始，到e的内容都返回给b

	fmt.Println("=使用Next可依次读出固定长度的内容==============")
	buf4 := bytes.NewBufferString("hello")
	fmt.Println(buf4.String())
	b4 := buf4.Next(2)         // 重头开始，取2个
	fmt.Println(buf4.String()) // 变小了
	fmt.Println(string(b4))    // 打印he

	b5 := buf4.Next(2)
	fmt.Println(string(b5)) // 打印ll，已经取走了两个he
}

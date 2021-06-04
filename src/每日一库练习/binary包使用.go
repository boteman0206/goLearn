package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

/**
todo binary.Write的第三个参数应该是固定大小的数据类型（uint8, int64...)，或者由固定大小的数据类型组成的复合类型。
	string等底层也是byte数组无法转换

*/

func main() {
	var num uint64
	num = 0x1234
	fmt.Printf("num = %x\n", num)

	enc := make([]byte, 8)
	fmt.Println("num : ", num)
	// 转化为大端
	binary.BigEndian.PutUint64(enc, num)
	fmt.Printf("bigendian enc = %x\n", enc)

	// 转化为小端
	//binary.LittleEndian.PutUint64(enc, num)
	//fmt.Printf("littleendian enc = %x\n", enc)

	var data uint64
	reader := bytes.NewReader(enc)
	err2 := binary.Read(reader, binary.BigEndian, &data)
	if err2 != nil {
		fmt.Println("err2 :", err2.Error())
	}
	fmt.Println("dtat : ", data)

	var pi float64
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.LittleEndian, &pi)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Println("pi :", pi)

	//todo 先将类型写入buf 再read

	a := int64(1009)

	newBuffer := new(bytes.Buffer)
	//i := make([]byte, binary.Size(a))  // todo 不能使用这种方式，长度必须确定
	//newBuffer := bytes.NewBuffer(i)

	err2 = binary.Write(newBuffer, binary.LittleEndian, a)
	if err2 != nil {
		fmt.Println("write err : ", err2.Error())
	}
	fmt.Println("new buffer : ", newBuffer.Bytes(), binary.Size(a))

	var num1 int64
	err2 = binary.Read(newBuffer, binary.LittleEndian, &num1)
	if err2 != nil {
		fmt.Println(" read : ", err2.Error())
	}

	fmt.Println(" num1 : ", num1) // 1009

	fmt.Println("-------------test1----------")
	test1()

	fmt.Println("=========test2==========")
	test2()
}

func test1() {
	data := []byte("hello")
	w := bytes.NewBuffer(nil) // todo 直接new一个nil也是ok的
	binary.Write(w, binary.LittleEndian, data)
	size := binary.Size(data)
	orig := make([]byte, size)
	binary.Read(w, binary.LittleEndian, orig)
	fmt.Println(string(orig))
}

func test2() {

	// todo string类型无法转换，但是byte类型是可以的
	data := []string{"hello"}
	//data := "hello"
	w := bytes.NewBuffer(nil) // todo 直接new一个nil也是ok的
	err := binary.Write(w, binary.LittleEndian, data)
	fmt.Println("w == ", w, err.Error())
	var orig string
	binary.Read(w, binary.LittleEndian, &orig)
	fmt.Println(orig)
}

package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/**

https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247487043&idx=1&sn=3896db731e63c516254f9a346b31006e&chksm=fa80dfd4cdf756c2d52d825f948119b66a09da6b0ba821b62ef6ef5ccb99bd80d0551d281de5&token=195030300&lang=zh_CN&scene=21#wechat_redirect

unsafe包只有两个类型，三个函数，但是功能很强大。
type ArbitraryType int  ArbitraryType是int的一个别名，在 Go 中ArbitraryType有特殊的意义。代表一个任意Go表达式类型。Pointer是int指针类型的一个别名，在 Go 中可以把任意指针类型转换成unsafe.Pointer类型。
type Pointer *ArbitraryType

func Sizeof(x ArbitraryType) uintptr   Sizeof接受任意类型的值(表达式)，返回其占用的字节数,这和c语言里面不同，c语言里面sizeof函数的参数是类型，而这里是一个值，比如一个变量。
func Offsetof(x ArbitraryType) uintptr  Offsetof：返回结构体成员在内存中的位置距离结构体起始处的字节数，所传参数必须是结构体的成员（结构体指针指向的地址就是结构体起始处的地址，即第一个成员的内存地址）。
func Alignof(x ArbitraryType) uintptr   Alignof返回变量对齐字节数量，这个函数虽然接收的是任何类型的变量，但是有一个前提，就是变量要是一个struct类型，且还不能直接将这个struct类型的变量当作参数，只能将这个struct类型变量的值当作参数，具体细节咱们到以后聊内存对齐的文章里再说。


unsafe.Pointer  unsafe.Pointer称为通用指针，官方文档对该类型有四个重要描述：
1：任何类型的指针都可以被转化为Pointer
2： Pointer可以被转化为任何类型的指针
3：uintptr可以被转化为Pointer
4：Pointer可以被转化为uintptr
unsafe.Pointer是特别定义的一种指针类型（译注：类似C语言中的void类型的指针），在Go 语言中是用于各种指针相互转换的桥梁，它可以持有任意类型变量的地址。
什么叫"可以持有任意类型变量的地址"呢？意思就是使用 unsafe.Pointer 转换的变量，该变量一定要是指针类型，否则编译会报错。
a := 1
b := unsafe.Pointer(a) //报错
b := unsafe.Pointer(&a) // 正确

*/
func main() {

	v2 := int(13)
	p := (*uint)(unsafe.Pointer(&v2))
	*p = 90
	fmt.Println(p, *p)
	fmt.Println(v2)

	//string 和 []byte 零拷贝转换 可以使用
	s := "Hello World"
	b := string2bytes(s)
	fmt.Println(b)
	s = bytes2string(b)
	fmt.Println(s)

}

func string2bytes(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))

	bh := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}

	return *(*[]byte)(unsafe.Pointer(&bh))
}

func bytes2string(b []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	sh := reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}

	return *(*string)(unsafe.Pointer(&sh))
}

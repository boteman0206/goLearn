package main

import (
	"fmt"
	"reflect"
)

//todo 判断interface里的指针是否为空

//1、知道类型的情况下，自然是可以使用类型断言后判空。如ai, ok := i.(*int)，之后判断ai == nil。

//2、不知道是何种类型的指针，就只好借助反射了vi := reflect.ValueOf(i)，后使用vi.IsNil()来判断。但如果i里放到不是一个指针，调用IsNil会出异常，则可能要写一个这样的函数来判空
func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

//其他注意事项

//todo nil 一般是判断结构体的指针是否为空
type test struct {
	Name     string
	Password string
}

func testNil() {
	var data *test
	if data == nil {
		fmt.Println("struct data is empty !")
	}
}

//todo len(t) 一般用于求数组、切片的长度的时候 nil slice 和 empty slice是不同的东西，

func main() {
	var test = make([]int, 0)
	fmt.Println("空数组： ", test, test == nil) // false
	test = nil
	fmt.Println("设置nil之后： ", test, test == nil) // true
	if len(test) == 0 {
		fmt.Println("[]string test is empty !")
	}
}

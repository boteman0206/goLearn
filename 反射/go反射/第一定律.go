package main

import (
	"fmt"
	"reflect"
)

type Coder struct {
	Name string
}

func (c Coder) String() string {
	return c.Name
}

func main() {

	coder := Coder{
		Name: "code name1",
	}

	typ := reflect.TypeOf(coder)
	val := reflect.ValueOf(&coder) // todo 使用val.Elem()，必须要是地址指针， 否则报错

	fmt.Println("type: ", typ, " valuer: ", val, val.Elem())
	fmt.Println("kind : ", typ.Kind(), typ.Field(0))

	of := reflect.ValueOf(coder) // todo 注意 val.Field(0) 会报错， 因为上面的valueOf是指针
	fmt.Println("of filed: ", of.Field(0))

	stringElem := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	fmt.Println("stringElem: ", stringElem)

	// 判断是否实现某个方法
	fmt.Println("IMPLEMENT: ", typ.Implements(stringElem))

}

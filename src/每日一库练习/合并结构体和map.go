package main

import (
	"fmt"
	"github.com/imdario/mergo"
	"log"
)

type Foo struct {
	A string
	B int64
}

//todo  https://segmentfault.com/a/1190000021997004 地址

type redisConfig struct {
	Address string
	address string
	Port    int
	DBs     []int
}

var defaultConfig = redisConfig{
	Address: "127.0.0.1",
	Port:    6381,
	DBs:     []int{1},
}

var defaultConfig1 = redisConfig{
	Address: "127.0.0.1",
	Port:    6381,
}

func main() {

	src := Foo{
		A: "one",
		B: 2,
	}
	dest := Foo{
		A: "two",
		B: 1,
	}

	//todo 参数 1 是目标对象，参数 2 是源对象，这两个函数的功能就是将源对象中的字段复制到目标对象的对应字段上。
	//1: 将src中的值合并到dest中
	mergo.Merge(&dest, src)
	fmt.Println(dest)
	// Will print
	// {two 1}

	//2: 将src中的值强制覆盖dest中所有的值
	mergo.Merge(&dest, src, mergo.WithOverride)
	fmt.Println(dest)
	// {one 2}

	// 添加切片的值
	var config redisConfig
	config.DBs = []int{2, 3}

	//todo将两个切片的值进行合并
	if err := mergo.Merge(&config, defaultConfig, mergo.WithAppendSlice); err != nil {
		log.Fatal(err)
	}

	fmt.Println("redis address: ", config.Address)
	fmt.Println("redis port: ", config.Port)
	fmt.Println("redis dbs: ", config.DBs)

	//todo 空值覆盖
	/**
		默认情况下，如果源对象中的字段为空值（数组、切片长度为 0 ，指针为nil，数字为 0，字符串为""等），即使我们使用了WithOverride选项也是不会覆盖的。下面两个选项就是强制这种情况下也覆盖：
	WithOverrideEmptySlice：源对象的空切片覆盖目标对象的对应字段；
	WithOverwriteWithEmptyValue：源对象中的空值覆盖目标对象的对应字段，其实这个对切片也有效。
	文档中这两个选项的介绍比较混乱，我通过看源码和自己试验下来发现：
	这两个选项都必须和WithOverride一起使用；
	WithOverwriteWithEmptyValue这个选项也可以处理切片类型的值。
	*/

	var config1 redisConfig
	config1.DBs = []int{2, 3}

	if err := mergo.Merge(&config1, defaultConfig1, mergo.WithOverride, mergo.WithOverrideEmptySlice); err != nil {
		log.Fatal(err)
	}

	fmt.Println("redis address: ", config1.Address)
	fmt.Println("redis port: ", config1.Port)
	fmt.Println("redis dbs: ", config1.DBs)
	//最终会输出空的DBs。

	m1 := make(map[string]interface{})
	m1["dbs"] = []uint32{2, 3}

	m2 := make(map[string]interface{})
	m2["DRbs"] = []int{1}

	if err := mergo.Map(&m1, &m2, mergo.WithOverride); err != nil {
		log.Fatal(err)
	}
	fmt.Println(m1)

	//todo  结构体到map的赋值
	mergo.Map(&m1, defaultConfig)

	fmt.Println("结构体到map： ", m1)

	//map到结构体的赋值
	var s1 redisConfig
	m := make(map[string]interface{})
	m["address"] = "121.12.31.23"
	m["port"] = 1223
	mergo.Map(&s1, m)
	fmt.Println("结构体到map：", s1.Port, s1.Address, s1.DBs, s1.address)

	//todo 如果map的类型不是interface
	/**
	todo panic: interface conversion: interface {} is map[string][]uint8, not map[string]interface {}
	//var s11 redisConfig
	//m3 := make(map[string][]byte, 1)
	//m3["address"] = []byte("121.12.31.23")
	//m3["port"] = []byte(string(1223))
	//mergo.Map(&s11, m3)
	//fmt.Println("复制非空接口的数据：", s1.Port, s1.Address, s1.DBs, s1.address)
	*/

}

/**
todo
	1: mergo不会赋值非导出字段；
	2: map中对应的键名首字母会转为小写；
	3: mergo可嵌套赋值，我们演示的只有一层结构。
*/

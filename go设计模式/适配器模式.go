package main

import "fmt"

// Target 是适配的目标接口
type Target interface {
	Request() string
}

// Adaptee 是适配的目标接口
type Adaptee interface {
	SpecificRequest() string
}

// NewAdaptee 是被适配接口的工厂函数
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

// 是适配器的目标类
type adapteeImpl struct {
}

func (a *adapteeImpl) SpecificRequest() string {
	return "apatee method"
}

// NewAdapter 是adapter的工厂函数 todo 将Adaptee转换成了Target接口
func NewAdapter(adaptee Adaptee) Target {
	return &adapter{adaptee}
}

//todo Adapter 是转换Adaptee为Target接口的适配器
type adapter struct {
	Adaptee
}

func (a adapter) Request() string {
	return a.SpecificRequest()
}

/**
	适配器模式用于转换一种接口适配另一种接口。

实际使用中Adaptee一般为接口，并且使用工厂函数生成实例。

在Adapter中匿名组合Adaptee接口，所以Adapter类也拥有SpecificRequest实例方法，
又因为Go语言中非入侵式接口特征， 其实Adapter也适配Adaptee接口。
*/
func main() {

	adaptee := NewAdaptee()
	newAdapter := NewAdapter(adaptee)
	request := newAdapter.Request()
	fmt.Println("request: ", request)

	//i, ok := newAdapter.(*adapter)
	//fmt.Println("test : ", i.SpecificRequest(), ok)
}

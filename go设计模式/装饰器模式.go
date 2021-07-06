package main

import "fmt"

/**
装饰模式
装饰模式使用对象组合的方式动态改变或增加对象行为。
Go语言借助于匿名组合和非入侵式接口可以很方便实现装饰模式。
使用匿名组合，在装饰器中不必显式定义转调原对象方法。
*/

type Component interface {
	Calc() int
}

type ConcreteComponent struct {
}

func (*ConcreteComponent) Calc() int {
	return 0
}

type MulDecorator struct {
	Component
	num int
}

func WarpMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		c,
		num,
	}
}

func (d *MulDecorator) Calc() int {

	return d.Component.Calc() * d.num
}

type AddDecorator struct {
	Component
	num int
}

func WarpAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		c,
		num,
	}
}

func (d *AddDecorator) Calc() int {
	return d.Component.Calc() + d.num
}

func main() {

	component := ConcreteComponent{}
	c := WarpAddDecorator(&component, 11)

	c = WarpMulDecorator(c, 90)

	c = WarpAddDecorator(c, 10)

	i := c.Calc()

	fmt.Println("WarpMulDecorator: ", i)

}

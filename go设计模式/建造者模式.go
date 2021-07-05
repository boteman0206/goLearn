package main

import "fmt"

// builder是生成器的接口

type Builder interface {
	Part1()
	Part2()
	Part3()
}

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

// 构造方法
func (d *Director) Construct() {

	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}

//==================Builder1==========================
type Builder1 struct {
	result string
}

func (b *Builder1) Part1() {
	b.result += "1"
}

func (b *Builder1) Part2() {
	b.result += "2"
}

func (b *Builder1) Part3() {
	b.result += "3"
}

func (b *Builder1) GetResult() string {
	return b.result
}

//++++++++++++++++++++Builder2+++++++++++++++++++++++++++++++++
type Builder2 struct {
	result int
}

func (b *Builder2) Part1() {
	b.result += 1
}

func (b *Builder2) Part2() {
	b.result += 2
}

func (b *Builder2) Part3() {
	b.result += 3
}
func (b *Builder2) GetResult() int {
	return b.result
}

func main() {

	//builder1建造
	builder1 := &Builder1{}
	director := NewDirector(builder1)
	director.Construct()
	fmt.Println("builder1 : ", builder1.GetResult())

	// builder2建造
	builder2 := &Builder2{}
	newDirector := NewDirector(builder2)
	newDirector.Construct()
	fmt.Println("builder2: ", builder2.GetResult())

}

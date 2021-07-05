package main

import "fmt"

/**
	原型模式使对象能复制自身，并且暴露到接口中，使客户端面向接口编程时，不知道接口实际对象的情况下生成新的对象。
原型模式配合原型管理器使用，使得客户端在不知道具体类的情况下，通过接口管理器得到新的实例，并且包含部分预设定配置。
*/

// Cloneable是原型对象需要实现的接口
type Cloneable interface {
	Clone() Cloneable
}

type Prototypemanager struct {
	prototypes map[string]Cloneable
}

func NewPrototypeManager() *Prototypemanager {
	return &Prototypemanager{
		prototypes: make(map[string]Cloneable),
	}
}

func (p *Prototypemanager) Get(name string) Cloneable {
	return p.prototypes[name].Clone()
}

func (p *Prototypemanager) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}

type Type1 struct {
	name string
}

func (t *Type1) Clone() Cloneable {
	tc := *t
	return &tc
}

type Type2 struct {
	name string
}

func (t *Type2) Clone() Cloneable {
	tc := *t
	return &tc
}

func main() {
	manager := NewPrototypeManager()
	t1 := &Type1{
		name: "type1",
	}
	manager.Set("t1", t1)
	c := manager.Get("t1").Clone()
	//t1.name = "ppp"  // 修改t1的原始值， 后面的c并没有发生变化
	t11 := c.(*Type1)
	fmt.Println("type1: ", t1, "克隆之后的对象：", t11, t11.name)
	fmt.Println("is true : ", t1 == c) // false

	t1.name = "ppp"

	fmt.Println()
}

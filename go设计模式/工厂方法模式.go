package main

func main() {

}

// operator是被封装的实际接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// 是工厂接口
type OperatorFactory interface {
	Create() Operator
}

// 是operator的接口实现的积累，封装
type OperatorBase struct {
	a, b int
}

// 设置A
func (o *OperatorBase) SetA(a int) {
	o.a = a
}

// 设置B
func (o *OperatorBase) SetB(b int) {
	o.b = b
}

// 是plusOperator的工厂类
type PlusOperatorFactory struct {
}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		&OperatorBase{},
	}
}

type PlusOperator struct {
	*OperatorBase
}

func (p PlusOperator) Result() int {
	return p.a + p.b
}

//MinusOperatorFactory 是 MinusOperator 的工厂类
type MinusOperatorFactory struct{}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}

//MinusOperator Operator 的实际减法实现
type MinusOperator struct {
	*OperatorBase
}

//Result 获取结果
func (o MinusOperator) Result() int {
	return o.a - o.b
}

// operator是被封装的实际接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// 是工厂接口
type OperatorFactory interface {
	Create() Operator
}

// 是operator的接口实现的积累，封装
type OperatorBase struct {
	a, b int
}

// 设置A
func (o *OperatorBase) SetA(a int) {
	o.a = a
}

// 设置B
func (o *OperatorBase) SetB(b int) {
	o.b = b
}

// 是plusOperator的工厂类
type PlusOperatorFactory struct {
}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		&OperatorBase{},
	}
}

type PlusOperator struct {
	*OperatorBase
}

func (p PlusOperator) Result() int {
	return p.a + p.b
}

//MinusOperatorFactory 是 MinusOperator 的工厂类
type MinusOperatorFactory struct{}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}

//MinusOperator Operator 的实际减法实现
type MinusOperator struct {
	*OperatorBase
}

//Result 获取结果
func (o MinusOperator) Result() int {
	return o.a - o.b
}

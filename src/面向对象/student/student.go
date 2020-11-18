package student

type person struct {
	name string
	age  int
	addr string
}

// 类似于构造方案
func NewPerson(name string, age int, addr string) *person {
	return &person{
		name: name,
		age:  age,
		addr: addr}
}

//获取name
func (p *person) GetName() string {
	return p.name
}

func (p *person) SetName(name string) {
	p.name = name
}

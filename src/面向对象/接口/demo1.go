package 接口

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Computer struct {
}

type Phone struct {
}

func (phone Phone) Start() {
	fmt.Println("phone call!!")
}

func (s Computer) Run(phone Phone) {
	phone.Start()
}

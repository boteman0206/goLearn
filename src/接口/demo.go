package main

import "fmt"

type Phone interface {
	call()
	tell()
}

type NokiaPhone struct {
}

type Iphone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("i am NokiaPhone, call ...")
}

func (nokiaPhone NokiaPhone) tell() {
	fmt.Println("i am NokiaPhone,  tell ...")
}

func (iPhone Iphone) call() {
	fmt.Println("i am Iphone, call ...")
}

func (iPhone Iphone) tell() {
	fmt.Println("i am Iphone, tell ....")
}

func main() {
	var phone Phone

	phone = new(Iphone)
	phone.call()
	phone.tell()

	phone = new(NokiaPhone)
	phone.call()
}

package main

import (
	"fmt"
	"github.com/pkg/errors"
	"runtime/debug"
)

// controller层的使用
func Controller() error {
	err := Service()
	//if err != nil {
	//	log.Printf("stack trace: \n%+v\n", err)
	//}

	fmt.Println("Controller:  ", err.Error())
	return nil
}

// service层的代码
func Service() error {
	err := Dao()
	return fmt.Errorf("Service failed, error=%v", err)
}

// dao层的使用
func Dao() error {

	err := getErr()
	if err != nil {
		return errors.Wrapf(err, "Dao fail")
	}

	return err

}

func getErr() error {
	return errors.New("division by zero")
}

func main() {

	Controller()
	//fmt.Println("err is ", err.Error())
	fmt.Println("===================")
	byte := debug.Stack()
	fmt.Println("byte:", string(byte))
}

package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

type User struct {
	Name string
	Age  int
	Pets []string
}

type Employee struct {
	Name string
	Age  int
	Role string
	Pets []string
}

func main() {
	user := User{Name: "dj", Age: 18}
	employee := Employee{}
	employee1 := Employee{Pets: []string{"lll"}}

	copier.Copy(&employee, &user)
	copier.CopyWithOption(&employee1, &user, copier.Option{false, false, nil})
	fmt.Printf("%#v\n", employee)
	fmt.Printf("另外的copy %#v\n", employee1)
	user.Name = "pop"
	fmt.Println("user : ", user)
	fmt.Println(employee)

	fmt.Println()
	// todo 切片赋值
	users := []User{
		{Name: "dj", Age: 18},
		{Name: "dj2", Age: 18},
	}
	employees := []Employee{}

	copier.Copy(&employees, &users)
	fmt.Printf("%#v\n", employees)

	fmt.Println()
	//todo 将结构体赋值到切片中
	user1 := User{Name: "dj", Age: 18}
	employees1 := []Employee{}

	copier.Copy(&employees1, &user1)
	fmt.Printf("%#v\n", employees1)
}

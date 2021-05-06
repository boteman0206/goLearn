package main

import (
	"fmt"

	"github.com/jinzhu/copier"
)

type User struct {
	Name string
	Age  int
}

type Employee struct {
	Name string
	Age  int
	Role string
}

func main() {
	user := User{Name: "dj", Age: 18}
	employee := Employee{}

	copier.Copy(&employee, &user)
	fmt.Printf("%#v\n", employee)
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

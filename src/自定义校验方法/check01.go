package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Users struct {
	Phone  string `form:"phone" json:"phone" validate:"required"`
	Passwd string `form:"passwd" json:"passwd" validate:"required,max=20,min=6"`
	Code   string `form:"code" json:"code" validate:"required,eqfield=Passwd"` // eqfield标识code字段和passwd的字段必须一样
}

func main() {

	users := &Users{
		Phone:  "1326654487",
		Passwd: "1232344",
		Code:   "1232344",
	}
	validate := validator.New()

	err := validate.Struct(users)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err) //Key: 'Users.Passwd' Error:Field validation for 'Passwd' failed on the 'min' tag
			return
		}
	}
	return
}

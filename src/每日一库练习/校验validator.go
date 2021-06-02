package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	vt "github.com/tricobbler/echo-tool/validate"
)

func main() {

	a := ""
	validate := validator.New()

	err := validate.Var(a, "required")
	if err != nil {
		translate := vt.Translate(err.(validator.ValidationErrors))
		fmt.Println("不能为空", translate)
	}

	num1 := 2
	err = validate.Var(num1, "gte=3")
	if err != nil {
		fmt.Println("gte 必须大于等于 3")
	}

}

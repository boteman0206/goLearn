package main

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

type Users struct {
	Name   string `form:"name" json:"name" validate:"required" label:"姓名"`
	Age    uint8  `form:"age" json:"age" validate:"required,gt=18" label:"年齡"` //todo  添加label标签作为中文的字段返回
	Passwd string `form:"passwd" json:"passwd" validate:"required,max=20,min=6" label:"密码"`
	Code   string `form:"code" json:"code" validate:"required,len=6" label:"编码"`
}

func main() {
	users := &Users{
		Name:   "",
		Age:    12,
		Passwd: "123",
		Code:   "123456",
	}
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh") // 添加中文错误返回
	validate := validator.New()
	//验证器注册翻译器
	// todo 注册一个函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})

	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println(err)
	}
	err = validate.Struct(users)
	if err != nil {
		//方式一：
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Translate(trans)) //Age必须大于18
		}

	}

	//方式二
	/**
	errors := err.(validator.ValidationErrors)
	translate := errors[0].Translate(trans)
	fmt.Println("translate: ", translate)

	*/

	return
}

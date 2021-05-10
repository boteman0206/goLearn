package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	myValidate "github.com/tricobbler/echo-tool/validate"
	"io/ioutil"
	"net/http"
)

type (
	User struct {
		Name  string `json:"name" validate:"required,gte=4,lowercase,excludes=a" label:"用户名称"`
		ID    int    `json:"id" validate:"required,gte=2" label:"用户的id"`
		Email string `json:"email" validate:"required" label:"用户的邮箱"` // 校验必填。并且邮箱必须未邮箱的格式
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()
	//e.Validator = &CustomValidator{validator: validator.New()}

	// todo 公司自己定义的validate测试
	e.Validator = myValidate.NewCustomValidator()

	//e.POST("/users", func(c echo.Context) (err error) {
	//	u := new(User)
	//	if err = c.Bind(u); err != nil {
	//		return
	//	}
	//	if err = c.Validate(u); err != nil {
	//		fmt.Println("this validate error is : ", err)
	//		return c.JSON(http.StatusBadRequest, err.Error())
	//	}
	//	return c.JSON(http.StatusOK, u)
	//})

	// todo 自己的校验器
	e.POST("/users", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return
		}
		if err := c.Validate(u); err != nil {
			fmt.Println("this is my validate...")
			//err := myValidate.Translate(err.(validator.ValidationErrors))
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, u)
	})
	// todo 输出为接口的json文件  routes.json
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		fmt.Println("error : ", err)
	}
	ioutil.WriteFile("routes.json", data, 0644)

	e.Logger.Fatal(e.Start(":9001"))

}

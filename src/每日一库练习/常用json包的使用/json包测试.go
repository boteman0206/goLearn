package 常用json包的使用

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
)

type Per1 struct {
	Nmae  string
	Age   int
	Num   int
	Phone string
}

func main() {
	sname := `{"name":{"first":"li",last:"dj"},"age":18}`
	valid := json.Valid([]byte(sname))
	if valid {
		fmt.Println("校验成功")
	} else {
		fmt.Println("校验失败")
	}
	//s1 := `{"Nmae":"jack","Age":12,"Num":90,"Phone":"12345676"}`

	perData := Per1{"jack", 12, 90, "12345676"}

	marshal, err := json.Marshal(perData)
	fmt.Println(string(marshal), "---", err)

	json := `{"name":{"first":"li","last":"dj"},"age":18}`

	lastName := gjson.Get(json, "name.last")
	fmt.Println("last name:", lastName.String())

	age := gjson.Get(json, "age")
	fmt.Println("age:", age.Int())

	fmt.Println("=========================")
	const json1 = `
	{
  "name":{"first":"Tom", "last": "Anderson"},
  "age": 37,
  "children": ["Sara", "Alex", "Jack"],
  "fav.movie": "Dear Hunter",
  "friends": [
    {"first": "Dale", "last":"Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ]
	}
`

	fmt.Println("last name:", gjson.Get(json1, "name.last"))
	fmt.Println("age:", gjson.Get(json1, "age"))
	fmt.Println("children:", gjson.Get(json1, "children"))
	fmt.Println("children count:", gjson.Get(json1, "children.#"))
	fmt.Println("second child:", gjson.Get(json1, "children.1"))
	fmt.Println("third child*:", gjson.Get(json1, "child*.2"))
	fmt.Println("first c?ild:", gjson.Get(json1, "c?ildren.0"))
	fmt.Println("fav.moive", gjson.Get(json1, `fav.\moive`))
	fmt.Println("first name of friends:", gjson.Get(json1, "friends.#.first"))
	fmt.Println("last name of second friend:", gjson.Get(json1, "friends.1.last"))
	rows := gjson.Get(json1, "friends.#.first")

	rows.ForEach(func(key, value gjson.Result) bool {
		fmt.Println(key, "==", value)
		return true
	})

	// todo 校验json
	if !gjson.Valid(json) {
		fmt.Println("error json!")
	} else {
		fmt.Println("ok json!")
	}

	// todo 一次获取多个值
	const name = `
	{
		  "name":"dj",
		  "age":18,
		  "pets": ["cat", "dog"],
		  "contact": {
			"phone": "123456789",
			"email": "dj@example.com"
		  }
		}`

	many := gjson.GetMany(name, "name", "age", "pets.#", "contact.phone")
	for key, result := range many {
		fmt.Println(key, " result ", result)
	}

}

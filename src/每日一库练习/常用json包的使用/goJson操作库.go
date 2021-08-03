package 常用json包的使用

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq/v2"
)

func main() {
	const json = `{"name":{"first":"Tom","last":"Hanks"},"age":61}`
	name := gojsonq.New().FromString(json).Find("name.first")
	fmt.Println(name.(string)) // Tom

	const json1 = `{"city":"dhaka","type":"weekly","temperatures":[30,39.9,35.4,33.5,31.6,33.2,30.7]}`
	avg := gojsonq.New().FromString(json1).From("temperatures").Avg()
	fmt.Printf("Average temperature: %.2f", avg) // 33.471428571428575

	fmt.Println("=================")

}

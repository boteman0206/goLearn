package if和switch

import (
	"fmt"
	"go/types"
)

func main() {

	var age int
	fmt.Println("请输入你的年龄 ： ")
	fmt.Scanln(&age)

	/**
	if分支
	*/
	if age > 58 { // 不需要加小括号
		fmt.Println("老油条")
	} else if age > 18 {
		fmt.Println("有点东西！")
	} else {
		fmt.Println("太年轻！")
	}

	/**
	switch分支； 不需要break 默认就有
	1： fallthrough可以穿透下一层， 而不需要判断
	*/
	fmt.Println("请输入月份 ： ")
	var month string
	fmt.Scanln(&month)

	switch month {
	case "1", "2", "3":
		fmt.Println("春")
		fallthrough // 穿透一层
	case "4", "5", "6":
		fmt.Println("夏")
	case "7", "8", "9":
		fmt.Println("秋")
	case "10", "11", "12":
		fmt.Println("冬天")
	default:
		fmt.Println("请输入正确的月份！")
	}

	// 用于判断某个interface中，实际指向的变量
	var x interface{}
	var y int = 10
	x = y
	switch i := x.(type) {
	case types.Nil:
		fmt.Println("x的类型 %T", i)
	case int:
		fmt.Println("x是int类型")
	case float32:
		fmt.Println("x是float32")
	default:
		fmt.Println("未知类型")
	}

	fmt.Println("x : ", x)
}

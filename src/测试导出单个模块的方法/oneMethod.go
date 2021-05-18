package oneMethod

/**
todo 没有测试出来
*/
import "fmt"

var ST interface{}

func init() {
	ST = Myone2
}

func MyOne1() {
	fmt.Println("this is one 1")
}

func Myone2() {
	fmt.Println("this is one 2")
}

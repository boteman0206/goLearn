package main

/**
基于 Go 1.18+ 泛型（map、filter、contains、find...）的 Lodash 风格的 Go 库

Lodash: 是一个javascript风格的库
*/

import (
	"errors"
	"fmt"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"reflect"
	"strconv"
	"time"
)

func main() {
	names := lo.Uniq[string]([]string{"Samuel", "Marc", "Samuel"})
	// []string{"Samuel", "Marc"}
	fmt.Println("names: ", names)

	even := lo.Filter[int]([]int{1, 2, 3, 4}, func(x int, y int) bool {
		fmt.Println(x, y) // 找出奇数的数据 y代表的是index索引
		return x%2 == 0
	})
	fmt.Println(even)

	imap := lo.Map[int64, string]([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
		//return cast.ToString(x)
	})
	fmt.Println(imap, reflect.TypeOf(imap))

	find, b := lo.Find([]int{1, 2, 3, 4, 1}, func(x int) bool {
		if x == 10 {
			return true
		}
		return false
	})
	fmt.Println("Find: ", find, b)

	of, i, b2 := lo.FindLastIndexOf([]int{1, 2, 3, 4, 1}, func(x int) bool {
		if x == 1 {
			return true
		}
		return false
	})
	fmt.Println("FindLastIndexOf: ", of, i, b2)

	// 查找重复的数据
	duplicates := lo.FindDuplicates([]int{1, 2, 3, 4, 1, 2, 3})
	fmt.Println("FindDuplicates: ", duplicates)

	min1 := lo.Min[int]([]int{1, 2, 3})
	// 1
	min2 := lo.Min[int]([]int{})
	// 0
	fmt.Println("min: ", min1, min2)

	minBY := lo.MinBy[string]([]string{"s1", "string2", "s3"}, func(item string, min string) bool {
		return len(item) < len(min)
	})
	// "s1"
	fmt.Println("minBy:", minBY)

	s := lo.Synchronize()
	var num = 0
	for i1 := 0; i1 < 10; i1++ {
		go s.Do(func() {
			num++
		})
	}
	time.Sleep(2 * time.Second)
	fmt.Println("Synchronize: ", num)

	//方法包装
	//Calls the function and return false in case of error and on panic. // 遇到panic或者err的时候返回false
	ok := lo.Try(func() error {
		//panic("error")
		return errors.New("err")
		//return nil
	})

	fmt.Println("try : ", ok)

	//golang的TryCatch 和try的方法相同，但在发生错误时调用catch函数。
	caught := false
	lo.TryCatch(func() error {
		panic("error")
		return nil
	}, func() {
		caught = true
	})
	// false
	// caught == true
	fmt.Println("TryCatch: ", caught)

	// 和try一样，但是会返回错误信息
	err, ok := lo.TryWithErrorValue(func() error {
		panic("error")
		return nil
	})
	// "error", false
	fmt.Println("TryWithErrorValue: ", err, ok)

	in := []string{"", "foo", "", "bar", ""} // 返回非空的集合
	slice := lo.Compact[string](in)
	// []string{"foo", "bar"}
	fmt.Println("Compact: ", slice)

	// go协程处理的，打印顺序不确定
	lop.ForEach[string]([]string{"hello", "world"}, func(x string, _ int) {
		println(x)
	})
	// prints "hello\nworld\n" or "world\nhello\n"

	// 循环处理打印顺序确定
	lo.ForEach[string]([]string{"hello", "world"}, func(x string, _ int) {
		println(x)
	})

	// 交集
	result1 := lo.Intersect[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
	// []int{0, 2}
	fmt.Println("Intersect: ", result1)

	// 差集
	left, right := lo.Difference[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 2, 6})
	// []int{1, 3, 4, 5}, []int{6}
	fmt.Println("Difference: ", left, right)

	//Union 并集
	union := lo.Union[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 2, 10})
	// []int{0, 1, 2, 3, 4, 5, 10}
	fmt.Println("Union: ", union)

	subset := lo.Without[int]([]int{0, 2, 10}, 2)
	fmt.Println("Without: ", subset)
}

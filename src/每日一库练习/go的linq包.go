package main

import (
	"fmt"
	"github.com/ahmetb/go-linq/v3"
	"github.com/spf13/cast"
	"strings"
)

/**
类似java中的stream流操作包，可以链式的做一些筛选和过滤操作
*/

func main() {

	type Car struct {
		year int
		Name string
	}

	var owners []string

	cars := []Car{{1, "1"}, {2016, "jack"}, {2017, "bob"}, {2015, "pop"}}

	// 1 todo、 返回特定的值
	linq.From(cars).Where(func(c interface{}) bool {
		return c.(Car).year > 2015
	}).Select(func(c interface{}) interface{} {
		var string = cast.ToString(c.(Car).Name) + " : " + cast.ToString(c.(Car).year)
		return string
	}).ToSlice(&owners)

	fmt.Println(owners)

	var ownerCar []Car

	carsList := []Car{{1, "1"}, {2016, "jack"}, {2017, "bob"}, {2015, "pop"}}
	//2 todo 返回全部的值
	linq.From(carsList).Where(func(c interface{}) bool {
		return c.(Car).year > 2015
	}).Where(func(i interface{}) bool {
		return len(i.(Car).Name) > 3
	}).Select(func(c interface{}) interface{} {
		return c.(Car)
	}).ToSlice(&ownerCar)

	fmt.Println("ownerCar: ", ownerCar)

	// 3 todo 排序倒序
	ages := []int{21, 46, 46, 55, 17, 21, 55, 55}

	var distinctAges []int // orderbydescentt是orderby降序的类型版本
	linq.From(ages).OrderByDescendingT(
		func(item interface{}) interface{} { return item },
	).Distinct().ToSlice(&distinctAges)

	fmt.Println(distinctAges)

	// 4 迭代器 todo  Generate a slice of integers from 1 to 10
	var squares []int
	linq.Range(1, 10).
		SelectT(
			func(x int) int { return x * x },
		).
		ToSlice(&squares)

	for _, num := range squares {
		fmt.Println("Generate a slice of integers from 1 to 10 : ", num)
	}

	// todo 重复数据
	var slice []string
	linq.Repeat("I like programming.", 5).
		ToSlice(&slice)

	for _, str := range slice {
		fmt.Println(str)
	}

	// todo Aggregate使用
	fruits := []string{"apple", "mango", "orange", "passionfruit", "grape"}

	// Determine which string in the slice is the longest.
	longestName := linq.From(fruits).Aggregate(func(r interface{}, i interface{}) interface{} {
		if len(r.(string)) > len(i.(string)) {
			return r
		}
		return i
	},
	)

	fmt.Println(longestName)

	fruits1 := []string{"apple", "mango", "orange", "passionfruit", "grape"}

	// Determine which string in the slice is the longest.
	longestName1 := linq.From(fruits1).Aggregate(func(r interface{}, i interface{}) interface{} {
		return r.(string) + " : " + i.(string)
	},
	)

	fmt.Println("Determine which string in the slice is the longest. == ", longestName1)

	// all是否全量满足
	pets := []string{"asasB", "223B", "445B"}
	allStartWithB := linq.From(pets).All(
		func(pet interface{}) bool { return strings.HasSuffix(pet.(string), "B") })

	fmt.Println("All pet names start with 'B'? %t", allStartWithB)

	numbers := []int{1, 2, 3, 4}
	hasElements := linq.From(numbers).Any()

	fmt.Println("Are there any element in the list? %t", hasElements)

	// append
	input := []int{1, 2, 3, 4}

	q := linq.From(input).Append(2)

	last := q.Last()

	fmt.Println("append  ==== ", last, q.Count(), q.First(), q.Results())
	iterate := q.Iterate()
	for item, ok := iterate(); ok; item, ok = iterate() {
		fmt.Println("----遍历迭代----", item)
	}

	// 求平均值
	grades := []int{78, 92, 100, 37, 81}
	average := linq.From(grades).Average()

	fmt.Println(average)

	//concat 合并元素
	q1 := linq.From([]int{1, 2, 3}).
		Concat(linq.From([]int{4, 5, 6}))

	fmt.Println(q1.Results())

	// 判断contains
	slice1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	has5 := linq.From(slice1).Contains(5)
	fmt.Println("Does the slice contains 5? %t", has5)

	// 去重
	ages1 := []string{"abc", "qweq", "opopo", "bob", "abc", "abc", "abc", "bob"}
	var distinctAges1 []string

	linq.From(ages1).Distinct().ToSlice(&distinctAges1)

	fmt.Println("数组去重： ", distinctAges1)

	type Product struct {
		Name string
		Code int
	}

	products := []Product{
		{Name: "apple", Code: 9},
		{Name: "orange", Code: 4},
		{Name: "apple", Code: 9},
		{Name: "lemon", Code: 12},
	}

	//Exclude duplicates.
	var noduplicates []Product
	linq.From(products).DistinctByT(func(item Product) int { return item.Code }).ToSlice(&noduplicates)
	for _, product := range noduplicates {
		fmt.Printf("%s %d\n", product.Name, product.Code)
	}

	str := "你n好 你好啊"
	//var strSlice []string
	r := linq.FromString(str).Distinct().Count()
	contains := linq.FromString(str).Contains("n") // todo 注意这种contains无效，不能这么使用
	fmt.Println("strSlice : ", r, contains)
	b := strings.Contains(str, "n")
	fmt.Println(b)

}

package main

import (
	"fmt"
	"github.com/spf13/cast"
)

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //[[1 2] [3 4] [5 6] [7 8] [9 10]]
	fmt.Println(arr)                            //[1 2 3 4 5 6 7 8 9 10]
	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}	//[[1 2] [3 4] [5 6] [7 8] [9]]

	var si []int

	for i := 0; i < 109; i++ {
		si = append(si, i)
	}

	//res := arrayInGroupsOf(arr, 2)
	res1 := arrayInGroupsOf(si, 10)
	fmt.Println(res1)

	for i := range res1 {
		ints := res1[i]
		fmt.Println(ints)
	}

	var strData []string
	for i := 0; i < 198; i++ {
		strData = append(strData, cast.ToString(i))
	}
	of := ArrayStrGroupsOf(strData, 10)
	fmt.Println("str of : ", of)
}

func arrayInGroupsOf(arr []int, num int64) [][]int {
	max := int64(len(arr))
	//判断数组大小是否小于等于指定分割大小的值，是则把原数组放入二维数组返回
	if max <= num {
		return [][]int{arr}
	}
	//获取应该数组分割为多少份
	var quantity int64
	if max%num == 0 {
		quantity = max / num
	} else {
		quantity = (max / num) + 1
	}
	//声明分割好的二维数组
	var segments = make([][]int, 0)
	//声明分割数组的截止下标
	var start, end, i int64
	for i = 1; i <= quantity; i++ {
		end = i * num
		if i != quantity {
			segments = append(segments, arr[start:end])
		} else {
			segments = append(segments, arr[start:])
		}
		start = i * num
	}
	return segments
}

//string分组
func ArrayStrGroupsOf(arr []string, num int64) [][]string {
	max := int64(len(arr))
	//判断数组大小是否小于等于指定分割大小的值，是则把原数组放入二维数组返回
	if max <= num {
		return [][]string{arr}
	}
	//获取应该数组分割为多少份
	var quantity int64
	if max%num == 0 {
		quantity = max / num
	} else {
		quantity = (max / num) + 1
	}
	//声明分割好的二维数组
	var segments = make([][]string, 0)
	//声明分割数组的截止下标
	var start, end, i int64
	for i = 1; i <= quantity; i++ {
		end = i * num
		if i != quantity {
			segments = append(segments, arr[start:end])
		} else {
			segments = append(segments, arr[start:])
		}
		start = i * num
	}
	return segments
}

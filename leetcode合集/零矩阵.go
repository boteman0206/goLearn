package main

import "fmt"

/**
编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。
	输入：
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
输出：
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]
https://leetcode.cn/problems/zero-matrix-lcci/
*/

func setZeroes(matrix [][]int) {

	var data [][]int
	for i := range matrix {
		ints := matrix[i]
		for is := range ints {
			if ints[is] == 0 {
				data = append(data, []int{i, is})
			}
		}
	}

	for _, v := range data {
		fmt.Println(v)
		for i := range matrix {
			ints := matrix[i]
			if v[0] == i {
				for i2 := range ints {
					ints[i2] = 0
				}
			}
			for i2 := range ints {
				if i2 == v[1] {
					ints[i2] = 0
				}
			}
		}
	}

}

func main() {

	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix)
	fmt.Println(matrix)
}

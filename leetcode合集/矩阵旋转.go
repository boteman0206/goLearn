package main

import "fmt"

/**
给你一幅由 N × N 矩阵表示的图像，其中每个像素的大小为 4 字节。请你设计一种算法，将图像旋转 90 度。

不占用额外内存空间能否做到？
https://leetcode.cn/problems/rotate-matrix-lcci/

给定 matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

原地旋转输入矩阵，使其变为:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/rotate-matrix-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

func rotate(matrix [][]int) {

}

func main() {

	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)

	var data uint8
	data = 78
	var data1 int8
	data1 = -67

	fmt.Println(data, data1)

}

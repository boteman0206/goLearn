package main

import (
	"fmt"
	"sort"
)

/**
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。

示例 1：
输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	fmt.Println(nums1)
	var idata float64
	if len(nums1)%2 == 0 {
		index1 := len(nums1)/2 - 1
		index2 := len(nums1) / 2
		idata = float64(nums1[index1]+nums1[index2]) / float64(2)

	} else {
		i := (len(nums1) - 1) / 2
		idata = float64(nums1[i])
	}

	return idata
}

func main() {

	nums1 := []int{1, 3}
	nums2 := []int{2}
	arrays := findMedianSortedArrays(nums1, nums2)
	fmt.Println(arrays)

	fmt.Println(float64(5) / float64(2))

}

package main

import "fmt"

/**
	编写一个函数，检查输入的链表是否是回文的。
示例 1：

输入： 1->2
输出： false
示例 2：

输入： 1->2->2->1
输出： true

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/palindrome-linked-list-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {

	// 使用数组来实现
	data := []int{}
	for true {
		if head != nil {
			data = append(data, head.Val)
			head = head.Next
		} else {
			break
		}
	}

	//fmt.Println(data, len(data))

	// 找中间节点,减少循环时间
	if len(data)%2 == 0 {
		middle := (len(data)) / 2
		for i := 0; i < middle; i++ {
			//fmt.Println(data[i], "======", data[len(data)-i-1])
			if data[i] != data[len(data)-i-1] {
				return false
			}
		}

	} else {
		middle := (len(data) - 1) / 2
		//i2 := data[i] // 中间的值
		for i := 0; i < middle; i++ {
			//fmt.Println(data[i], "======", data[len(data)-1-i])
			if data[i] != data[len(data)-i-1] {
				return false
			}
		}
	}

	return true
}

func main() {

	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
	}
	palindrome := isPalindrome(l2)
	fmt.Println(palindrome)
}

package main

import "fmt"

/**

实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。

注意：本题相对原题稍作改动

示例：

输入： 1->2->3->4->5 和 k = 2
输出： 4

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/kth-node-from-end-of-list-lcci
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

func kthToLast(head *ListNode, k int) int {

	var data []int

	for true {
		if head == nil {
			break
		} else {
			data = append(data, head.Val)
		}

		head = head.Next

	}

	return data[len(data)-k]

}

func main() {
	//1->2->3->4->5
	l := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					4,
					&ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	last := kthToLast(&l, 2)

	fmt.Println(last)
}

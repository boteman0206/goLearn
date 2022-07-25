package main

import (
	"encoding/json"
	"fmt"
)

/**

编写代码，移除未排序链表中的重复节点。保留最开始出现的节点
https://leetcode.cn/problems/remove-duplicate-node-lcci/


输入：[1, 2, 3, 3, 2, 1]
输出：[1, 2, 3]

输入：[1, 1, 1, 1, 2]
输出：[1, 2]

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	occurred := map[int]bool{head.Val: true}
	pos := head
	for pos.Next != nil {
		cur := pos.Next
		if !occurred[cur.Val] {
			occurred[cur.Val] = true
			pos = pos.Next
		} else {
			pos.Next = pos.Next.Next
		}
	}
	//pos.Next = nil
	return head
}

func main() {
	//输入：[1, 2, 3, 3, 2, 1]
	l := new(ListNode)
	l.Val = 1
	l.Next = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}

	nodes := removeDuplicateNodes(l)

	marshal, err := json.Marshal(nodes)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
}

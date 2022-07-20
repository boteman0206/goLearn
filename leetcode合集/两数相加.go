package main

import (
	"encoding/json"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode.cn/problems/add-two-numbers/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//
//	var numL1 []int64
//	for true {
//		if l1 != nil {
//			numL1 = append(numL1, int64(l1.Val))
//		} else {
//			break
//		}
//		l1 = l1.Next
//	}
//
//	var numL2 []int64
//
//	for true {
//		if l2 != nil {
//			numL2 = append(numL2, int64(l2.Val))
//		} else {
//			break
//		}
//		l2 = l2.Next
//	}
//	fmt.Println(numL1, " === ", numL2)
//	//lenNum1 := len(numL1)
//	//lenNum2 := len(numL2)
//	//
//	//var allNum1 int64
//	//for i := range numL1 {
//	//	dumD := lenNum1 - i
//	//
//	//
//
//	//}
//
//	return nil
//}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	//定义一个尾结点，或者可以理解为临时节点
	var tail *ListNode
	//余数carry
	carry := 0
	//依次遍历两个链表，只要元素不为空就进行下一步
	for l1 != nil || l2 != nil {
		//定义两个变量存储各个节点的值
		n1, n2 := 0, 0
		//从第一个链表开始
		if l1 != nil {
			//把每个节点的值赋给n1
			n1 = l1.Val
			//节点后移
			l1 = l1.Next
		}
		//l2同上
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		//此时是两个链表第一个元素的和 + 余数
		sum := n1 + n2 + carry
		//sum%10是节点的当前值，如果是10,取余后当前节点值为0，sum/10是求十位的那个数
		sum, carry = sum%10, sum/10
		//此时申请一个新的链表存储两个链表的和
		if head == nil {
			//申请新的链表
			head = &ListNode{Val: sum}
			//这一步是为了保持头结点不变的情况下指针可以右移，所以说tail相当于临时节点，理解成尾节点也可以，因
			//为此时新链表中只有一个节点，所以头结点和尾结点都指向同一个元素。
			tail = head
		} else {
			//第二个节点后开始逐渐往尾结点增加元素
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	//把最后一位的余数加到链表最后。
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}

func main() {

	node1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	}

	node2 := ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val:  6,
				Next: nil,
			},
		},
	}

	numbers := addTwoNumbers(&node1, &node2)
	marshal, err := json.Marshal(numbers)
	if err != nil {
		return
	}
	fmt.Println(" numbers : ", string((marshal)))

	sum, carry := 10%10, 18/10
	fmt.Println(sum, carry)

}

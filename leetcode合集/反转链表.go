package main

import (
	"container/list"
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param pHead ListNode类
 * @return ListNode类
 */
func ReverseList(pHead *ListNode) *ListNode {
	// ret返回值
	var ret *ListNode
	// 节点为空，直接返回
	if pHead == nil {
		return ret
	}
	// cur为当前节点，pre为当前节点的前一个结点，next为当前节点的后一个节点
	// 需要让 pre 和 next 的目的是让当前节点从 pre->cur->next1->next2 变成 pre<-cur next1->next2
	//即 pre 让节点可以反转所指方向，但反转之后如果不用 next 节点保存 next1节点的话，此单链表就断开了
	var cur, pre, next *ListNode
	cur = pHead
	for cur != nil {
		//先用 next 节点保存 cur 的下一个结点的信息，保证单链表不会因为失去 cur节点的原next节点而就此断裂
		next = cur.Next
		//保存完 next，就可以让cur从指向next变成指向pre了
		cur.Next = pre
		//cur指向pre后，就继续依次反转下一个节点
		//让pre、cur、next依次向后移动一个节点，继续下一次的指针反转
		pre = cur
		cur = next
	}
	//如果cur为nil时，pre就为最后一个节点了，此时链表已经反转完毕，pre就是反转后链表的第一个节点
	ret = pre
	return ret
}

func main() {
	// 声明链表
	l := list.New()

	// 数据添加到尾部
	l.PushFront(4)
	l.PushFront(5)
	l.PushFront(6)

	// 遍历
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v\n", e.Value)
	}

	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	reverseList := ReverseList(l2)
	fmt.Println(jsoniter.MarshalToString(reverseList))

	var nums = []int{1, 2, 3, 4, 5}
	for i := range nums {
		fmt.Println(nums[len(nums)-i-1])
	}

}

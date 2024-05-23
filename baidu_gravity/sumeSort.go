package main

import (
	"fmt"
)

func main() {

	fmt.Println("-----", "run")

	maopao001 := maopaoSort([]int64{23, 31, 4343, 516, 3, 566, 56})
	fmt.Println("maopao001:", maopao001)

	fmt.Println("quicksort: ", quickSort([]int64{23, 31, 4343, 516, 3, 566, 56}))

	fmt.Println("selectSort:", selectSort([]int64{23, 31, 4343, 516, 3, 566, 56}))

	ll := LinkedList{}
	ll.InsertAtEnd(3)
	ll.InsertAtEnd(4)
	ll.InsertAtEnd(1)
	ll.InsertAtEnd(2)

	ll.DisplayNode()

	ll.DeleteNode(3)
	fmt.Println("----ll..DeleteNode()----")
	ll.DisplayNode()

}

func maopaoSort(list []int64) []int64 {
	if len(list) <= 1 {
		return list
	}
	length := len(list)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
}

func quickSort(list []int64) []int64 {

	if len(list) <= 1 {
		return list
	}

	provid := list[0]
	var left, right []int64

	for _, v := range list[1:] {
		if v <= provid {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, provid), right...)
}

func selectSort(list []int64) []int64 {

	if len(list) <= 0 {
		return list
	}

	for i := 0; i < len(list)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(list); j++ {
			if list[j] < list[minIndex] {
				minIndex = j
			}
		}
		list[i], list[minIndex] = list[minIndex], list[i]
	}
	return list
}

type ListNode struct {
	Val  int64
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func (ll *LinkedList) InsertAtEnd(val int64) {

	nowNode := ListNode{
		Val: val,
	}
	if ll.Head == nil {
		ll.Head = &nowNode
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = &nowNode

}

func (ll *LinkedList) DeleteNode(val int64) {
	if ll.Head == nil {
		return
	}

	if ll.Head.Val == val {
		ll.Head = ll.Head.Next
		return
	}

	preNode := ll.Head
	for preNode.Next != nil {
		if preNode.Next.Val == val {
			preNode.Next = preNode.Next.Next
			return
		}

		preNode = preNode.Next
	}

}

func (ll *LinkedList) DisplayNode() {

	if ll.Head == nil {
		return
	}

	currentNode := ll.Head
	for currentNode != nil {
		fmt.Println("currentNode.Val: ", currentNode.Val)
		currentNode = currentNode.Next
	}

}

type DubbleNode struct {
	Val  int64
	Pre  *DubbleNode
	Next *DubbleNode
}

type DoubbleList struct {
	Head *DubbleNode
	Tail *DubbleNode
}

func (du *DoubbleList) InsertAtEnd(val int64) {

	newNode := DubbleNode{
		Val:  val,
		Pre:  nil,
		Next: nil,
	}
	if du.Head == nil {
		du.Head = &newNode
		du.Tail = &newNode
		return
	}

	newNode.Pre = du.Tail
	du.Tail.Next = &newNode
	du.Tail = &newNode
}

func (du *DoubbleList) DisplayNodePre() {

}

func (du *DoubbleList) DeleteNode(val int64) {

}

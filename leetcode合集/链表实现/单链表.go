package main

import (
	"fmt"
	"sync"
)

/**
单向链表（单链表）是链表的一种，其特点是链表的链接方向是单向的，对链表的访问要通过顺序读取从头部开始；链表是使用指针进行构造的列表；又称为结点列表，因为链表是由一个个结点组装起来的；其中每个结点都有指针成员变量指向列表中的下一个结点；
列表是由结点构成，head指针指向第一个成为表头结点，而终止于最后一个指向nuLL的指针。

*/

//单链表的节点
type SingleNode struct {
	Data interface{}
	Next *SingleNode
}

//单链表
type SingleList struct {
	lock *sync.Mutex
	Head *SingleNode // 头结点
	Tail *SingleNode // 尾节点
	Size int         // 链表的长度
}

func NewList() *SingleList {
	return &SingleList{
		lock: new(sync.Mutex),
		Head: nil,
		Tail: nil,
		Size: 0,
	}
}

func (list *SingleList) Append(node *SingleNode) bool {
	if node == nil {
		return false
	}
	list.lock.Lock()
	defer list.lock.Unlock()

	if list.Size == 0 {
		list.Head = node
		list.Tail = node
		list.Size = 1

	} else {
		tail := list.Tail
		tail.Next = node
		list.Tail = node
		list.Size += 1
	}
	return true
}

// 插入节点到指定位置
func (list *SingleList) Insert(index int, node *SingleNode) bool {
	if node == nil {
		return false
	}

	if index > list.Size {
		return false
	}

	list.lock.Lock()
	defer list.lock.Unlock()

	if index == 0 {
		head := list.Head
		list.Head = node
		list.Size += 1
		node.Next = head
		return true
	}

	ptr := list.Head
	for i := 1; i < index; i++ {
		ptr = ptr.Next
	}
	next := ptr.Next
	node.Next = next

	ptr.Next = node
	list.Size += 1

	return true

}

// 删除指定位置的节点
func (list *SingleList) Delete(index int) bool {
	if list == nil || list.Size == 0 || index > list.Size-1 {
		return false
	}

	list.lock.Lock()
	defer list.lock.Unlock()

	if index == 0 {
		next := list.Head.Next
		list.Head = next
		if list.Size == 1 {
			list.Tail = nil
		}
		list.Size -= 1
		return true
	}
	ptr := list.Head
	for i := 1; i < index; i++ {
		ptr = ptr.Next
	}
	next := ptr.Next
	ptr.Next = next.Next
	if index == list.Size-1 {
		list.Tail = ptr
	}
	list.Size -= 1

	return true
}

// 输出链表
func (list *SingleList) Display() {
	if list == nil {
		fmt.Println("this single list is nil")
		return
	}
	list.lock.Lock()
	defer list.lock.Unlock()
	fmt.Printf("this single list size is %d \n", list.Size)
	ptr := list.Head
	var i int
	for i = 0; i < list.Size; i++ {
		fmt.Printf("No%3d data is %v\n", i+1, ptr.Data)
		ptr = ptr.Next
	}
}

func main() {

	list := NewList()
	list.Append(&SingleNode{
		"a",
		nil,
	})
	list.Append(&SingleNode{
		Data: "b",
		Next: nil,
	})
	list.Append(&SingleNode{
		Data: "c",
		Next: nil,
	})

	list.Insert(2, &SingleNode{
		Data: "e",
		Next: nil,
	})
	list.Display()

	println("=================")
	list.Delete(2)
	list.Display()
}

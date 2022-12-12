package main

import (
	"fmt"
	"sync"
)

/**
双向链表也叫双链表，是链表的一种，它的每个数据结点中都有两个指针，分别指向直接后继和直接前驱。所以，从双向链表中的任意一个结点开始，都可以很方便地访问它的前驱结点和后继结点。
*/

//节点的数据
type DoubleNode struct {
	Data string
	Prev *DoubleNode
	Next *DoubleNode
}

// 链表结构

type DoubleList struct {
	mutex *sync.Mutex

	Size int
	Head *DoubleNode // 头结点
	Tail *DoubleNode // 尾节点
}

func NewDoubleList() *DoubleList {
	return &DoubleList{
		mutex: &sync.Mutex{},
		Size:  0,
		Head:  nil,
		Tail:  nil,
	}
}

//双向链表也叫双链表，是链表的一种，它的每个数据结点中都有两个指针，分别指向直接后继和直接前驱。所以，从双向链表中的任意一个结点开始，都可以很方便地访问它的前驱结点和后继结点。
func (d *DoubleList) Append(node *DoubleNode) bool {
	if node == nil {
		return false
	}
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if d.Size == 0 {

		d.Head = node
		d.Tail = node

		node.Next = nil
		node.Prev = nil

	} else {

		tail := d.Tail // 尾节点

		tail.Next = node // 尾节点的下一个指向node

		node.Prev = tail // node的上一个指向tail
		node.Next = nil  // node的下一个为nil
		d.Tail = node    // 当前的tail为node

	}
	d.Size += 1
	return true
}

// 插入指定的位置元素
func (d *DoubleList) Insert(index int, node *DoubleNode) bool {
	if index > d.Size || node == nil {
		return false
	}

	if index == d.Size {
		return d.Append(node)
	}

	d.mutex.Lock()
	defer d.mutex.Unlock()

	if index == 0 {
		node.Next = d.Head
		node.Prev = nil
		node.Next.Prev = node // 设置双向链表节点

		d.Head = node
		d.Size += 1
		return true
	}

	nextNode := d.Get(index)

	node.Prev = nextNode.Prev
	node.Next = nextNode

	nextNode.Prev.Next = node
	nextNode.Prev = node

	d.Size += 1
	return true
}

func (d *DoubleList) Get(index int) *DoubleNode {
	if d.Size == 0 || index > d.Size-1 {
		return nil
	}
	if index == 0 {
		return d.Head
	}

	head := d.Head

	for i := 0; i < index; i++ {
		head = head.Next
	}
	return head
}

func (list *DoubleList) Display() {
	if list == nil || list.Size == 0 {
		fmt.Println("this double list is nil or empty")
		return
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	fmt.Printf("this double list size is %d \n", list.Size)
	ptr := list.Head
	for ptr != nil {
		fmt.Printf("data is %v\n", ptr.Data)
		ptr = ptr.Next
	}
}

func (d *DoubleList) Delete(index int) bool {
	if index > d.Size || d.Size == 0 {
		return false
	}

	d.mutex.Lock()
	defer d.mutex.Unlock()

	if index == 0 {
		if d.Size == 1 {
			d.Head = nil
			d.Tail = nil

		} else {
			next := d.Head.Next // 头节点的下一个节点
			next.Prev = nil

			d.Head = next // 提升为head节点
		}
		d.Size--
		return true
	}

	if index == d.Size-1 { // 删除最后一个节点
		prev := d.Tail.Prev
		prev.Next = nil

		d.Tail = prev
		d.Size--
		return true
	}

	node := d.Get(index)
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev

	d.Size--
	return true
}
func main() {
	list := NewDoubleList()

	//list.Append(&DoubleNode{
	//	Data: "a",
	//	Prev: nil,
	//	Next: nil,
	//})
	//
	//list.Append(&DoubleNode{
	//	Data: "b",
	//	Prev: nil,
	//	Next: nil,
	//})
	//list.Append(&DoubleNode{
	//	Data: "c",
	//	Prev: nil,
	//	Next: nil,
	//})
	list.Insert(0, &DoubleNode{
		Data: "c",
		Prev: nil,
		Next: nil,
	})
	list.Insert(0, &DoubleNode{
		Data: "b",
		Prev: nil,
		Next: nil,
	})
	list.Insert(0, &DoubleNode{
		Data: "a",
		Prev: nil,
		Next: nil,
	})

	list.Delete(2)
	list.Display()

}

package main

import "fmt"

type MyQueue struct {
	Data []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{Data: []int{}}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.Data = append(this.Data, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	ft := this.Data[1:len(this.Data)]
	var front int
	if len(this.Data) > 0 {
		front = this.Data[0]
		this.Data = ft
	}
	return front
}

/** Get the front element. */
func (this *MyQueue) Peek() int {

	var front int
	if len(this.Data) > 0 {
		front = this.Data[0]
	}
	return front
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.Data) > 0 {
		return false
	}
	return true
}

func main() {

	constructor := Constructor()
	constructor.Push(1)
	constructor.Push(2)
	pop := constructor.Pop()
	fmt.Println(pop)
	peek := constructor.Peek()
	fmt.Println(peek)

}

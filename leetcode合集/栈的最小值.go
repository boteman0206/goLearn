package main

import "sync"

type MinStack struct {
	Lock sync.Mutex
	Data []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Lock: sync.Mutex{},
		Data: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.Lock.Lock()
	defer this.Lock.Unlock()
	this.Data = append(this.Data, x)

}

func (this *MinStack) Pop() {

	this.Lock.Lock()
	defer this.Lock.Unlock()
	this.Data = this.Data[:len(this.Data)-1]

}

func (this *MinStack) Top() int {
	this.Lock.Lock()
	defer this.Lock.Unlock()

	return this.Data[len(this.Data)-1]
}

func (this *MinStack) GetMin() int {
	this.Lock.Lock()
	defer this.Lock.Unlock()
	var min int
	if len(this.Data) > 0 {
		min = this.Data[0]
	}
	for i := range this.Data {
		if this.Data[i] < min {
			min = this.Data[i]
		}
	}
	return min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
func main() {

}

package main

import (
	"errors"
	"fmt"
	"sync"
)

/**
	请你实现一个队列。 先进先出
操作：
push x：将 x\x 加入队尾，保证 x\x 为 int 型整数。
pop：输出队首，并让队首出队
front：输出队首：队首不出队

*/

type list struct {
	list []int
	lock sync.Mutex
}

func (l *list) push(data int) {
	l.list = append(l.list, data)
}

func (l *list) pop() (int, error) {
	l.lock.Lock()
	defer l.lock.Unlock()

	var data int
	if len(l.list) > 0 {
		data = l.list[0]
		l.list = l.list[1:]
		return data, nil
	}

	return data, errors.New("err")
}

func (l *list) front() (int, error) {
	l.lock.Lock()
	defer l.lock.Unlock()

	if len(l.list) > 0 {
		return l.list[0], nil
	}
	return 0, errors.New("err")
}

func main() {
	l := list{
		list: nil,
		lock: sync.Mutex{},
	}

	//l.push(45)
	//l.push(46)
	//l.push(47)
	//
	//pop, err := l.pop()
	//fmt.Println(pop, err)
	//pop1, err := l.pop()
	//fmt.Println(pop1, err)
	//pop2, err := l.pop()
	//fmt.Println(pop2, err)
	//pop3, err := l.pop()
	//fmt.Println(pop3, err)

	//fro, err := l.front()
	//fmt.Println(fro, err)
	//fro1 := l.front()
	//fmt.Println(fro1)
	//fro2 := l.front()
	//fmt.Println(fro2)
	//fro3 := l.front()
	//fmt.Println(fro3)

	fro, err := l.front()
	fmt.Println(fro, err)
	pop3, err := l.pop()
	fmt.Println(pop3, err)
	l.push(7)
	front, err := l.front()
	fmt.Println(front, err)

}

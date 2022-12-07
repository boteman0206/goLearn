package main

import (
	"fmt"
	"sync"
)

var wait = &sync.WaitGroup{}
var c1 = make(chan struct{}, 1)
var c2 = make(chan struct{}, 1)
var c3 = make(chan struct{}, 1)

func T1() {

	for i := 0; i < 5; i++ {
		<-c1
		fmt.Println("a")

		c2 <- struct{}{}
	}

	wait.Done()
}

func T2() {

	for i := 0; i < 5; i++ {
		<-c2
		fmt.Println("b")

		c3 <- struct{}{}
	}

	wait.Done()
}

func T3() {

	for i := 0; i < 5; i++ {
		<-c3
		fmt.Println("c")

		c1 <- struct{}{}
	}

	wait.Done()

}

func main() {

	wait.Add(3)

	go T1()
	go T2()
	go T3()
	c1 <- struct{}{}

	wait.Wait()

}

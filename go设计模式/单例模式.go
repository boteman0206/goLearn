package main

import (
	"sync"
	"sync/atomic"
)

type Instance struct{}

var instance *Instance

// todo 1：非线程安全
func NewInstance1() *Instance {
	if instance == nil {
		instance = &Instance{}
	}
	return instance
}

var lock sync.Mutex

// todo 2：check-lock模式
func NewInstance2() *Instance {
	lock.Lock()
	defer lock.Unlock()

	if instance != nil {
		instance = &Instance{}
	}

	return instance
}

// todo 3: check-lock-check模式
func NewInstance3() *Instance {

	if instance == nil {
		lock.Lock()
		if instance == nil {
			instance = &Instance{}
		}
		lock.Unlock()
	}

	return instance
}

// todo 4:原子操作 check-lock-check
var initlized int32

func NewInstance4() *Instance {
	if atomic.LoadInt32(&initlized) == 1 {
		return instance
	}
	lock.Lock()
	defer lock.Unlock()

	if initlized == 0 {
		instance = &Instance{}
		atomic.StoreInt32(&initlized, 1)
	}

	return instance
}

// todo 5: once方法
var once sync.Once

func NewInstance5() *Instance {
	if instance == nil {
		once.Do(func() {
			instance = &Instance{}
		})
	}
	return instance
}

package main

import "sync"

type Obj struct {
	mu sync.Mutex
	// ... 其他字段
}

func (o *Obj) Lock()        { o.mu.Lock() }
func (o *Obj) Dosomething() {}
func (o *Obj) Unlock()      { o.mu.Unlock() }

func main() {
	o := Obj{}

	o.Lock()
	o.Dosomething()
	o.Unlock()

	o.Lock()
	o.Dosomething()
	o.Unlock()
}

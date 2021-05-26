package main

import (
	"fmt"
	"sync"
)

type Stats struct {
	mu sync.Mutex

	counters map[string]int
}

func (s *Stats) Snapshot() map[string]int {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make(map[string]int, len(s.counters))
	for k, v := range s.counters {
		result[k] = v
	}
	return result
}

func main() {
	// snapshot 现在是一个拷贝 对snapshot的修改不会影响counters

	var stats Stats
	stats.counters = make(map[string]int)
	stats.counters = map[string]int{
		"name": 1,
		"age":  19,
		"addr": 123,
	}

	fmt.Println("修改前： counters ： ", stats.counters)
	snapshot := stats.Snapshot()
	snapshot["name"] = 12

	fmt.Println("修改后面 snapshot： ", snapshot)
	fmt.Println("修改后面 counters： ", stats.counters)

}

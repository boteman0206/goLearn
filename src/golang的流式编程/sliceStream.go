package main

import (
	"fmt"
	"sync"
)

type int32Filter struct {
	lock sync.Mutex
	data []int32
}

func NewInt32Filter(item []int32) *int32Filter {
	s := &int32Filter{
		data: make([]int32, 0),
	}

	s.data = append(s.data, item...)

	return s
}

//过滤等小于
func (s *int32Filter) FilterLtNum(num int32) *int32Filter {
	s.lock.Lock()
	defer s.lock.Unlock()

	var res []int32
	for _, v := range s.data {
		v := v
		if v < num {
			res = append(res, v)
		}
	}
	s.data = res

	return s
}

//过滤大于
func (s *int32Filter) FilterGtNum(num int32) *int32Filter {
	s.lock.Lock()
	defer s.lock.Unlock()
	var res []int32
	for _, v := range s.data {
		v := v
		if v > num {
			res = append(res, v)
		}
	}
	s.data = res

	return s
}

// 过滤等于
func (s *int32Filter) FilterEtNum(num int32) *int32Filter {
	s.lock.Lock()
	defer s.lock.Unlock()
	var res []int32
	for _, v := range s.data {
		v := v
		if v == num {
			res = append(res, v)
		}
	}
	s.data = res

	return s
}

// 截取到len的长度
func (s *int32Filter) FilterLen(len int32) *int32Filter {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.data = s.data[:len]
	return s
}

// 数组去重
func (s *int32Filter) RemoveDuplicate() *int32Filter {
	s.lock.Lock()
	defer s.lock.Unlock()
	var mapt = make(map[int32]struct{}, 0)

	for _, v := range s.data {
		v := v
		mapt[v] = struct{}{}
	}

	var data []int32
	for k := range mapt {
		data = append(data, k)
	}

	s.data = data
	return s
}

// 分组
func (s *int32Filter) GroupBy() map[int32]int {
	s.lock.Lock()
	defer s.lock.Unlock()

	var data = make(map[int32]int, 0)

	for _, v := range s.data {
		v := v
		if num, ok := data[v]; ok {
			data[v] = num + 1
		} else {
			data[v] = 1
		}
	}
	return data
}

func (s *int32Filter) Max() int32 {
	s.lock.Lock()
	defer s.lock.Unlock()

	var maxData int32
	for _, v := range s.data {
		v := v
		if maxData < v {
			maxData = v
		}
	}

	return maxData
}
func (s *int32Filter) Min() int32 {
	s.lock.Lock()
	defer s.lock.Unlock()

	var minData int32
	if len(s.data) > 1 {
		minData = s.data[1]
	}
	for _, v := range s.data {
		v := v
		if minData > v {
			minData = v
		}
	}

	return minData
}

// 平均值 会丢失精度
func (s *int32Filter) AVG() float64 {
	s.lock.Lock()
	defer s.lock.Unlock()

	var sum int32
	for _, v := range s.data {
		v := v
		sum += v
	}

	return float64(sum) / float64(len(s.data))
}

// 返回int数组
func (s *int32Filter) ToInt32() []int32 {
	return s.data
}

func main() {

	var data = []int32{1, 2, 310, 4, 4, 5, 6, 67, 7, 87}

	filter := NewInt32Filter(data).FilterGtNum(3).ToInt32()
	fmt.Println("过滤大于3的数据：", filter)

	toInt32 := NewInt32Filter(data).FilterEtNum(4).ToInt32()
	fmt.Println("过滤等于4的数据： ", toInt32)

	//分组
	by := NewInt32Filter(data).GroupBy()
	fmt.Println("分组by :", by)

	// 求平均值
	avg := NewInt32Filter(data).AVG()
	fmt.Println("avg: ", avg)

	// 求最大值
	max := NewInt32Filter(data).Max()
	fmt.Println("max: ", max)

	// 最小值
	min := NewInt32Filter(data).Min()
	fmt.Println("min: ", min)

}

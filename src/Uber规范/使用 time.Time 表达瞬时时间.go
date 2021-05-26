package main

import "time"

/**
todo 在处理时间的瞬间时使用 time.Time，在比较、添加或减去时间时使用 time.Time 中的方法。

 */

// todo 错误的使用方式
//func isActive(now, start, stop int) bool {
//	return start <= now && now < stop
//}

// 正确的使用方式
func isActive(now, start, stop time.Time) bool {
	return (start.Before(now) || start.Equal(now)) && now.Before(stop)
}


/**
todo 错误用法
func poll(delay int) {
	for {
		// ...
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}
poll(10) // 是几秒钟还是几毫秒?

 */

// 正确用法
func poll(delay time.Duration) {
	for {
		// ...
		time.Sleep(delay)
	}
}
poll(10*time.Second)


/**
回到第一个例子，在一个时间瞬间加上 24 小时，我们用于添加时间的方法取决于意图。如果我们想要下一个日历日(当前天的下一天)的同一个时间点，
我们应该使用 Time.AddDate。但是，如果我们想保证某一时刻比前一时刻晚 24 小时，我们应该使用 Time.Add。


 */

newDay := t.AddDate(0 /* years */, 0 /* months */, 1 /* days */)
maybeNewDay := t.Add(24 * time.Hour)





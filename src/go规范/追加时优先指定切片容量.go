package main

/**

追加时优先指定切片容量
在尽可能的情况下，在初始化要追加的切片时为make()提供一个容量值。
*/

func main() {

	/**
		for n := 0; n < b.N; n++ {
		//todo 错误使用方式
		data := make([]int, 0)
		for k := 0; k < size; k++{
			data = append(data, k)
		}
	}

	//BenchmarkBad-4    100000000    2.48s

	*/

	/**
	for n := 0; n < b.N; n++ {
		// todo 正确使用方式
	  data := make([]int, 0, size)
	  for k := 0; k < size; k++{
	    data = append(data, k)
	  }
	}


	BenchmarkGood-4   100000000    0.21s
	*/

}

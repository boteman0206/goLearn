package main

import "fmt"

/*
	最灵活性的方法: fmt.Sscanf(…)
	fmt.Sscanf为格式字符串提供了更大的灵活性, 您可以在输入中指定数字格式（例如宽度,基数等）以及其他额外的字符string,
*/
func main() {
	s := "id:00123"

	var i int
	if _, err := fmt.Sscanf(s, "id:%5d", &i); err == nil {
		fmt.Println(i) // Outputs 123
	}

	var in int
	sscanf, err := fmt.Sscanf("123", "%d", &in)
	if err != nil {
		return
	}

	fmt.Println(sscanf, " in : ", in, in+1)

}

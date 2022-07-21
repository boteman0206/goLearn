package main

import "fmt"

//https://leetcode.cn/problems/is-unique-lcci/

func isUnique(astr string) bool {

	m := make(map[string]struct{})

	var flaa bool = true
	for i := range astr {
		str := string(astr[i])
		if _, ok := m[str]; !ok {
			m[str] = struct{}{}
		} else {
			flaa = false
		}
	}

	return flaa
}

func main() {
	hule := "helloworld"
	isUnique(hule)

	for i := range hule {
		fmt.Println(string(hule[i]))
	}
}

package main

//go get github.com/rs/xid

import (
	"fmt"
	"github.com/rs/xid"
)

/*
*
可排序的id生成
*/
func main() {

	var xids []xid.ID
	for i := 0; i < 1000; i++ {
		guid := xid.New()
		println(guid.String())

		xids = append(xids, guid)
	}
	// Output: 9m4e2mr0ui3e8a215n4g

	// 排序
	xid.Sort(xids)

	for i := range xids {
		fmt.Println("====", xids[i].String())
	}

}

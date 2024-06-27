package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// 生成格式化的sql
/**
**表按照user_id来进行分表查询的，将一批user_id分组合并起来，这样可以减少查询的次数
**
 */

// 3: 按照userID分表，分成组
func getTableName(userIDs []int64) ([]int64, map[int64][]int64) {

	mapv1 := make(map[int64]int64, 0)
	mapv2 := make(map[int64][]int64, 0)
	tableIDs := []int64{}
	fmt.Println("len user_ids: ", len(userIDs))
	for i := range userIDs {
		userId := userIDs[i]
		tableID := getTableID(userId)
		mapv2[tableID] = append(mapv2[tableID], userId)
		if _, ok := mapv1[tableID]; !ok {
			tableIDs = append(tableIDs, tableID)
			mapv1[tableID] = 0
		}
	}
	byteSID, _ := json.Marshal(tableIDs)
	userTablestr, _ := json.Marshal(mapv2)
	fmt.Println("生成的表id: ", string(byteSID))
	fmt.Println("分组的userIDs: ", string(userTablestr))
	return tableIDs, mapv2
}

// 4 生成sql
func getTableScript(tablePrefix, queryTemplate string, tableids []int64, mapTabIdUserID map[int64][]int64) {
	queries := make([]string, 0)

	// 生成每个表的查询
	var lenUid = 0
	for i := range tableids {
		uids := mapTabIdUserID[tableids[i]]
		userIDStrings := make([]string, 0)
		lenUid += len(uids)
		for k := range uids {
			userIDStrings = append(userIDStrings, fmt.Sprintf("%d", uids[k]))
		}
		// 将字符串数组连接为一个用逗号分隔的字符串
		inClause := strings.Join(userIDStrings, ", ")
		fmt.Println("inClause: lenUid :", lenUid, "tableID: ", tableids[i], " alluserID: ", inClause)

		// 将所有查询用 " UNION ALL " 连接
		strTemp := queryTemplate + fmt.Sprintf(` and user_id in (%s)`, inClause)

		queries = append(queries, fmt.Sprintf(strTemp, tablePrefix, tableids[i]))
	}

	finalQuery := strings.Join(queries, " UNION ALL\n") + ";"

	fmt.Println("----------start-------------------")
	// 打印生成的查询
	fmt.Println(finalQuery)

	fmt.Println("----------end-------------------")

}

// 1: 定义userID
var ALLuserIDs = []int64{
	1531008566, 1531018576, 1531028579, 1531052863, 1531063255, 1531092017,
	1531092018, 1531093110, 1531093114, 1531093117, 1537581653, 1537583931, 1537584468,
	1537584606, 1537584682, 1537585321, 1537585348, 1537585699, 1537585731, 1537585736,
	1537585758, 1537585798, 1537585814, 1537585820, 1537585866, 1537585911, 1537585944,
	1537585946, 1537585985, 1537585988, 1537585994, 1537586030, 1537586347, 1537586387,
	1537586392, 1537586410, 1537586427, 1537586433, 1537586458, 1537586490,
	1537586497, 1537586506, 1537586547, 1537586568, 1537586752,
}

// 2: 获取分表的规则 需要动态的调整
func getTableID(userID int64) int64 {
	return (userID % 10) + 1
}

func main() {

	tablePrefixId, tablegroupUserID := getTableName(ALLuserIDs)
	// 定义表名前缀和总表数量
	tablePrefixstr := "tb_user_bind_"
	// 定义查询模板
	queryTemplate := "SELECT user_id, idfa FROM `%s%d` WHERE "
	getTableScript(tablePrefixstr, queryTemplate, tablePrefixId, tablegroupUserID)
}

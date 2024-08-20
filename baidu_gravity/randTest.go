package main

import (
	"encoding/json"
	"fmt"
)

type DepositRewardInfo struct {
	RewardType string `json:"reward_type"` // 奖励类型：coin_package：充值套餐
	RewardID   string `json:"reward_id"`   // 奖励ID
	ExpireTime int64  `json:"expire_time"` // 有效期
	CoinValue  int64  `json:"coin_value "` // 价值星粒数
	Count      int64  `json:"count"`       // 奖励数量
}

func main() {

	t := []DepositRewardInfo{{
		RewardType: "frame",
		RewardID:   "10",
		ExpireTime: 7,
		CoinValue:  78,
		Count:      1,
	}}
	bytes, _ := json.Marshal(t)
	fmt.Println(string(bytes))

}

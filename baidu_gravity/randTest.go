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
	}, {
		RewardType: "user_tag",
		RewardID:   "18",
		ExpireTime: 4,
		CoinValue:  48,
		Count:      1,
	}, {
		RewardType: "prize",
		RewardID:   "66",
		ExpireTime: 3,
		CoinValue:  34,
		Count:      1,
	}, {
		RewardType: "coin_package",
		RewardID:   "gravitycoin_50coin_120jpy",
		ExpireTime: 0,
		CoinValue:  0,
		Count:      0,
	}}
	bytes, _ := json.Marshal(t)
	fmt.Println(string(bytes))

}

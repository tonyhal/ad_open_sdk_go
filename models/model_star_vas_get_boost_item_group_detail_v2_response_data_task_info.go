/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarVasGetBoostItemGroupDetailV2ResponseDataTaskInfo 助推组信息
type StarVasGetBoostItemGroupDetailV2ResponseDataTaskInfo struct {
	// 手动出价，乘百取整（如：0.1 -> 10）。若没有则代表自动出价
	BidAmount   *int64                                                  `json:"bid_amount,omitempty"`
	BidType     *StarVasGetBoostItemGroupDetailV2DataTaskInfoBidType    `json:"bid_type,omitempty"`
	BoostAction StarVasGetBoostItemGroupDetailV2DataTaskInfoBoostAction `json:"boost_action"`
	// 助推预算，单位元
	BoostAmount int64 `json:"boost_amount"`
	// 助推消耗，单位千分之一元
	BoostCost int64 `json:"boost_cost"`
	// 自定义投放时长，单位时。若没有则代表非自定义投放时长
	BoostHours *int64                                                `json:"boost_hours,omitempty"`
	BoostType  StarVasGetBoostItemGroupDetailV2DataTaskInfoBoostType `json:"boost_type"`
	// 创建时间
	CreateTime int64 `json:"create_time"`
	// 关联人群包，仅boost_type为CUSTOM_PACKAGE有效
	PackId *int64                                             `json:"pack_id,omitempty"`
	Status StarVasGetBoostItemGroupDetailV2DataTaskInfoStatus `json:"status"`
	// 助推组ID
	TaskId int64 `json:"task_id"`
	// 助推组名称
	TaskName string `json:"task_name"`
}

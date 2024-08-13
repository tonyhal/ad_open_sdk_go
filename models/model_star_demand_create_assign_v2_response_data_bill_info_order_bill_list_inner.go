/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandCreateAssignV2ResponseDataBillInfoOrderBillListInner struct for StarDemandCreateAssignV2ResponseDataBillInfoOrderBillListInner
type StarDemandCreateAssignV2ResponseDataBillInfoOrderBillListInner struct {
	// 任务ID
	OrderId *int64 `json:"order_id,omitempty"`
	// 服务费
	PlatformFee *int64 `json:"platform_fee,omitempty"`
	// 精确服务费
	PrecisePlatformFee *float64 `json:"precise_platform_fee,omitempty"`
	// 精确待付金额
	PreciseRemaining *float64 `json:"precise_remaining,omitempty"`
	// 精确任务金额
	PreciseTaskCost *float64 `json:"precise_task_cost,omitempty"`
	// 精确总金额（单位元，下同）
	PreciseTotal *float64 `json:"precise_total,omitempty"`
	// 待付金额
	Remaining *int64 `json:"remaining,omitempty"`
	// 任务金额
	TaskCost *int64 `json:"task_cost,omitempty"`
	// 总金额（单位元，下同）
	Total *int64 `json:"total,omitempty"`
}

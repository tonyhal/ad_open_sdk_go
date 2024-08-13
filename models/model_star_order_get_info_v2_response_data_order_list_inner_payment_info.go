/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarOrderGetInfoV2ResponseDataOrderListInnerPaymentInfo 资金信息
type StarOrderGetInfoV2ResponseDataOrderListInnerPaymentInfo struct {
	// 已扣除金额
	DeductAmount *int64 `json:"deduct_amount,omitempty"`
	// 服务费
	PlatformFee *int64 `json:"platform_fee,omitempty"`
	// 精确已扣除金额
	PreciseDeductAmount *float64 `json:"precise_deduct_amount,omitempty"`
	// 精确服务费
	PrecisePlatformFee *float64 `json:"precise_platform_fee,omitempty"`
	// 精确已退还金额
	PreciseRefundAmount *float64 `json:"precise_refund_amount,omitempty"`
	// 精确任务金额
	PreciseTaskCost *float64 `json:"precise_task_cost,omitempty"`
	// 精确任务总金额
	PreciseTotal *float64 `json:"precise_total,omitempty"`
	// 精确已付金额（元，下同）
	PreciseTotalPaid *float64 `json:"precise_total_paid,omitempty"`
	// 已退还金额
	RefundAmount *int64 `json:"refund_amount,omitempty"`
	// 任务金额
	TaskCost *int64 `json:"task_cost,omitempty"`
	// 任务总金额
	Total *int64 `json:"total,omitempty"`
	// 已付金额（元，下同）
	TotalPaid *int64 `json:"total_paid,omitempty"`
}

/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueCallVirtualNumberRefundDetailGetV2ResponseDataRefundDetailListInner struct for ToolsClueCallVirtualNumberRefundDetailGetV2ResponseDataRefundDetailListInner
type ToolsClueCallVirtualNumberRefundDetailGetV2ResponseDataRefundDetailListInner struct {
	// 广告主id
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	// 广告主名称
	AdvertiserName *string `json:"advertiser_name,omitempty"`
	// 呼叫线索数
	CallClueCnt *int64 `json:"call_clue_cnt,omitempty"`
	// 呼叫总时长，单位：分钟
	CallDurationAll *int64 `json:"call_duration_all,omitempty"`
	// 呼叫话费返款总额
	CallRefundPayMoney *float64 `json:"call_refund_pay_money,omitempty"`
	// 主叫号码
	CallerNumber *string `json:"caller_number,omitempty"`
	// 查询账单月份，格式yyyyMM
	Month *string `json:"month,omitempty"`
}

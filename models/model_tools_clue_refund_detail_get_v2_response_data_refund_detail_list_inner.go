/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueRefundDetailGetV2ResponseDataRefundDetailListInner struct for ToolsClueRefundDetailGetV2ResponseDataRefundDetailListInner
type ToolsClueRefundDetailGetV2ResponseDataRefundDetailListInner struct {
	// 广告计划产生线索的日期，格式yyyyMMdd
	AdCreateClueDate *string `json:"ad_create_clue_date,omitempty"`
	// 广告计划id
	AdId *int64 `json:"ad_id,omitempty"`
	// 广告主id
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	// 广告主名称
	AdvertiserName *string `json:"advertiser_name,omitempty"`
	// 查询账单月份，格式yyyyMM
	Month *string `json:"month,omitempty"`
	// 参与返款的线索数
	ParticipateRefundClueCnt *int64 `json:"participate_refund_clue_cnt,omitempty"`
	// 支付赔付的线索数
	PayRefundClueCnt *int64 `json:"pay_refund_clue_cnt,omitempty"`
	// 赔付金额
	RefundPayMoney *float64 `json:"refund_pay_money,omitempty"`
}

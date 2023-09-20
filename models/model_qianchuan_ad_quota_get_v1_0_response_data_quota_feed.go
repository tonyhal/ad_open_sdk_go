/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdQuotaGetV10ResponseDataQuotaFeed 通投广告在投计划限额信息
type QianchuanAdQuotaGetV10ResponseDataQuotaFeed struct {
	DeliveryInfo *QianchuanAdQuotaGetV10ResponseDataQuotaFeedDeliveryInfo `json:"delivery_info,omitempty"`
	// 当前最高月消耗
	MonthCost *float64 `json:"month_cost,omitempty"`
	// 当前所在限额档位：1,2,3,4,5,6,7,8,9
	QuotaGear *int64                                                `json:"quota_gear,omitempty"`
	QuotaInfo *QianchuanAdQuotaGetV10ResponseDataQuotaFeedQuotaInfo `json:"quota_info,omitempty"`
	StageInfo *QianchuanAdQuotaGetV10ResponseDataQuotaFeedStageInfo `json:"stage_info,omitempty"`
}

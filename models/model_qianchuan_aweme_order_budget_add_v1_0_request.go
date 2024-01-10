/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeOrderBudgetAddV10Request struct for QianchuanAwemeOrderBudgetAddV10Request
type QianchuanAwemeOrderBudgetAddV10Request struct {
	// 广告主id
	AdvertiserId int64 `json:"advertiser_id"`
	// 需要追加预算的订单id
	OrderId int64 `json:"order_id"`
	// 追加的预算
	RenewalBudget int64 `json:"renewal_budget"`
	// 延长的投放时间 1. 短视频，0-7天（步进单位为1天） 2. 0-24小时（步进单位为0.5小时）
	RenewalDeliverySeconds float64 `json:"renewal_delivery_seconds"`
}

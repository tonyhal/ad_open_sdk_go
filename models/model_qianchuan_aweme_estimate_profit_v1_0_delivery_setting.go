/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeEstimateProfitV10DeliverySetting
type QianchuanAwemeEstimateProfitV10DeliverySetting struct {
	//
	Amount  int64                                                  `json:"amount"`
	BidMode QianchuanAwemeEstimateProfitV10DeliverySettingBidMode  `json:"bid_mode"`
	BidType *QianchuanAwemeEstimateProfitV10DeliverySettingBidType `json:"bid_type,omitempty"`
	//
	CpaBid         *float64                                                     `json:"cpa_bid,omitempty"`
	ExternalAction QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction `json:"external_action"`
	// 当ea为326时需要传
	PayRoi *float64 `json:"pay_roi,omitempty"`
}

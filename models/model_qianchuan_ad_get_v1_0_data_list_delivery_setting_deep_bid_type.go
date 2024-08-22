/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdGetV10DataListDeliverySettingDeepBidType
type QianchuanAdGetV10DataListDeliverySettingDeepBidType string

// List of qianchuan_ad_get_v1.0_data_list_delivery_setting_deep_bid_type
const (
	COMM_ROI_QianchuanAdGetV10DataListDeliverySettingDeepBidType         QianchuanAdGetV10DataListDeliverySettingDeepBidType = "COMM_ROI"
	MIN_QianchuanAdGetV10DataListDeliverySettingDeepBidType              QianchuanAdGetV10DataListDeliverySettingDeepBidType = "MIN"
	MIN_SECOND_STAGE_QianchuanAdGetV10DataListDeliverySettingDeepBidType QianchuanAdGetV10DataListDeliverySettingDeepBidType = "MIN_SECOND_STAGE"
	PACING_QianchuanAdGetV10DataListDeliverySettingDeepBidType           QianchuanAdGetV10DataListDeliverySettingDeepBidType = "PACING"
	PAY_ROI_QianchuanAdGetV10DataListDeliverySettingDeepBidType          QianchuanAdGetV10DataListDeliverySettingDeepBidType = "PAY_ROI"
)

// Ptr returns reference to qianchuan_ad_get_v1.0_data_list_delivery_setting_deep_bid_type value
func (v QianchuanAdGetV10DataListDeliverySettingDeepBidType) Ptr() *QianchuanAdGetV10DataListDeliverySettingDeepBidType {
	return &v
}

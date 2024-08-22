/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CustomerCenterAdvertiserListV2AccountSource
type CustomerCenterAdvertiserListV2AccountSource string

// List of customer_center_advertiser_list_v2_account_source
const (
	AD_CustomerCenterAdvertiserListV2AccountSource         CustomerCenterAdvertiserListV2AccountSource = "AD"
	LOCAL_CustomerCenterAdvertiserListV2AccountSource      CustomerCenterAdvertiserListV2AccountSource = "LOCAL"
	ENTERPRISE_CustomerCenterAdvertiserListV2AccountSource CustomerCenterAdvertiserListV2AccountSource = "ENTERPRISE"
	QIANCHUAN_CustomerCenterAdvertiserListV2AccountSource  CustomerCenterAdvertiserListV2AccountSource = "QIANCHUAN"
)

// Ptr returns reference to customer_center_advertiser_list_v2_account_source value
func (v CustomerCenterAdvertiserListV2AccountSource) Ptr() *CustomerCenterAdvertiserListV2AccountSource {
	return &v
}

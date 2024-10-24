/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2DataOptStatus
type AdGetV2DataOptStatus string

// List of ad_get_v2_data_opt_status
const (
	AD_STATUS_DISABLE_AdGetV2DataOptStatus          AdGetV2DataOptStatus = "AD_STATUS_DISABLE"
	AD_STATUS_ENABLE_AdGetV2DataOptStatus           AdGetV2DataOptStatus = "AD_STATUS_ENABLE"
	AD_STATUS_FROZEN_AdGetV2DataOptStatus           AdGetV2DataOptStatus = "AD_STATUS_FROZEN"
	AD_STATUS_DISABLE_BY_QUOTA_AdGetV2DataOptStatus AdGetV2DataOptStatus = "AD_STATUS_DISABLE_BY_QUOTA"
)

// Ptr returns reference to ad_get_v2_data_opt_status value
func (v AdGetV2DataOptStatus) Ptr() *AdGetV2DataOptStatus {
	return &v
}

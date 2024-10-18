/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserInfoV2DataStatus
type AdvertiserInfoV2DataStatus string

// List of advertiser_info_v2_data_status
const (
	STATUS_SELF_SERVICE_UNAUDITED_AdvertiserInfoV2DataStatus AdvertiserInfoV2DataStatus = "STATUS_SELF_SERVICE_UNAUDITED"
	STATUS_WAIT_FOR_BPM_AUDIT_AdvertiserInfoV2DataStatus     AdvertiserInfoV2DataStatus = "STATUS_WAIT_FOR_BPM_AUDIT"
	STATUS_PENDING_CONFIRM_AdvertiserInfoV2DataStatus        AdvertiserInfoV2DataStatus = "STATUS_PENDING_CONFIRM"
	STATUS_ENABLE_AdvertiserInfoV2DataStatus                 AdvertiserInfoV2DataStatus = "STATUS_ENABLE"
	STATUS_LIMIT_AdvertiserInfoV2DataStatus                  AdvertiserInfoV2DataStatus = "STATUS_LIMIT"
	STATUS_CONFIRM_MODIFY_FAIL_AdvertiserInfoV2DataStatus    AdvertiserInfoV2DataStatus = "STATUS_CONFIRM_MODIFY_FAIL"
	STATUS_PENDING_VERIFIED_AdvertiserInfoV2DataStatus       AdvertiserInfoV2DataStatus = "STATUS_PENDING_VERIFIED"
	STATUS_DISABLE_AdvertiserInfoV2DataStatus                AdvertiserInfoV2DataStatus = "STATUS_DISABLE"
	STATUS_CONFIRM_FAIL_AdvertiserInfoV2DataStatus           AdvertiserInfoV2DataStatus = "STATUS_CONFIRM_FAIL"
	STATUS_WAIT_FOR_PUBLIC_AUTH_AdvertiserInfoV2DataStatus   AdvertiserInfoV2DataStatus = "STATUS_WAIT_FOR_PUBLIC_AUTH"
	STATUS_PENDING_CONFIRM_MODIFY_AdvertiserInfoV2DataStatus AdvertiserInfoV2DataStatus = "STATUS_PENDING_CONFIRM_MODIFY"
	STATUS_CONFIRM_FAIL_END_AdvertiserInfoV2DataStatus       AdvertiserInfoV2DataStatus = "STATUS_CONFIRM_FAIL_END"
)

// Ptr returns reference to advertiser_info_v2_data_status value
func (v AdvertiserInfoV2DataStatus) Ptr() *AdvertiserInfoV2DataStatus {
	return &v
}

/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarVasGetBoostGroupListV2DataBoostGroupInfosStatus
type StarVasGetBoostGroupListV2DataBoostGroupInfosStatus string

// List of star_vas_get_boost_group_list_v2_data_boost_group_infos_status
const (
	AUDITING_StarVasGetBoostGroupListV2DataBoostGroupInfosStatus     StarVasGetBoostGroupListV2DataBoostGroupInfosStatus = "AUDITING"
	CANCEL_StarVasGetBoostGroupListV2DataBoostGroupInfosStatus       StarVasGetBoostGroupListV2DataBoostGroupInfosStatus = "CANCEL"
	CLOSE_StarVasGetBoostGroupListV2DataBoostGroupInfosStatus        StarVasGetBoostGroupListV2DataBoostGroupInfosStatus = "CLOSE"
	DOING_StarVasGetBoostGroupListV2DataBoostGroupInfosStatus        StarVasGetBoostGroupListV2DataBoostGroupInfosStatus = "DOING"
	FAIL_CREATE_StarVasGetBoostGroupListV2DataBoostGroupInfosStatus  StarVasGetBoostGroupListV2DataBoostGroupInfosStatus = "FAIL_CREATE"
	FINISH_StarVasGetBoostGroupListV2DataBoostGroupInfosStatus       StarVasGetBoostGroupListV2DataBoostGroupInfosStatus = "FINISH"
	PREPARE_StarVasGetBoostGroupListV2DataBoostGroupInfosStatus      StarVasGetBoostGroupListV2DataBoostGroupInfosStatus = "PREPARE"
	REJECT_AUDIT_StarVasGetBoostGroupListV2DataBoostGroupInfosStatus StarVasGetBoostGroupListV2DataBoostGroupInfosStatus = "REJECT_AUDIT"
	SETTLE_StarVasGetBoostGroupListV2DataBoostGroupInfosStatus       StarVasGetBoostGroupListV2DataBoostGroupInfosStatus = "SETTLE"
	WAIT_CANCEL_StarVasGetBoostGroupListV2DataBoostGroupInfosStatus  StarVasGetBoostGroupListV2DataBoostGroupInfosStatus = "WAIT_CANCEL"
)

// Ptr returns reference to star_vas_get_boost_group_list_v2_data_boost_group_infos_status value
func (v StarVasGetBoostGroupListV2DataBoostGroupInfosStatus) Ptr() *StarVasGetBoostGroupListV2DataBoostGroupInfosStatus {
	return &v
}

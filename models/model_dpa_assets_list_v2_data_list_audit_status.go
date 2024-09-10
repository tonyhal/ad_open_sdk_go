/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaAssetsListV2DataListAuditStatus
type DpaAssetsListV2DataListAuditStatus string

// List of dpa_assets_list_v2_data_list_audit_status
const (
	ASSETS_AUDIT_SUCCESS_DpaAssetsListV2DataListAuditStatus   DpaAssetsListV2DataListAuditStatus = "ASSETS_AUDIT_SUCCESS"
	ASSETS_AUDIT_DENY_DpaAssetsListV2DataListAuditStatus      DpaAssetsListV2DataListAuditStatus = "ASSETS_AUDIT_DENY"
	ASSETS_TO_SUBMIT_AUDIT_DpaAssetsListV2DataListAuditStatus DpaAssetsListV2DataListAuditStatus = "ASSETS_TO_SUBMIT_AUDIT"
	ASSETS_PENDING_AUDIT_DpaAssetsListV2DataListAuditStatus   DpaAssetsListV2DataListAuditStatus = "ASSETS_PENDING_AUDIT"
)

// Ptr returns reference to dpa_assets_list_v2_data_list_audit_status value
func (v DpaAssetsListV2DataListAuditStatus) Ptr() *DpaAssetsListV2DataListAuditStatus {
	return &v
}

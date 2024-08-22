/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsThirdSiteGetV2DataListAuditStatus
type ToolsThirdSiteGetV2DataListAuditStatus string

// List of tools_third_site_get_v2_data_list_audit_status
const (
	AUDIT_UNKNOW_ToolsThirdSiteGetV2DataListAuditStatus   ToolsThirdSiteGetV2DataListAuditStatus = "AUDIT_UNKNOW"
	AUDIT_ACCEPTED_ToolsThirdSiteGetV2DataListAuditStatus ToolsThirdSiteGetV2DataListAuditStatus = "AUDIT_ACCEPTED"
	AUDIT_REJECTED_ToolsThirdSiteGetV2DataListAuditStatus ToolsThirdSiteGetV2DataListAuditStatus = "AUDIT_REJECTED"
	AUDIT_BANNED_ToolsThirdSiteGetV2DataListAuditStatus   ToolsThirdSiteGetV2DataListAuditStatus = "AUDIT_BANNED"
	AUDITING_ToolsThirdSiteGetV2DataListAuditStatus       ToolsThirdSiteGetV2DataListAuditStatus = "AUDITING"
	AWAIT_AUDIT_ToolsThirdSiteGetV2DataListAuditStatus    ToolsThirdSiteGetV2DataListAuditStatus = "AWAIT_AUDIT"
)

// Ptr returns reference to tools_third_site_get_v2_data_list_audit_status value
func (v ToolsThirdSiteGetV2DataListAuditStatus) Ptr() *ToolsThirdSiteGetV2DataListAuditStatus {
	return &v
}

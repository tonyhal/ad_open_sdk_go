/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementBookingGetV2Status
type ToolsAppManagementBookingGetV2Status string

// List of tools_app_management_booking_get_v2_status
const (
	ALL_ToolsAppManagementBookingGetV2Status            ToolsAppManagementBookingGetV2Status = "ALL"
	AUDIT_DOING_ToolsAppManagementBookingGetV2Status    ToolsAppManagementBookingGetV2Status = "AUDIT_DOING"
	AUDIT_REJECTED_ToolsAppManagementBookingGetV2Status ToolsAppManagementBookingGetV2Status = "AUDIT_REJECTED"
	BOOKING_ToolsAppManagementBookingGetV2Status        ToolsAppManagementBookingGetV2Status = "BOOKING"
	ENABLE_ToolsAppManagementBookingGetV2Status         ToolsAppManagementBookingGetV2Status = "ENABLE"
	PAST_DUE_ToolsAppManagementBookingGetV2Status       ToolsAppManagementBookingGetV2Status = "PAST_DUE"
)

// Ptr returns reference to tools_app_management_booking_get_v2_status value
func (v ToolsAppManagementBookingGetV2Status) Ptr() *ToolsAppManagementBookingGetV2Status {
	return &v
}

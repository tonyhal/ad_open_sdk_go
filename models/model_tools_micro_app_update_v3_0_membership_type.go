/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsMicroAppUpdateV30MembershipType
type ToolsMicroAppUpdateV30MembershipType string

// List of tools_micro_app_update_v3.0_membership_type
const (
	ANNUAL_ToolsMicroAppUpdateV30MembershipType       ToolsMicroAppUpdateV30MembershipType = "ANNUAL"
	LIFETIME_ToolsMicroAppUpdateV30MembershipType     ToolsMicroAppUpdateV30MembershipType = "LIFETIME"
	MONTHLY_ToolsMicroAppUpdateV30MembershipType      ToolsMicroAppUpdateV30MembershipType = "MONTHLY"
	NONE_ToolsMicroAppUpdateV30MembershipType         ToolsMicroAppUpdateV30MembershipType = "NONE"
	WEEKLY_DAILY_ToolsMicroAppUpdateV30MembershipType ToolsMicroAppUpdateV30MembershipType = "WEEKLY_DAILY"
)

// Ptr returns reference to tools_micro_app_update_v3.0_membership_type value
func (v ToolsMicroAppUpdateV30MembershipType) Ptr() *ToolsMicroAppUpdateV30MembershipType {
	return &v
}
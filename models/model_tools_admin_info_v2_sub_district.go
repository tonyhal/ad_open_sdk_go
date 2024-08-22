/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAdminInfoV2SubDistrict
type ToolsAdminInfoV2SubDistrict string

// List of tools_admin_info_v2_sub_district
const (
	FOUR_LEVEL_ToolsAdminInfoV2SubDistrict  ToolsAdminInfoV2SubDistrict = "FOUR_LEVEL"
	NONE_ToolsAdminInfoV2SubDistrict        ToolsAdminInfoV2SubDistrict = "NONE"
	ONE_LEVEL_ToolsAdminInfoV2SubDistrict   ToolsAdminInfoV2SubDistrict = "ONE_LEVEL"
	THREE_LEVEL_ToolsAdminInfoV2SubDistrict ToolsAdminInfoV2SubDistrict = "THREE_LEVEL"
	TWO_LEVEL_ToolsAdminInfoV2SubDistrict   ToolsAdminInfoV2SubDistrict = "TWO_LEVEL"
)

// Ptr returns reference to tools_admin_info_v2_sub_district value
func (v ToolsAdminInfoV2SubDistrict) Ptr() *ToolsAdminInfoV2SubDistrict {
	return &v
}

/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarVasGetBoostItemGroupDetailV2DataTaskInfoBoostType
type StarVasGetBoostItemGroupDetailV2DataTaskInfoBoostType string

// List of star_vas_get_boost_item_group_detail_v2_data_task_info_boost_type
const (
	CUSTOM_PACKAGE_StarVasGetBoostItemGroupDetailV2DataTaskInfoBoostType StarVasGetBoostItemGroupDetailV2DataTaskInfoBoostType = "CUSTOM_PACKAGE"
	CUSTOM_TAG_StarVasGetBoostItemGroupDetailV2DataTaskInfoBoostType     StarVasGetBoostItemGroupDetailV2DataTaskInfoBoostType = "CUSTOM_TAG"
	NONE_StarVasGetBoostItemGroupDetailV2DataTaskInfoBoostType           StarVasGetBoostItemGroupDetailV2DataTaskInfoBoostType = "NONE"
)

// Ptr returns reference to star_vas_get_boost_item_group_detail_v2_data_task_info_boost_type value
func (v StarVasGetBoostItemGroupDetailV2DataTaskInfoBoostType) Ptr() *StarVasGetBoostItemGroupDetailV2DataTaskInfoBoostType {
	return &v
}

/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarVasGetBoostGroupListV2DataBoostGroupInfosBoostType
type StarVasGetBoostGroupListV2DataBoostGroupInfosBoostType string

// List of star_vas_get_boost_group_list_v2_data_boost_group_infos_boost_type
const (
	CUSTOM_PACKAGE_StarVasGetBoostGroupListV2DataBoostGroupInfosBoostType StarVasGetBoostGroupListV2DataBoostGroupInfosBoostType = "CUSTOM_PACKAGE"
	CUSTOM_TAG_StarVasGetBoostGroupListV2DataBoostGroupInfosBoostType     StarVasGetBoostGroupListV2DataBoostGroupInfosBoostType = "CUSTOM_TAG"
	NONE_StarVasGetBoostGroupListV2DataBoostGroupInfosBoostType           StarVasGetBoostGroupListV2DataBoostGroupInfosBoostType = "NONE"
)

// Ptr returns reference to star_vas_get_boost_group_list_v2_data_boost_group_infos_boost_type value
func (v StarVasGetBoostGroupListV2DataBoostGroupInfosBoostType) Ptr() *StarVasGetBoostGroupListV2DataBoostGroupInfosBoostType {
	return &v
}
